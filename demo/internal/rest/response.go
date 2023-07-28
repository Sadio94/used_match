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
