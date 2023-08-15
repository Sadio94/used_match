package rest

import "time"

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

type Inner struct {
	EqCSFieldString  string
	NeCSFieldString  string
	GtCSFieldString  string
	GteCSFieldString string
	LtCSFieldString  string
	LteCSFieldString string
}

type KindValid struct {
	Inner                 Inner
	RequiredString        string    `form:"requiredString" json:"requiredString" validate:"required"`  // 必传
	OmitemptyString       string    `form:"omitemptySting" json:"omitemptySting" validate:"omitempty"` // 非必传
	RequiredNumber        int       `form:"requiredNumber" json:"requiredNumber" validate:"required"`
	RequiredNumberError   int       `form:"requiredNumberError" json:"requiredNumberError" validate:"required, oneof=0, 30"`  // 如果这里指定`required`, 那这个值就不能是改类型的默认值
	RequiredMultiple      []string  `form:"requiredMultiple" json:"requiredMultiple" validate:"required,max=7,dive,required"` // 数组长度最大是7, 数组内元素必传    														// RequiredMultiple是nil或是长度大于0的数组且数组中的元素不能是空字符串
	LenString             string    `form:"lenString" json:"lenString" validate:"len=1"`
	LenNumber             float64   `form:"lenNumber" json:"lenNumber" validate:"len=1113.00"`
	EqString              string    `form:"eqString" json:"eqString" validate:"min=1,max=90"`                                  // 字符串长度1<=n<=90								// 字符串长度1<=n<=90
	NeString              string    `form:"neString" json:"neString" validate:"ne="`                                           // NeString不能等于空字符串
	LtString              string    `form:"ltString" json:"ltString" validate:"lt=3"`                                          // LtString长度必须小于3个字符
	LtTime                time.Time `form:"ltTime" json:"ltTime" validate:"lt"`                                                // LtTime必须小于当前日期和时间
	LteTime               time.Time `form:"lteTime" json:"lteTime" validate:"lte"`                                             // LteTime必须小于或等于当前日期和时间
	GtTime                time.Time `form:"gtTime" json:"gtTime" validate:"gt"`                                                // GtTime必须大于当前日期和时间
	GteTime               time.Time `form:"gteTime" json:"gteTime" validate:"gte"`                                             // GteTime必须大于或等于当前日期和时间
	EqFieldString         string    `form:"eqFieldString" json:"eqFieldString" validate:"eqfield=EqString"`                    // EqFieldString必须等于EqString
	EqCSFieldString       string    `form:"eqCSFieldString" json:"eqCSFieldString" validate:"eqcsfield=Inner.EqCSFieldString"` // EqCSFieldString必须等于Inner.EqCSFieldString
	NeCSFieldString       string    `form:"neCSFieldString" json:"neCSFieldString" validate:"necsfield=Inner.NeCSFieldString"` // NeCSFieldString不能等于Inner.NeCSFieldString
	AlphaString           string    `form:"alphaString" json:"alphaString" validate:"alpha"`                                   // AlphaString只能包含字母
	AlphanumString        string    `form:"alphanumString" json:"alphanumString" validate:"alphanum"`                          // AlphanumString只能包含字母和数字
	AlphanumUnicodeString string    `form:"alphanumUnicodeString" json:"alphanumUnicodeString" validate:"alphanumunicode"`     // AlphanumUnicodeString只能包含字母数字和Unicode字符
	NumberString          string    `form:"numberString" json:"numberString" validate:"number"`                                // NumberString必须是一个有效的数字
	Email                 string    `form:"email" json:"email" validate:"email"`                                               // Email必须是一个有效的邮箱
	URL                   string    `form:"url" json:"url" validate:"url"`                                                     // URL必须是一个有效的URL
	URI                   string    `form:"uri" json:"uri" validate:"uri"`                                                     // URI必须是一个有效的URI
	Contains              string    `form:"contains" json:"contains" validate:"contains=purpose"`                              // Contains必须包含文本'purpose'
	ContainsAny           string    `form:"containsAny" json:"containsAny" validate:"containsany=!@#$"`                        // ContainsAny必须包含至少一个以下字符'!@#$'
	Excludes              string    `form:"excludes" json:"excludes" validate:"excludes=text"`                                 // Excludes不能包含文本'text'
	EndsWith              string    `form:"endsWith" json:"endsWith" validate:"endswith=end"`                                  // EndsWith必须以文本'end'结尾
	StartsWith            string    `form:"startsWith" json:"startsWith" validate:"startswith=start"`                          // StartsWith必须以文本'start'开头
	ASCII                 string    `form:"ascii" json:"ascii" validate:"ascii"`                                               // ASCII必须只包含ascii字符
	IP                    string    `form:"ip" json:"ip" validate:"ip"`                                                        // IP必须是一个有效的IP地址
	IPv4                  string    `form:"ipv4" json:"ipv4" validate:"ipv4"`                                                  // IPv4必须是一个有效的IPv4地址
	IPv6                  string    `form:"ipv6" json:"ipv6" validate:"ipv6"`                                                  // IPv6必须是一个有效的IPv6地址
	JsonString            string    `form:"jsonString" json:"jsonString" validate:"json"`                                      // 是一个json字符串
	LowercaseString       string    `form:"lowercaseString" json:"lowercaseString" validate:"lowercase"`                       // LowercaseString必须是小写字母
	UppercaseString       string    `form:"uppercaseString" json:"uppercaseString" validate:"uppercase"`                       // UppercaseString必须是大写字母
	Datetime              string    `form:"datetime" json:"datetime" validate:"datetime=2006-01-02"`                           // Datetime的格式必须是2006-01-02
}

type QueryRequest struct {
	Index int `form:"index" json:"index" binding:"omitempty,min=0"`         // 分页起始位置(默认0)
	Count int `form:"count" json:"count" binding:"omitempty,min=1,max=100"` // 分页数量(默认15)
}

type RegisterInfo struct {
	NickName       string `form:"nickName" json:"nickName" binding:"required,min=1,max=30"` // 最长不超过30
	Email          string `form:"email" json:"email" binding:"required,email"`              // 必传 符合邮箱格式
	Password       string `form:"password" json:"password" binding:"required"`
	RepeatPassword string `form:"repeatPassword" json:"repeatPassword" binding:"required,eqfield=Password"` // 需要和Password字段一样
	RegisterType   int    `form:"registerType" json:"registerType" binding:"required,oneof=1 2 3 4"`        // 之一
}

type UserList struct {
	QueryRequest
	FilterName string `form:"filterName" json:"filterName" binding:"omitempty,min=1,max=30"` // 非必传 但是传了之后需要满足tag对应的规则
	UserStatus int    `form:"userStatus" json:"userStatus" binding:"required,oneof=1 2"`
}

type DayTypeRequest struct {
	Dates string `form:"dates" json:"dates" binding:"required"` // 2006-01-02形式 多个之间以英文逗号','分隔
}

type IdRequest struct {
	Id string `uri:"id" json:"id" binding:"required,len=32" validate:"required,len=32"`
}

type Lsyz struct {
	StartDate string `uri:"start_time" json:"start_date" binding:"omitempty" validate:"omitempty"` // 起始时间 2022-07-23
	EndDate   string `uri:"end_time" json:"end_date" binding:"omitempty" validate:"omitempty"`     // 结束日期 2022-07-24
	DocIds    string `uri:"doc_ids" json:"doc_ids" binding:"omitempty" validate:"omitempty"`       // 当前项目下筛选的文档 多个以英文逗号','拼接 不超过30个
	ProjectId string `uri:"project_id" json:"project_id" binding:"required" validate:"required"`   // 项目id
	Token     string `uri:"token" json:"token" binding:"required" validate:"required"`             // 用户token
}

type PageRequest struct {
	Page  int `form:"page" json:"page" binding:"omitempty,min=1"`   // 第一页开始
	Limit int `form:"limit" json:"limit" binding:"omitempty,min=5"` // 默认5条/页
}

type LsyzC struct {
	PageRequest
	Lsyz
}

type AddProjectRequest struct {
	Token string `uri:"token" json:"token" binding:"required" validate:"required"`                       // 用户token
	Name  string `uri:"name" json:"name" binding:"required" validate:"required"`                         // 项目名称
	Type  int64  `uri:"type" json:"type" binding:"required,oneof=1 2 3" validate:"required,oneof=1 2 3"` // 项目类型 1:企业 2:个人 3:其他
	Note  string `uri:"note" json:"note" binding:"omitempty" validate:"omitempty"`                       // 项目备注
}

type ProjectListRequest struct {
	PageRequest
	Token     string `uri:"token" json:"token" binding:"required" validate:"required"`                               // 用户token
	DateDesc  int64  `uri:"date_desc" json:"date_desc" binding:"omitempty,oneof=0 1" validate:"omitempty,oneof=0 1"` // 项目创建时间排序搜索项 0:默认倒序 1:正序
	Type      int64  `uri:"type" json:"type" binding:"omitempty,oneof=1 2 3" validate:"omitempty,oneof=1 2 3"`       // 项目类型搜索项 1:企业 2:个人 3:其他
	KeyWord   string `uri:"keyword" json:"keyword" binding:"omitempty" validate:"omitempty"`                         // 关键字搜索 包含项目名称和备注的模糊匹配
	StartTime int64  `uri:"start_time" json:"start_time" binding:"omitempty" validate:"omitempty"`                   // 开始时间 秒级时间戳
	EndTime   int64  `uri:"end_time" json:"end_time" binding:"omitempty" validate:"omitempty"`                       // 结束时间 秒级时间戳
}

type UpdateProjectRequest struct {
	Token     string `uri:"token" json:"token" binding:"required" validate:"required"`                         // 用户token
	ProjectId string `uri:"project_id" json:"project_id" binding:"required" validate:"required"`               // 项目id
	Name      string `uri:"name" json:"name" binding:"omitempty" validate:"omitempty"`                         // 项目名称
	Type      int64  `uri:"type" json:"type" binding:"omitempty,oneof=1 2 3" validate:"omitempty,oneof=1 2 3"` // 项目类型 1:企业 2:个人 3:其他
	Note      string `uri:"note" json:"note" binding:"omitempty" validate:"omitempty"`                         // 项目备注
}

type DeleteProjectRequest struct {
	Token     string `uri:"token" json:"token" binding:"required" validate:"required"`           // 用户token
	ProjectId string `uri:"project_id" json:"project_id" binding:"required" validate:"required"` // 待删除项目id
}

type JydsOverviewRequest struct {
	Lsyz
	DispType int64 `uri:"disp_type" json:"disp_type" binding:"omitempty,oneof=0 1 2" validate:"omitempty,oneof=0 1 2"` // 对手方概览展示 0:汇总 1:收入 2:支出 默认展示汇总页
}
type JydsClassRequest struct {
	Lsyz
	DispType int64 `uri:"disp_type" json:"disp_type" binding:"omitempty,oneof=0 1 2" validate:"omitempty,oneof=0 1 2"` // 对手方分类展示 0:汇总 1:收入 2:支出 默认展示汇总页
}

type JydsMonitorObjectRequest struct {
	PageRequest
	Lsyz
	DispType int64 `uri:"disp_type" json:"disp_type" binding:"omitempty,oneof=0 1 2" validate:"omitempty,oneof=0 1 2"` // 对手方分类展示 0:汇总 1:收入 2:支出 默认展示汇总页
}

type JyhzBalanceFluctuationRequest struct {
	Lsyz
}

type JyhzTransactionDistributionRequest struct {
	Lsyz
	DispType int64 `uri:"disp_type" json:"disp_type" binding:"omitempty,oneof=0 1" validate:"omitempty,oneof=0 1"` // 交易分布展示 0:交易金额 1:交易笔数 默认展示交易金额页
}

type JydsAbnormalTransactionHighFrequencyRequest struct {
	Lsyz
}

type JydsAbnormalTransactionSuspiciousRequest struct {
	PageRequest
	Lsyz
}

type JydsAbnormalTransactionLargeRequest struct {
	PageRequest
	Lsyz
}

type JyfxTopUserRequest struct {
	Lsyz
	Type string `uri:"type" json:"type" binding:"omitempty,oneof=income expenses " validate:"omitempty,oneof=income expenses"` // 筛选类型 income:收入 expenses:支出 默认收入
}

type JyfxTopUserDetailRequest struct {
	Lsyz
	Type string `uri:"type" json:"type" binding:"omitempty,oneof=income expenses " validate:"omitempty,oneof=income expenses"` // 筛选类型 income:收入 expenses:支出 默认收入
}

type JyfxTopSummaryRequest struct {
	Lsyz
	Type string `uri:"type" json:"type" binding:"omitempty,oneof=income expenses " validate:"omitempty,oneof=income expenses"` // 筛选类型 income:收入 expenses:支出 默认收入
}

type DocListRequest struct {
	PageRequest
	Token             string `uri:"token" json:"token" binding:"required" validate:"required"`                                                   // 用户token
	DateDesc          int64  `uri:"date_desc" json:"date_desc" binding:"omitempty,oneof=0 1" validate:"omitempty,oneof=0 1"`                     // 文件上传时间排序搜索项 0:默认倒序 1:正序
	KeyWord           string `uri:"keyword" json:"keyword" binding:"omitempty" validate:"omitempty"`                                             // 关键字搜索 文件名称或者备注模糊匹配
	StartTime         int64  `uri:"start_time" json:"start_time" binding:"omitempty" validate:"omitempty"`                                       // 开始时间 秒级时间戳
	EndTime           int64  `uri:"end_time" json:"end_time" binding:"omitempty" validate:"omitempty"`                                           // 结束时间 秒级时间戳
	HeaderCheckStatus int64  `uri:"header_check_status" json:"header_check_status" binding:"omitempty,oneof=0 1" validate:"omitempty,oneof=0 1"` // 表头校验搜索项 0:不通过 1:通过
	DocIds            string `uri:"doc_ids" json:"doc_ids" binding:"omitempty" validate:"omitempty"`                                             // 当前项目下筛选的文档 多个以英文逗号','拼接
	OcrStatus         int64  `uri:"ocr_status" json:"ocr_status" binding:"omitempty,oneof=-1 0 1" validate:"omitempty,oneof=-1 0 1"`             // ocr识别状态 -1:识别失败 0:识别中 1:识别成功
}

type UpdateDocRequest struct {
	Token string `uri:"token" json:"token" binding:"required" validate:"required"`   // 用户token
	DocId string `uri:"doc_id" json:"doc_id" binding:"required" validate:"required"` // 待编辑文档
	Note  string `uri:"note" json:"note" binding:"omitempty" validate:"omitempty"`   // 文档备注
	Title string `uri:"title" json:"title" binding:"omitempty" validate:"omitempty"` // 文档名
}

type DocDetailRequest struct {
	Token string `uri:"token" json:"token" binding:"required" validate:"required"`   // 用户token
	DocId string `uri:"doc_id" json:"doc_id" binding:"required" validate:"required"` // 待查看文档id
}

type DeleteDocRequest struct {
	Token string `uri:"token" json:"token" binding:"required" validate:"required"`   // 用户token
	DocId string `uri:"doc_id" json:"doc_id" binding:"required" validate:"required"` // 待删除文档id
}

type DocUploadRequest struct {
	Token     string `uri:"token" json:"token" binding:"required" validate:"required"`           // 用户token
	FileType  string `uri:"file_type" json:"file_type" binding:"required" validate:"required"`   // 上传文档类型 默认.jpg
	Func      string `uri:"func" json:"func" binding:"required" validate:"required"`             // 方法名 可能是为了区别业务 银行流水项目默认 'bankbills'
	Platform  string `uri:"platform" json:"platform" platform:"required" validate:"required"`    // 默认 'web'
	ProjectId string `uri:"project_id" json:"project_id" binding:"required" validate:"required"` // 项目id
	DocId     string `uri:"doc_id" json:"doc_id" binding:"required" validate:"required"`         // 文档id
	Title     string `uri:"title" json:"title" binding:"required" validate:"required"`           // 文档名
	Password  string `uri:"password" json:"password" binding:"omitempty" validate:"omitempty"`   // 密码
}

type OCRRequest struct {
	Token      string `uri:"token" json:"token" binding:"required" validate:"required"`       // 用户token
	DocId      string `uri:"doc_id" json:"doc_id" binding:"required" validate:"required"`     // ocr识别文档id
	PageId     string `uri:"page_id" json:"page_id" binding:"omitempty" validate:"omitempty"` // 文档下的某一页
	Position   bool   `uri:"position" json:"position" binding:"omitempty" validate:"omitempty"`
	MoreThread int64  `uri:"more_thread" json:"more_thread" binding:"omitempty" validate:"omitempty"`
}
