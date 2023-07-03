package models

//import (
//	"github.com/Zeus/service/hypersla/sla-core/types"
//	"github.com/Zeus/service/hypersla/sla-core/utils"
//	"time"
//)
//
//type Group struct {
//	Id             string                   `xorm:"'id' not null pk comment('主键') VARCHAR(32)"`
//	Type           types.BusinessType       `xorm:"'type' not null comment('业务类型') TINYINT(2)"`
//	Name           string                   `xorm:"'name' unique(combine_name_to_user) not null comment('策略名称或SLA名') VARCHAR(255)"`
//	Remark         string                   `xorm:"'remark'  comment('备注')  VARCHAR(255)"`
//	ValidatePeriod types.ValidatePeriodType `xorm:"'validate_period' comment('有效期') INT(11)"`
//	PeriodSpan     int                      `xorm:"'period_span' comment('有效期时长') INT(11)"`
//	PeriodType     types.ValidateType       `xorm:"'period_type' comment('有效期时长类型') INT(11)"`
//	EffectiveType  types.EffectiveCType     `xorm:"'effective_type' comment('生效时间') INT(11)"`
//	EffectiveTime  int64                    `xorm:"'effective_time' comment('生效时间, 时间戳') BIGINT(20)"`
//	TimeZone       string                   `xorm:"'time_zone'  comment('时区字段  ,默认用户时区')  VARCHAR(32)"`
//	UserId         string                   `xorm:"'user_id' unique(combine_name_to_user) index not null comment('创建用户id') VARCHAR(32)"`
//	TenantId       string                   `xorm:"'tenant_id' index not null comment('用户所属租户id') VARCHAR(32)"`
//	UpdateTime     int64                    `xorm:"'update_time' not null comment('更新时间') BIGINT(20)"`
//	ExpiredTime    int64                    `xorm:"'expired_time' not null comment('到期时间, 时间戳') BIGINT(20)"`
//	DisableMark    bool                     `xorm:"'disable_mark' comment('禁用启用状态') TINYINT(1)"`
//	InstancedMark  types.CronInstancedType  `xorm:"'instanced_mark' comment('实例化状态 主要用于最终一致性的判断') INT(11)"`
//}
//
//func (this *Group) BeforeInsert() {
//	if this.Id == "" {
//		this.Id = utils.NewUUID()
//	}
//
//	this.UpdateTime = time.Now().UnixNano() / 1e6
//
//	if this.ValidatePeriod == types.ValidatePeriodF {
//		// 永久有效
//		this.ExpiredTime = 0
//	} else {
//		this.ExpiredTime = utils.GetExpiredTime(this.PeriodSpan, this.PeriodType, this.EffectiveTime)
//	}
//}
//
//func (this *Group) BeforeUpdate() {
//	this.UpdateTime = time.Now().UnixNano() / 1e6
//
//	if this.ValidatePeriod == types.ValidatePeriodF {
//		// 永久有效
//		this.ExpiredTime = 0
//	} else {
//		this.ExpiredTime = utils.GetExpiredTime(this.PeriodSpan, this.PeriodType, this.EffectiveTime)
//	}
//}
//
//func (this *Group) TableName() string {
//	return "strategy_group"
//}
