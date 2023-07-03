package rest

//import "gitlab2.aishu.cn/Zeus/service/hypersla/sla-core/types"
//
//type IdRequest struct {
//	Id string `uri:"id" json:"id" binding:"required,len=32" validate:"required,len=32"`
//}
//
//type QueryRequest struct {
//	Index int `form:"index" json:"index" binding:"omitempty,min=0"         validate:"omitempty,min=0"`         // 分页起始位置(默认0)
//	Count int `form:"count" json:"count" binding:"omitempty,min=1,max=100" validate:"omitempty,min=1,max=100"` // 分页数量(默认15)
//}
//
//// 基本信息
//type SLABaseInfo struct {
//	Name           string                   `form:"name"       json:"name"       binding:"required,min=1,max=90"   validate:"required,min=1,max=90"`
//	Remark         string                   `form:"remark"       json:"remark"       binding:"omitempty,min=1,max=384"   validate:"omitempty,min=1,max=384"`
//	ValidatePeriod types.ValidatePeriodType `form:"validatePeriod"       json:"validatePeriod"       binding:"required,oneof=1 2"   validate:"required,oneof=1 2"`
//	PeriodType     types.ValidateType       `form:"periodType"       json:"periodType"       binding:"omitempty,oneof=1 2 3 4 5 6 7"   validate:"omitempty,oneof=1 2 3 4 5 6 7"`
//	PeriodSpan     int                      `form:"periodSpan"       json:"periodSpan"    binding:"omitempty,min=1,max=9999"   validate:"omitempty,min=1,max=9999"`
//	EffectiveType  types.EffectiveCType     `form:"effectiveType"       json:"effectiveType"       binding:"required,oneof=1 2"   validate:"required,oneof=1 2"`
//	EffectiveTime  int64                    `form:"effectiveTime"       json:"effectiveTime"    binding:"required,min=1"   validate:"required,min=1"`
//	TimeZone       string                   `form:"timeZone"       json:"timeZone"    binding:"omitempty"   validate:"omitempty"`
//}
//
//// 备份策略配置
//type BackupConfigC struct {
//	PlanCommon
//	WindowsCommon
//	RetentionCommon
//}
//
//// 如果后续策略加进来没有那么多计划的配置 可以继续拆分成更小的结构体
//type PlanCommon struct {
//	MinEnable            bool                   `form:"minEnable"     json:"minEnable"    binding:"omitempty"   validate:"omitempty"`
//	MinConfiguration     []*types.MinConfig     `form:"minConfiguration"     json:"minConfiguration"   binding:"omitempty,max=4,dive,required"  validate:"omitempty,max=4,dive,required"`
//	HourEnable           bool                   `form:"hourEnable"     json:"hourEnable"    binding:"omitempty"   validate:"omitempty"`
//	HourConfiguration    []*types.HourConfig    `form:"hourConfiguration"     json:"hourConfiguration"   binding:"omitempty,max=4,dive,required"  validate:"omitempty,max=4,dive,required"`
//	DayEnable            bool                   `form:"dayEnable"     json:"dayEnable"    binding:"omitempty"   validate:"omitempty"`
//	DayConfiguration     []*types.DayConfig     `form:"dayConfiguration"     json:"dayConfiguration"   binding:"omitempty,max=4,dive,required"  validate:"omitempty,max=4,dive,required"`
//	WeekEnable           bool                   `form:"weekEnable"     json:"weekEnable"    binding:"omitempty"   validate:"omitempty"`
//	WeekConfiguration    []*types.WeekConfig    `form:"weekConfiguration"     json:"weekConfiguration"   binding:"omitempty,max=4,dive,required"  validate:"omitempty,max=4,dive,required"`
//	MonthEnable          bool                   `form:"monthEnable"     json:"monthEnable"    binding:"omitempty"   validate:"omitempty"`
//	MonthConfiguration   []*types.MonthConfig   `form:"monthConfiguration"     json:"monthConfiguration"   binding:"omitempty,max=4,dive,required"  validate:"omitempty,max=4,dive,required"`
//	QuarterEnable        bool                   `form:"quarterEnable"     json:"quarterEnable"    binding:"omitempty"   validate:"omitempty"`
//	QuarterConfiguration []*types.QuarterConfig `form:"quarterConfiguration"     json:"quarterConfiguration"   binding:"omitempty,max=4,dive,required"  validate:"omitempty,max=4,dive,required"`
//	YearEnable           bool                   `form:"yearEnable"     json:"yearEnable"    binding:"omitempty"   validate:"omitempty"`
//	YearConfiguration    []*types.YearConfig    `form:"yearConfiguration"     json:"yearConfiguration"   binding:"omitempty,max=4,dive,required"  validate:"omitempty,max=4,dive,required"`
//}
//
//// 窗口配置也是一样 这里就把禁用和限速窗口都放进去了
//type WindowsCommon struct {
//	DisableWindow *types.DisableWindowConfig `form:"disableWindow"     json:"disableWindow"   binding:"omitempty"  validate:"omitempty"`
//	LimitWindow   *types.LimitWindowConfig   `form:"limitWindow"     json:"limitWindow"   binding:"omitempty"  validate:"omitempty"`
//}
//
//type RetentionCommon struct {
//	IsRetentionOpen       bool                  `form:"isRetentionOpen"     json:"isRetentionOpen"    binding:"omitempty"   validate:"omitempty"`
//	RetentionType         types.RetentionCType  `form:"retentionType"       json:"retentionType"       binding:"omitempty,oneof=1 2"   validate:"omitempty,oneof=1 2"`
//	DurationConfiguration *types.DurationConfig `form:"durationConfiguration"     json:"durationConfiguration"   binding:"omitempty"  validate:"omitempty"`
//	QuantityNum           int                   `form:"quantityNum"       json:"quantityNum"       binding:"omitempty,min=1,max=9999"   validate:"omitempty,min=1,max=9999"`
//	IsGfs                 bool                  `form:"isGfs"     json:"isGfs"    binding:"omitempty"   validate:"omitempty"`
//	GfsConfiguration      *types.GfsConfig      `form:"gfsConfiguration"     json:"gfsConfiguration"   binding:"omitempty"  validate:"omitempty"`
//}
//
//// 添加SLA
//type AddSLARequest struct {
//	SLABaseInfo
//	BackupConfig *BackupConfigC `form:"backupConfig"       json:"backupConfig"    binding:"required"   validate:"required"`
//}
//
//// 编辑SLA备份配置
//type UpdateBackupRequest struct {
//	BackupConfig *BackupConfigC `form:"backupConfig"       json:"backupConfig"    binding:"required"   validate:"required"`
//}
//
//// 编辑SLA基本信息
//type UpdateBaseInfoRequest struct {
//	SLABaseInfo
//}
//
//// 获取SLA列表
//type GetSLAListRequest struct {
//	QueryRequest
//	Filter         *types.ListFilter        `form:"filter"       json:"filter"    binding:"omitempty"   validate:"omitempty"` // 模糊过滤参数
//	ValidateStatus types.ValidateStatusType `form:"validateStatus"     json:"validateStatus"    binding:"required,oneof=1 2 3 4"   validate:"required,oneof=1 2 3 4"`
//}
//
//// 保护对象下编辑SLA 本质上是另存为
//type AddSLAForResourceRequest struct {
//	SLABaseInfo
//	ObjectId     string         `form:"objectId" json:"objectId" binding:"required,len=32" validate:"required,len=32"`
//	ResourceId   string         `form:"resourceId" json:"resourceId" binding:"required,len=32" validate:"required,len=32"`
//	BackupConfig *BackupConfigC `form:"backupConfig"       json:"backupConfig"    binding:"required"   validate:"required"`
//}
//
//// 批量删除SLA
//type BatchDeleteSLARequest struct {
//	IdSet []string `form:"idSet"  json:"idSet"  binding:"required,max=100,dive,required,len=32" validate:"required,max=100,dive,required,len=32"` // 待删除SLA集合
//}
//
//// 批量启用禁用SLA
//type BatchUpdateSLAStatusRequest struct {
//	IdSet  []string `form:"idSet"  json:"idSet"  binding:"required,max=100,dive,required,len=32" validate:"required,max=100,dive,required,len=32"` // SLA集合
//	Status bool     `form:"status"     json:"status"    binding:"omitempty"   validate:"omitempty"`
//}
//
//// 名称存在性校验
//type SLANameExistsRequest struct {
//	Name string `form:"name"       json:"name"       binding:"required,min=1,max=90"   validate:"required,min=1,max=90"`
//}
//
//// SLA下绑定的保护对象信息 弹窗接口
//type HasBindRequest struct {
//	IdStrings string `form:"idStrings"       json:"idStrings"       binding:"required,min=32,max=3299"   validate:"required,min=32,max=3299"` // SLA id 前端以','拼接
//}
//
//// 保护对象信息
//type RemoveObjectsInfo struct {
//	ObjectId   string `form:"objectId"  json:"objectId"  binding:"required,len=32" validate:"required,len=32"`
//	ObjectName string `form:"objectName"       json:"objectName"       binding:"required,min=1,max=255"   validate:"required,min=1,max=255"` // 不做代码逻辑处理 主要是提高弹窗可读性体验
//}
//
//// 批量解绑一个SLA下的多个保护对象
//type SLARemoveObjectsRequest struct {
//	Id          string               `form:"id" json:"id" binding:"required,len=32" validate:"required,len=32"`
//	ObjectInfos []*RemoveObjectsInfo `form:"objectInfos"  json:"objectInfos"  binding:"required,max=100,dive,required" validate:"required,max=100,dive,required"`
//}
//
//// 单个SLA下绑定的保护对象列表信息
//type BindObjectsListRequest struct {
//	QueryRequest
//	Name    string `form:"name"       json:"name"       binding:"omitempty,min=1,max=255"   validate:"omitempty,min=1,max=255"`       // 保护对象名模糊过滤
//	AppType string `form:"appType"       json:"appType"       binding:"omitempty,min=1,max=255"   validate:"omitempty,min=1,max=255"` // 应有类型筛选
//}
