/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  response
 * @Date: 2023/07/03 11:32
 */

package rest

import "time"

type UserInfo struct {
	Id           string    `json:"id"`
	NickName     string    `json:"nickName"`
	UserStatus   int       `json:"userStatus"`
	RegisterTime time.Time `json:"registerTime"`
	UpdateTime   time.Time `json:"updateTime"`
}

type Paging struct {
	TotalNum uint64      `json:"totalNum"`
	MetaInfo []*UserInfo `json:"metaInfo"`
}

type Result struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type Result1 struct {
	Ret  int         `json:"ret"`
	Err  interface{} `json:"err"`
	Data interface{} `json:"data"`
}

type DateInfo []DateI

type DateI struct {
	Date string `json:"date"` // 具体日期 2022-09-10
	Type int    `json:"type"` // 当天类型枚举 0：工作日 1：假日 2：节假日
}

// LsyzAResp 流水验证-数据真实性
type LsyzAResp struct {
	ShouldInterest int64             `json:"should_interest"` // 应结息
	ActualInterest int64             `json:"actual_interest"` // 实结息
	AbnormalTimes  int64             `json:"abnormal_times"`  // 结息金额异常次数
	TotalInterest  int64             `json:"total_interest"`  // 总结息金额
	InterestReport []InterestReportI `json:"interest_report"` // 结息报表
	StartDate      string            `json:"start_date"`      // 筛选之后显示的流水开始时间
	EndDate        string            `json:"end_date"`        // 筛选之后显示的流水结束时间
}

type InterestReportI struct {
	InterestTime     string `json:"interest_time"`     // 结息日期 后端格式化之后返回
	TradeTime        string `json:"trade_time"`        // 交易日期 后端格式化之后返回
	Summary          string `json:"summary"`           // 交易摘要
	Interest         int64  `json:"interest"`          // 利息 单位元
	EstimateInterest int64  `json:"estimate_interest"` // 预估利息 单位元
}

// LsyzCAccountResp 流水验证-数据完整性
type LsyzCAccountResp struct {
	ThemeAccount        []ThemeAccountInfo `json:"theme_account"`          // 主体账户
	NoneSameNameAccount []NSMInfo          `json:"none_same_name_account"` // 非同名账户信息
}

type ThemeAccountInfo struct {
	AssociateInfo []AssociateInfoI `json:"associate_info"` // 主体账户涉及的银行账户
	ThemeName     string           `json:"theme_name"`     // 主体账户名
}

type AssociateInfoI struct {
	Name    string `json:"name"`    // 隶属银行
	Account string `json:"account"` // 银行账户
}

type NSMInfo struct {
	TradeName        string `json:"trade_name"`         // 交易户名
	TradeAccount     string `json:"trade_account"`      // 交易账号
	TradePeriod      string `json:"trade_period"`       // 交易时间段
	TransferIn       int64  `json:"transfer_in"`        // 向同名账户转入 单位元
	TransferInTimes  int64  `json:"transfer_in_times"`  // 向同名账户转入次数
	TransferOut      int64  `json:"transfer_out"`       // 向同名账户转出 单位元
	TransferOutTimes int64  `json:"transfer_out_times"` // 向同名账户转出次数
}

// LsyzBalanceResp 流水验证-余额完整
type LsyzBalanceResp struct {
	IntervalType  int64            `json:"interval_type"`  // 区间类型 0月 1周
	BalanceType   []BalanceTypeI   `json:"balance_type"`   // 月余额类型
	TotalNum      int64            `json:"total_num"`      // 余额总月数
	ErrorNum      int64            `json:"error_num"`      // 月错误总月数
	AbnormalTrade []AbnormalTradeI `json:"abnormal_trade"` // 异常记录
}

type BalanceTypeI struct {
	Timer string `json:"timer"` // 2023-01
	TType int64  `json:"type"`  // 1 缺失 2 错误 3 正确
}

type AbnormalTradeI struct {
	TradeTime      string `json:"trade_time"`      // 交易时间
	InitialBalance int64  `json:"initial_balance"` // 初期余额
	TradeAmount    int64  `json:"trade_amount"`    // 交易金额
	EndingBalance  int64  `json:"ending_balance"`  // 期末余额
	Summary        string `json:"summary"`         // 交易摘要
	Name           string `json:"name"`            // 交易户名
	Account        string `json:"account"`         // 交易账号
}

// LsyzSummaryResp 流水验证-摘要完整
type LsyzSummaryResp struct {
	IntervalType  int64             `json:"interval_type"`  // 区间类型 0月 1周
	SummaryType   []SummaryTypeI    `json:"summary_type"`   // 月摘要类型
	TotalNum      int64             `json:"total_num"`      // 摘要总月数
	ErrorNum      int64             `json:"error_num"`      // 摘要缺失月数
	AbnormalTrade []AbnormalTradeI1 `json:"abnormal_trade"` // 异常记录
}

type SummaryTypeI struct {
	Timer string `json:"timer"` // 2023-01
	TType int64  `json:"type"`  // 1 摘要完整 2 摘要缺失
}

type AbnormalTradeI1 struct {
	TradeTime     string `json:"trade_time"`     // 交易时间
	TradeAmount   int64  `json:"trade_amount"`   // 交易金额
	EndingBalance int64  `json:"ending_balance"` // 期末余额
	Summary       string `json:"summary"`        // 交易摘要
	Name          string `json:"name"`           // 交易户名
	Account       string `json:"account"`        // 交易账号
}

// AddProjectResp 新建项目
type AddProjectResp struct {
	ProjectId int64 `json:"project_id"` // 新建成功之后的项目id
}

// ProjectListResp 项目列表
type ProjectListResp struct {
	Count           int64          `json:"count"`              // 项目总数
	DocCount        int64          `json:"doc_count"`          // 文档总数
	ReportCount     int64          `json:"report_count"`       // 报告总数
	AmountCount     string         `json:"amount_count"`       // 金额总数 从金额个位开始 三位使用','拼接 例如 123,333,333
	OverBillion     bool           `json:"over_billion"`       // 金额总数是否超过10亿
	ProjectList     []ProjectListI `json:"project_list"`       // 项目详情
	RiskLevel1Count int            `json:"risk_level_1_count"` // 高风险项目数
	RiskLevel2Count int            `json:"risk_level_2_count"` // 中风险项目数
	RiskLevel3Count int            `json:"risk_level_3_count"` // 低风险项目数
}

type ProjectListI struct {
	Note       string `json:"note"`        // 项目备注
	ProjectId  int64  `json:"project_id"`  // 项目id
	CreateTime int64  `json:"create_time"` // 项目创建时间
	Name       string `json:"name"`        // 项目名称
	DocCount   int64  `json:"doc_count"`   // 当前项目下文件总数
	Type       int64  `json:"type"`        // 项目类型
}

// JydsOverviewResp 交易对手-对手方概览
type JydsOverviewResp struct {
	TradeInfo []JydsOverviewTradeInfo `json:"trade_info"` // 交易信息 top15+其他 数组最大长度16
}

type JydsOverviewTradeInfo struct {
	Id          int64  `json:"id"`           // 序号
	Name        string `json:"name"`         // 交易户名
	Account     string `json:"account"`      // 交易账号
	Amount      int64  `json:"amount"`       // 交易金额
	StartTime   string `json:"start_time"`   // 交易时间段开始 2022-08-09
	EndTime     string `json:"end_time"`     // 交易时间段结束 2023-02-18
	TradeNumber int64  `json:"trade_number"` // 交易笔数
}

// JydsClassResp 交易对手-对手方分类
type JydsClassResp struct {
	Class JydsClassSummary `json:"class"` // 分类
}

type JydsClassSummary struct {
	CounterpartyClass
	TradeInfo []ClassTradeInfo `json:"trade_info"` // 对手方交易明细 top8
}

type ClassTradeInfo struct {
	Name      string `json:"name"`       // 对手方名称
	OutAmount int64  `json:"out_amount"` // 支出金额
	InAmount  int64  `json:"in_amount"`  // 收入金额
}

// CounterpartyClass 对手方分类
type CounterpartyClass struct {
	InstitutionalClients int64 `json:"institutional_clients"` // 机构客户数量
	NaturalPerson        int64 `json:"natural_person"`        // 自然人数量
	Own                  int64 `json:"own"`                   // 自己数量
	Other                int64 `json:"other"`                 // 其他数量
}

// JydsMonitorObjectResp 交易对手-重点监测对象
type JydsMonitorObjectResp struct {
	Monitor JydsMonitorObjectSummary `json:"monitor"` // 监测
}

type JydsMonitorObjectSummary struct {
	TradeClassType
	TotalNum  int64              `json:"total_num"`  // 所有流水的条数
	TradeInfo []MonitorTradeInfo `json:"trade_info"` // 所有流水详情
}

type TradeClassType struct {
	Loan       int64 `json:"loan"`        // 贷款
	Investment int64 `json:"investment"`  // 投资
	Financial  int64 `json:"financial"`   // 金融服务
	RealEstate int64 `json:"real_estate"` // 房地产
	Other      int64 `json:"other"`       // 其他
}

type MonitorTradeInfo struct {
	Signal      string `json:"signal"`       // 预警信号
	Name        string `json:"name"`         // 交易户名
	Account     string `json:"account"`      // 交易账号
	StartTime   string `json:"start_time"`   // 交易时间段开始 2022-08-09
	EndTime     string `json:"end_time"`     // 交易时间段结束 2023-02-18
	TradeNumber int64  `json:"trade_number"` // 交易笔数
	Amount      int64  `json:"amount"`       // 交易金额
}

// BalanceFluctuationResp 交易汇总-余额波动
type BalanceFluctuationResp struct {
	BalanceFluctuation []BalanceInfo `json:"balance_fluctuation"` // 余额波动
}

type BalanceInfo struct {
	Timer      string `json:"timer"`       // 月 2022-01
	AvgBalance int64  `json:"avg_balance"` // 月均余额
	Volatility int64  `json:"volatility"`  // 波动率 标注差 20
}

// TransactionDistribution 交易汇总-交易分布
type TransactionDistribution struct {
	TradeData []TradeAmountI `json:"trade_data"` // 交易金额或者交易笔数信息 根据disp_type来填充返回值
}

type TradeAmountI struct {
	Timer  string `json:"timer"`  // 月 2022-01
	Income int64  `json:"income"` // 当月收入或者收入笔数
	Pay    int64  `json:"pay"`    // 当月支出或者支出笔数
	Total  int64  `json:"total"`  // 当月总流水或者交易总笔数
}

// AbnormalTransactionHighFrequencyResp 交易汇总-异常交易-高频交易
type AbnormalTransactionHighFrequencyResp struct {
	HighFrequencyTrading []HighFrequencyTradingI `json:"high_frequency_trading"` // 高频交易
}

type HighFrequencyTradingI struct {
	Name      string `json:"name"`       // 对手方
	Income    int64  `json:"income"`     // 收入
	Pay       int64  `json:"pay"`        // 支出
	IncomeNum int64  `json:"income_num"` // 收入笔数
	PayNum    int64  `json:"pay_num"`    // 支出笔数
}

// AbnormalTransactionSuspiciousResp 交易汇总-异常交易-可疑交易
type AbnormalTransactionSuspiciousResp struct {
	SuspiciousTransaction []SuspiciousTransactionI `json:"suspicious_transaction"` // 可疑交易
	SuspiciousNum         int64                    `json:"suspicious_num"`         // 可以交易流水总数
}

type SuspiciousTransactionI struct {
	Time         string `json:"time"`           // 时间
	Amount       int64  `json:"amount"`         // 金额
	Name         string `json:"name"`           // 交易户名
	Account      string `json:"account"`        // 交易账号
	IsWorkingDay bool   `json:"is_working_day"` // 时间是否工作日
}

// AbnormalTransactionLargeResp 交易汇总-异常交易-大额交易
type AbnormalTransactionLargeResp struct {
	LargeTransaction []LargeTransactionI `json:"large_transaction"` // 大额交易
	LargeNum         int64               `json:"large_num"`         // 大额交易流水总数
}

type LargeTransactionI struct {
	Time       string `json:"time"`       // 时间
	Amount     int64  `json:"amount"`     // 金额
	Name       string `json:"name"`       // 交易户名
	Account    string `json:"account"`    // 交易账号
	Proportion string `json:"proportion"` // 当月占比
}

// JyfxSummaryResp 经营分析-经营概览
type JyfxSummaryResp struct {
	StartRestAmount         int64         `json:"start_rest_amount"`          // 期初余额
	EndRestAmount           int64         `json:"end_rest_amount"`            // 期末余额
	TotalIncome             int64         `json:"total_income"`               // 总收入
	TotalExpenses           int64         `json:"total_expenses"`             // 总支出
	AvgDailyRestAmount      int64         `json:"avg_daily_rest_amount"`      // 日均余额
	InterestBasedRestAmount int64         `json:"interest_based_rest_amount"` // 利息反推余额
	RestAmountList          []AmountListI `json:"rest_amount_list"`           // 月份概览详情
}

type AmountListI struct {
	Keys       string `json:"keys"`        // 月份
	Income     int64  `json:"income"`      // 收入
	Expenses   int64  `json:"expenses"`    // 支出
	RestAmount int64  `json:"rest_amount"` // 余额
}

// JyfxTopUserResp 经营分析-交易集中度
type JyfxTopUserResp struct {
	AmountTop        []TopInfoI `json:"amount_top"`         // 收入之和占总收入
	AmountTopPerson  []TopInfoI `json:"amount_top_person"`  //top5个人客户
	AmountTopCompany []TopInfoI `json:"amount_top_company"` // top5企业客户
	TotalAmount      int64      `json:"total_amount"`       // 总计金额
}

type TopInfoI struct {
	OtherName   string `json:"other_name"`   // 客户账号
	UserType    string `json:"user_type"`    // person 个人/company 公司
	TotalAmount int64  `json:"total_amount"` // 总计金额
}

// JyfxTopUserDetailResp 经营分析-交易集中度详情
type JyfxTopUserDetailResp struct {
	AmountTop   []TopInfoI `json:"amount_top"`   // 收入 支出信息
	TotalAmount int64      `json:"total_amount"` // 总计金额
	AllKey      []string   `json:"all_key"`      // 所有月份数组
}

type TopUserDetailI struct {
	OtherName   string         `json:"other_name"`   // 客户账号
	UserType    string         `json:"user_type"`    // person 个人/company 公司
	TotalAmount int64          `json:"total_amount"` // 总计金额
	TotalCount  int64          `json:"total_count"`  // 交易笔数
	AmountList  []AmountListII `json:"amount_list"`  // 月信息
}

type AmountListII struct {
	Key    string  `json:"key"`    // 月份
	Amount float64 `json:"amount"` // 收入
}

// JyfxTopSummaryResp 经营分析-流水构成
type JyfxTopSummaryResp struct {
	AmountTop   []TopSummaryI `json:"amount_top"`   // 摘要流水信息
	TotalAmount int64         `json:"total_amount"` // 总计金额
}

type TopSummaryI struct {
	Summary string  `json:"summary"` // 流水摘要
	Count   int64   `json:"count"`   // 交易笔数
	Amount  float64 `json:"amount"`  // 此摘要下金额
}

// DocListResp 项目下文档列表
type DocListResp struct {
	FileNum  int64      `json:"file_num"`  // 文档数
	FileList []DocListI `json:"file_list"` // 文档详情
}

type DocListI struct {
	DocId             int64  `json:"doc_id"`              // 文档id
	FirstPageId       string `json:"first_page_id"`       // 若该文档包含多页 这是第一页id
	HeaderCheckStatus int64  `json:"header_check_status"` // 表头校验结果
	OcrStatus         int64  `json:"ocr_status"`          // ocr识别结果
	Note              string `json:"note"`                // 项目备注
	CreateTime        int64  `json:"create_time"`         // 项目创建时间
	ModifyTime        int64  `json:"modify_time"`         // 文档变更时间
	UploadTime        int64  `json:"upload_time"`         // 文档上传时间
	Title             string `json:"title"`               // 文档名
}

// DocDetailResp 文档详情
type DocDetailResp struct {
	NextDocId string   `json:"next_doc_id"` // 按时间顺序的下个文档id
	LastDocId string   `json:"last_doc_id"` // 按时间顺序的上个文档id
	DocInfo   DocInfoI `json:"doc_info"`    // 文档详情信息
}

type DocInfoI struct {
	DocId      string         `json:"doc_id"`      // 文档id
	Title      string         `json:"title"`       // 文档名
	UploadTime int64          `json:"upload_time"` // 文档上传时间
	CreateTime int64          `json:"create_time"` // 项目创建时间
	PageList   []DocPageListI `json:"page_list"`   // 文档下页信息 一个文档可能是多页的
}

type DocPageListI struct {
	ModifyTime int64  `json:"modify_time"` // 页变更时间
	Title      string `json:"title"`       // 页标题
	Note       string `json:"note"`        // 页备注
	FileName   string `json:"file_name"`   // 页名
	PageId     string `json:"page_id"`     // 页id
	Rotate     int64  `json:"rotate"`
}

// LsglListResp 流水管理-流水详情
type LsglListResp struct {
	TotalNum  int64        `json:"total_num"`  // 流水总条数
	TradeList []TradeListI `json:"trade_list"` // 流水详情
}

type TradeListI struct {
	TradeTime string `json:"trade_time"` // 交易时间
	Amount    int64  `json:"amount"`     // 交易金额
	Summary   string `json:"summary"`    // 交易摘要
	Name      string `json:"name"`       // 交易户名
	Account   string `json:"account"`    // 交易账号
	IsRelated bool   `json:"is_related"` // 是否关联方
}

// GljyRelatedOverviewResp 关联交易-关联方概览/关联方交易概览
type GljyRelatedOverviewResp struct {
	RelatedParty       []RelatedPartyI       `json:"related_party"`       // 关联方概览
	Quantity           int64                 `json:"quantity"`            // 关联方数量
	IncomeTotal        int64                 `json:"income_total"`        // 关联方收入
	PayTotal           int64                 `json:"pay_total"`           // 关联方支出
	Difference         int64                 `json:"difference"`          // 差额
	RelatedTransaction []RelatedTransactionI `json:"related_transaction"` // 关联交易概览详情
}

type RelatedPartyI struct {
	Name    string `json:"name"`    // 户名
	Account string `json:"account"` // 账号
}

type RelatedTransactionI struct {
	Name       string `json:"name"`       // 关联方
	Income     int64  `json:"income"`     // 总收入
	Pay        int64  `json:"pay"`        // 总支出
	IncomeNum  int64  `json:"income_num"` // 总收入笔数
	PayNum     int64  `json:"pay_num"`    // 总支出笔数
	Difference int64  `json:"difference"` // 差额
}

// GljyRelatedInformationResp 关联交易-关联方信息
type GljyRelatedInformationResp struct {
	Information []InformationI `json:"information"` // 关联方详细信息
}

type InformationI struct {
	RelatedTransactionI
	Reason string `json:"reason"` // 关联方原因
}

// GljyRelatedDetailResp 关联交易-关联交易详情
type GljyRelatedDetailResp struct {
	TotalNum          int64                `json:"total_num"`          // 流水总条数
	TransactionDetail []TransactionDetailI `json:"transaction_detail"` // 流水详情
}

type TransactionDetailI struct {
	TradeTime string `json:"trade_time"` // 交易时间
	Amount    int64  `json:"amount"`     // 交易金额
	Summary   string `json:"summary"`    // 交易摘要
	Name      string `json:"name"`       // 交易户名
	Account   string `json:"account"`    // 交易账号
	Reason    string `json:"reason"`     // 关联关系
}

// GljyRelatedTransactionResp 关联交易-关联交易
type GljyRelatedTransactionResp struct {
	IncomeTop5        []Top5I              `json:"income_top_5"`       // 关联方收入top5
	PayTop5           []Top5I              `json:"pay_top_5"`          // 关联方支出top5
	TradeDistribution []TradeDistributionI `json:"trade_distribution"` // 关联方交易时间分布
	AllRelatedParty   []string             `json:"all_related_party"`  // 展示在关联交易时间分布下的所有关联方数组
}

type Top5I struct {
	Name   string `json:"name"`   // 关联方户名
	Amount int64  `json:"amount"` // 交易金额
}

type TradeDistributionI struct {
	Timer  string  `json:"timer"`  // 时间 2022-01
	Income []Top5I `json:"income"` // 存在收入的关联方数组
	Pay    []Top5I `json:"pay"`    // 存在支出的关联方数组
}

// ReportListResp 分析报告-主页列表
type ReportListResp struct {
	TotalNum   int64         `json:"total_num"`   // 所有报告的条数
	ReportInfo []ReportListI `json:"report_info"` // 报告详情
}

type ReportListI struct {
	Id            int64  `json:"id"`             // 报告id 后台生成
	Name          string `json:"name"`           // 报告名称
	GenerateTime  string `json:"generate_time"`  // 报告生成时间 格式化2022-01-10 12:00:00的形式
	BelongProject string `json:"belong_project"` // 报告所属项目名
	ProjectType   int64  `json:"project_type"`   // 项目类型 1:企业 2:个人 3:其他
	Status        int64  `json:"status"`         // 报告生成状态 0:生成中 1:生成完成
}

// ReportDetailResp 分析报告-项目报告详情-项目历史报告
type ReportDetailResp struct {
	HistoricalReport []ReportDetailI `json:"historical_report"` // 项目历史报告
}

type ReportDetailI struct {
	Id           int64  `json:"id"`            // 报告id 后台生成
	Name         string `json:"name"`          // 报告名称
	GenerateTime string `json:"generate_time"` // 报告生成时间 格式化2022-01-10 12:00:00的形式
	Status       int64  `json:"status"`        // 报告生成状态 0:生成中 1:生成完成
}

// OCRResultQueryResp 文档ocr识别结果
type OCRResultQueryResp struct {
	Result OCRResultI `json:"result"` // 返回结果
}

type OCRResultI struct {
	DocHeaders DocHeadersI `json:"doc_headers"` // 表头信息等
	Pages      []PagesI    `json:"pages"`       // 表格完整信息
}

type DocHeadersI struct {
	ALLKVItemsI
}

type PagesI struct {
	ALLKVItemsI
	ExceptTableCellAmountItems []ExceptTableCellAmountItemsI `json:"except_table_cell_amount_items"` // 表格内金额错误校验信息
	ExceptTableCellDateItems   []ExceptTableCellDateItemsI   `json:"except_table_cell_date_items"`   // 表格内日期错误校验信息
	HeaderTemplate             int64                         `json:"header_template"`                // 表头模板
	Tables                     []TablesI                     `json:"tables"`                         // 表格内信息
}

type ExceptTableCellAmountItemsI struct {
	StartRow     int64  `json:"start_row"`
	EndRow       int64  `json:"end_row"`
	StartCol     int64  `json:"start_col"`
	EndCol       int64  `json:"end_col"`
	ExtraAmount  int64  `json:"extra_amount"`  // 原始识别后的金额
	ExpectAmount int64  `json:"expect_amount"` // 建议修改金额
	Text         string `json:"text"`          // 原始识别后的金额
}

type ExceptTableCellDateItemsI struct {
	StartRow   int64  `json:"start_row"`
	EndRow     int64  `json:"end_row"`
	StartCol   int64  `json:"start_col"`
	EndCol     int64  `json:"end_col"`
	ErrType    int64  `json:"err_type"`    // 错误值
	ExceptText string `json:"except_text"` // 建议修改时间
	Text       string `json:"text"`        // 原始识别后的时间
}

type ALLKVItemsI struct {
	KVItems            []KVItemsI      `json:"kv_items"`              // 表格外信息抽取
	ExpectKVItems      []KVItemsI      `json:"expect_kv_items"`       // 按需 表格外信息
	TableKVItems       []TableKVItemsI `json:"table_kv_items"`        // 表格内信息抽取
	ExpectTableKVItems []TableKVItemsI `json:"expect_table_kv_items"` // 按需 表格内信息抽取
}

type KVItemsI struct {
	Key   string `json:"key"`   // 对应键名
	Value string `json:"value"` // 键对应的值
	Text  string `json:"text"`  // 键对应的值
}

type TableKVItemsI struct {
	Key      string `json:"key"`       // 对应键名
	Value    string `json:"value"`     // 键对应的值
	Text     string `json:"text"`      // 键对应的值
	StartCol int64  `json:"start_col"` // 第几列
	StartRow int64  `json:"start_row"` // 第几行
}

type TablesI struct {
	TableCols  int64        `json:"table_cols"`  // 表格多少列
	TableRows  int64        `json:"table_rows"`  // 表格多少行
	TableCells []TableCells `json:"table_cells"` // 单元格信息
}

type TableCells struct {
	StartRow int64  `json:"start_row"`
	EndRow   int64  `json:"end_row"`
	StartCol int64  `json:"start_col"`
	EndCol   int64  `json:"end_col"`
	Text     string `json:"text"` // 当前单元格对应的值
}

type UpdateHeaderResp struct {
	UploadTime int64 `json:"upload_time"` // 更新时间戳
}
