package dao

//import (
//	errorCode "gitlab2.aishu.cn/Zeus/framework/i18n-sdk-go/errors"
//	db "gitlab2.aishu.cn/Zeus/service/hypersla/sla-core"
//	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-core/models"
//	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-core/types"
//	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-core/utils"
//	"time"
//	"xorm.io/builder"
//	"xorm.io/xorm"
//)
//
//var GroupOperate = GroupDao{}
//
//type GroupDao struct{}
//
//func (this *GroupDao) Insert(session *xorm.Session, group *models.Group) (string, error) {
//	if session != nil {
//		_, err := session.Insert(group)
//		if err != nil {
//			return "", utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//		return group.Id, nil
//	} else {
//		_, err := db.ClientMgrDB.Insert(group)
//		if err != nil {
//			return "", utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//		return group.Id, nil
//	}
//
//}
//
//func (this *GroupDao) GetById(id string) (*models.Group, error) {
//	group := &models.Group{Id: id}
//	has, err := db.ClientMgrDB.Get(group)
//	if err != nil {
//		return nil, utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//	}
//
//	if !has {
//		return nil, nil
//	}
//
//	return group, nil
//}
//
//func (this *GroupDao) GetByIds(ids []string) (int64, []*models.Group, error) {
//	var groups []*models.Group
//	n, err := db.ClientMgrDB.In("id", ids).FindAndCount(&groups)
//	if err != nil {
//		return 0, nil, utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//	}
//
//	return n, groups, nil
//}
//
//func (this *GroupDao) GetByIdsWithQuery(ids []string, query map[string]interface{}) (int64, []*models.Group, error) {
//	var groups []*models.Group
//	n, err := db.ClientMgrDB.In("id", ids).Where(query).FindAndCount(&groups)
//	if err != nil {
//		return 0, nil, utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//	}
//
//	return n, groups, nil
//}
//
//func (this *GroupDao) GetByQuery(query map[string]interface{}) (int64, []*models.Group, error) {
//	var groups []*models.Group
//	n, err := db.ClientMgrDB.Where(query).FindAndCount(&groups)
//	if err != nil {
//		return 0, nil, utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//	}
//
//	return n, groups, nil
//}
//
//func (this *GroupDao) Update(session *xorm.Session, id string, group *models.Group) (string, error) {
//	if session != nil {
//		_, err := session.Where("id = ?", id).AllCols().Update(group)
//		if err != nil {
//			return "", utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//
//		return group.Id, nil
//	} else {
//		_, err := db.ClientMgrDB.Where("id = ?", id).AllCols().Update(group)
//		if err != nil {
//			return "", utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//
//		return group.Id, nil
//	}
//}
//
//func (this *GroupDao) UpdateById(session *xorm.Session, id string, param map[string]interface{}) error {
//	if session != nil {
//		_, err := session.Table(new(models.Group)).Where("id = ?", id).Update(param)
//		if err != nil {
//			return utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//		return nil
//	} else {
//		_, err := db.ClientMgrDB.Table(new(models.Group)).Where("id = ?", id).Update(param)
//		if err != nil {
//			return utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//		return nil
//	}
//}
//
//func (this *GroupDao) UpdateByIdsWithQuery(session *xorm.Session, ids []string, param map[string]interface{}) error {
//	if session != nil {
//		_, err := session.Table(new(models.Group)).In("id", ids).Update(param)
//		if err != nil {
//			return utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//		return nil
//	} else {
//		_, err := db.ClientMgrDB.Table(new(models.Group)).In("id", ids).Update(param)
//		if err != nil {
//			return utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//		return nil
//	}
//}
//
//func (this *GroupDao) FindList(filterName string, ids []interface{}, index, count int, status types.ValidateStatusType, userId, tenantId string, isOc bool) (int64, []*models.Group, error) {
//	var groups []*models.Group
//	cond := builder.NewCond()
//	// 权限和用户名过滤
//	if isOc {
//		cond = cond.And(builder.Eq{"user_id": userId})
//	} else {
//		cond = cond.And(builder.Eq{"tenant_id": tenantId})
//		if len(ids) > 0 {
//			cond = cond.And(builder.In("user_id", ids...))
//		}
//	}
//
//	// SLA名过滤
//	if len(filterName) > 0 {
//		cond = cond.And(builder.Like{"name", filterName})
//	}
//
//	/*
//		已生效意味着到期时间大于等于当前时间且生效时间小于等于当前时间
//		未生效意味着生效时间大于等于当前时间
//		已到期意味着在自定义到期时间的时候到期时间小于等于当前时间
//	*/
//	switch status {
//	case types.ValidateStatusV:
//		cond = (cond.And(builder.Eq{"validate_period": types.ValidatePeriodC}).And(builder.Gte{"expired_time": time.Now().UnixNano() / 1e6}).And(builder.Lte{"effective_time": time.Now().UnixNano() / 1e6})).
//			Or(cond.And(builder.Eq{"validate_period": types.ValidatePeriodF}).And(builder.Lte{"effective_time": time.Now().UnixNano() / 1e6}))
//	case types.ValidateStatusU:
//		cond = cond.And(builder.Gte{"effective_time": time.Now().UnixNano() / 1e6})
//	case types.ValidateStatusE:
//		cond = cond.And(builder.Eq{"validate_period": types.ValidatePeriodC}).And(builder.Lte{"expired_time": time.Now().UnixNano() / 1e6})
//	}
//
//	n, err := db.ClientMgrDB.Where(cond).Limit(count, index).Desc("update_time").FindAndCount(&groups)
//	if err != nil {
//		return 0, nil, utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//	}
//
//	return n, groups, nil
//}
//
//func (this *GroupDao) DeleteByIds(session *xorm.Session, ids []string) error {
//	if session != nil {
//		_, err := session.In("id", ids).Delete(new(models.Group))
//		if err != nil {
//			return utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//
//		return nil
//	} else {
//		_, err := db.ClientMgrDB.In("id", ids).Delete(new(models.Group))
//		if err != nil {
//			return utils.New(errorCode.HyperSLAMgm_OperationDatabaseError, map[string]string{"detail": err.Error()})
//		}
//
//		return nil
//	}
//}
