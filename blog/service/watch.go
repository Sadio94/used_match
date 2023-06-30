package watch

import (
	"context"
	"encoding/json"
	"github.com/go-co-op/gocron"
	"gitlab2.aishu.cn/Zeus/service/hyperbackup/gprc-proto-go/backup_worker"
	"gitlab2.aishu.cn/Zeus/service/hyperbackup/gprc-proto-go/common"
	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-core/types"
	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-core/utils"
	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-worker/internal/business"
	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-worker/internal/dao"
	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-worker/internal/external"
	. "gitlab2.aishu.cn/Zeus/service/hypersla/sla-worker/internal/initials"
	. "gitlab2.aishu.cn/Zeus/service/hypersla/sla-worker/internal/utils"
	"runtime/debug"
	"strings"
	"time"
)

type IWatch interface {
	LoadScheduleInstance(ctx context.Context)
	LeaderScheduleFailover(ctx context.Context)
	LeaderRetentionTrigger() *gocron.Scheduler
}

type BusinessWatch struct{}

func NewBusinessWatch() IWatch {
	return &BusinessWatch{}
}

// 本节点监控
func (this *BusinessWatch) LoadScheduleInstance(ctx context.Context) {
	Logger.Info("Load Schedule Instance...")
	defer func() {
		if e := recover(); e != nil {
			Logger.Errorf("Load Schedule Instance Err:%#v, stack:%s", e, string(debug.Stack()[:]))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		_, nodeSlots, err := dao.SlotOperate.GetByQuery(map[string]interface{}{"slot_code": MachineCode})
		if err != nil {
			Logger.Error(err)
			// 释放cpu占用时间
			time.Sleep(time.Second * 2)
			continue
		}

		for _, slot := range nodeSlots {
			relativeGroup, err := dao.GroupOperate.GetById(slot.LoadId)
			if err != nil || relativeGroup == nil {
				Logger.Warnf("groupId[%v] To RelativeGroup Not Found...", slot.LoadId)
				continue
			}

			// 划分不同维度
			if relativeGroup.InstancedMark != types.CronInstancedC {
				// 表中是未实例化状态 可能是由于rpc通知晚于监控 这里走加载失败重试的逻辑
				business.AddJobForNotify(slot.LoadId)
			} else {
				if !business.HasInCronTask(slot.LoadId) {
					// 内存中没有 但是这是前期已经加载成功的 只是可能由于服务重启 直接加载且不重试
					err := business.AddJobForWatch(slot.LoadId)
					if err != nil {
						go func() {
							// 尝试获取sla信息
							var eventSLAId string = slot.LoadId
							var eventSLAName string = slot.LoadId
							eventGroup, errI := dao.GroupOperate.GetById(slot.LoadId)
							if errI == nil && eventGroup != nil {
								eventSLAId = eventGroup.Id
								eventSLAName = eventGroup.Name
							}
							SLAByte, _ := json.Marshal(map[string]string{"SLAName": eventSLAName, "SLAId": eventSLAId})
							SLAStr := string(SLAByte)

							// 尝试获取报错原因
							ex := utils.NewError(err)
							eventDesc, eventSolution := GetErrorInfoToEvent(ex.GetCode(), ex.GetParams())
							AddCronTaskFailedPushEvent(SLAStr, eventDesc, eventSolution)
						}()
					}
				}
			}
		}

		time.Sleep(40 * time.Second)
	}
}

// leader故障迁移
func (this *BusinessWatch) LeaderScheduleFailover(ctx context.Context) {
	Logger.Info("Leader Schedule Failover...")
	defer func() {
		if e := recover(); e != nil {
			Logger.Errorf("Leader Schedule Failover Err:%#v, stack:%s", e, string(debug.Stack()[:]))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		time.Sleep(80 * time.Second)
		if external.NewETCDSDKService().IsLeader() {
			downCodeArr, waitCodeArr, err := external.NewETCDSDKService().GetRealAllNodeInfo()
			if err != nil {
				Logger.Warn(err)
			}

			Logger.Debugf("Leader Get downCodeArr is:%v, waitCodeArr is:%v", downCodeArr, waitCodeArr)
			for _, downCode := range downCodeArr {
				// 节点故障 毫无疑问直接启动迁移
				downCount, downSlot, err := dao.SlotOperate.GetByQuery(map[string]interface{}{"slot_code": downCode})
				if err != nil {
					Logger.Warn(err)
					continue
				}

				// 这里存在计算哈希的过程 考虑到时候后台监控线程 所以这里选择在for循环下面写表
				for _, slot := range downSlot {
					works, err := external.NewETCDSDKService().GetServiceInstance()
					if err != nil {
						Logger.Warn(err)
						downCount -= 1
						continue
					}

					if works == nil || len(works) < 1 {
						Logger.Warn("no HyperSLAWorkerService")
						downCount -= 1
						continue
					}

					hash := utils.GetHash(slot.LoadId)
					slotCode := strings.Split(works[int(hash)%len(works)], ";")[0]
					err = dao.SlotOperate.UpdateById(nil, slot.Id, map[string]interface{}{"slot_code": slotCode})
					if err != nil {
						Logger.Warn(err)
						downCount -= 1
						continue
					}
				}

				if downCount > 0 {
					// 上报故障迁移完成事件
					go SLAFailoverPushEvent(downCode)
				}
			}

			// 保护机制 休眠30s之后查看本时刻离线服务的状态
			if waitCodeArr != nil && len(waitCodeArr) > 0 {
				time.Sleep(30 * time.Second)
				for _, waitCode := range waitCodeArr {
					isOnline, err := external.NewETCDSDKService().IsNodeOnline(waitCode)
					if err != nil {
						Logger.Warn(err)
						continue
					}
					if !isOnline {
						// 仍离线 启动迁移
						waitCont, waitSlot, err := dao.SlotOperate.GetByQuery(map[string]interface{}{"slot_code": waitCode})
						if err != nil {
							Logger.Warn(err)
							continue
						}

						for _, slot := range waitSlot {
							works, err := external.NewETCDSDKService().GetServiceInstance()
							if err != nil {
								Logger.Warn(err)
								waitCont -= 1
								continue
							}

							if works == nil || len(works) < 1 {
								Logger.Warn("no HyperSLAWorkerService")
								waitCont -= 1
								continue
							}

							hash := utils.GetHash(slot.LoadId)
							slotCode := strings.Split(works[int(hash)%len(works)], ";")[0]
							err = dao.SlotOperate.UpdateById(nil, slot.Id, map[string]interface{}{"slot_code": slotCode})
							if err != nil {
								Logger.Warn(err)
								waitCont -= 1
								continue
							}
						}

						if waitCont > 0 {
							// 上报故障迁移成功事件
							go SLAFailoverPushEvent(waitCode)
						}
					}
				}
			}
		}
	}

}

// leader触发保留策略
func (this *BusinessWatch) LeaderRetentionTrigger() *gocron.Scheduler {
	c := gocron.NewScheduler(time.Local)
	c.Every(1).Days().At("00:00").Do(func() {
		defer func() {
			if e := recover(); e != nil {
				Logger.Errorf("Leader Retention Trigger Err:%#v, stack:%s", e, string(debug.Stack()[:]))
			}
		}()

		Logger.Info("Leader Retention Trigger...")
		if external.NewETCDSDKService().IsLeader() {
			binds, err := dao.BindOperate.GetAll()
			if err != nil {
				Logger.Error(err)
				return
			}

			if binds == nil || len(binds) < 1 {
				return
			}

			var groupIds []string
			for _, item := range binds {
				if !utils.IntStringInRealize(item.WorkerId, groupIds) {
					groupIds = append(groupIds, item.WorkerId)
				}
			}

			// 一次性取表
			_, groups, err := dao.GroupOperate.GetByIds(groupIds)
			if err != nil {
				Logger.Error(err)
				return
			}

			if groups == nil || len(groups) < 1 {
				return
			}

			for _, bind := range binds {
				req := &backup_worker.RunRetentionStrategyReqProto{
					ObjectId: bind.ObjectId,
				}

				// 禁用 有效期判断
				var bindU bool
				for _, iGroup := range groups {
					if iGroup.Id == bind.WorkerId {
						if iGroup.DisableMark || (utils.GetSLAValidateStatus(iGroup.ValidatePeriod, iGroup.EffectiveTime, iGroup.ExpiredTime) != types.ValidateStatusV) {
							bindU = true
						}
					}
				}

				retentionInfo := &backup_worker.RetentionStrategyProto{}

				if bindU {
					// 这里给保留全部副本
					retentionInfo.Mode = common.StrategyRetentionModeProto_STRATEGY_RETENTION_MODE_UNKNOWN
				} else {
					// 这里给用户配置
					_, bindRetention, err := dao.RetentionOperate.GetByQuery(map[string]interface{}{"group_id": bind.WorkerId, "type": types.SlaBackup})
					if err != nil {
						Logger.Error(err)
						continue
					}

					if bindRetention == nil || len(bindRetention) != 1 {
						continue
					}

					if bindRetention[0].IsRetentionOpen {
						retentionInfo.Mode = common.StrategyRetentionModeProto(bindRetention[0].RetentionType)
						retentionInfo.ReplicasNum = int64(bindRetention[0].QuantityNum)
						if bindRetention[0].RetentionType == types.RetentionTypeByDuration && bindRetention[0].DurationConfiguration != nil {
							retentionInfo.DataPeriodType = common.DataRetentionPeriodProto(bindRetention[0].DurationConfiguration.Type)
							retentionInfo.DataPeriodCount = int64(bindRetention[0].DurationConfiguration.Num)
						}
						retentionInfo.IsPeriodRetentionEnabled = bindRetention[0].IsGfs
						if bindRetention[0].GfsConfiguration != nil {
							retentionInfo.IsDayPeriodRetentionOptionEnabled = bindRetention[0].GfsConfiguration.GfsDayEnable
							retentionInfo.DayPeriodRetentionNum = int64(bindRetention[0].GfsConfiguration.GfsDayNum)
							retentionInfo.IsWeekPeriodRetentionOptionEnabled = bindRetention[0].GfsConfiguration.GfsWeekEnable
							retentionInfo.WeekPeriodRetentionNum = int64(bindRetention[0].GfsConfiguration.GfsWeekNum)
							retentionInfo.IsMonthPeriodRetentionOptionEnabled = bindRetention[0].GfsConfiguration.GfsMonthEnable
							retentionInfo.MonthPeriodRetentionNum = int64(bindRetention[0].GfsConfiguration.GfsMonthNum)
							retentionInfo.IsQuarterPeriodRetentionOptionEnabled = bindRetention[0].GfsConfiguration.GfsQuarterEnable
							retentionInfo.QuarterPeriodRetentionNum = int64(bindRetention[0].GfsConfiguration.GfsQuarterNum)
							retentionInfo.IsYearPeriodRetentionOptionEnabled = bindRetention[0].GfsConfiguration.GfsYearEnable
							retentionInfo.YearPeriodRetentionNum = int64(bindRetention[0].GfsConfiguration.GfsYearNum)
						}
					} else {
						retentionInfo.Mode = common.StrategyRetentionModeProto_STRATEGY_RETENTION_MODE_UNKNOWN
					}
					req.RetentionStrategy = retentionInfo
					_, err = external.NewGrpcFuncManage("", 0, 0).RunRetentionStrategy(req)
					if err != nil {
						Logger.Error(err)
					}
				}
			}
		}
	})

	c.StartAsync()
	return c
}
