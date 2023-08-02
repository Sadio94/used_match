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
	Count       int64          `json:"count"`        // 项目总数
	DocCount    int64          `json:"doc_count"`    // 文档总数
	ReportCount int64          `json:"report_count"` // 报告总数
	AmountCount int64          `json:"amount_count"` // 金额总数
	ProjectList []ProjectListI `json:"project_list"` // 项目详情
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
}

// BalanceFluctuationResp 交易汇总-余额波动
type BalanceFluctuationResp struct {
	BalanceFluctuation []BalanceInfo `json:"balance_fluctuation"` // 余额波动
}

type BalanceInfo struct {
	Timer      string `json:"timer"`       // 月 2022-01
	AvgBalance int    `json:"avg_balance"` // 月均余额
	Volatility string `json:"volatility"`  // 波动率 20%
}

// TransactionDistribution 交易汇总-交易分布
type TransactionDistribution struct {
	TradeAmount []TradeAmountI `json:"trade_amount"` // 交易金额
	TradeNumber []TradeNumberI `json:"trade_number"` // 交易笔数
}

type TradeAmountI struct {
	Timer  string `json:"timer"`  // 月 2022-01
	Income int64  `json:"income"` // 当月收入
	Pay    int64  `json:"pay"`    // 当月支出
	Total  int64  `json:"total"`  // 当月总流水
}

type TradeNumberI struct {
	Timer     string `json:"timer"`      // 月 2022-01
	IncomeNum int64  `json:"income_num"` // 当月收入笔数
	PayNum    int64  `json:"pay_num"`    // 当月支出笔数
	Total     int64  `json:"total"`      // 当月总交易笔数
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
	Time    string `json:"time"`    // 时间
	Amount  int64  `json:"amount"`  // 金额
	Name    string `json:"name"`    // 交易户名
	Account string `json:"account"` // 交易账号
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
