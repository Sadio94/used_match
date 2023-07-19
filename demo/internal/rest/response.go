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
	HasBalanceR   []string         `json:"has_balance_r"`  // 哪几个月有余额字段且余额正确 ["2022/1", "2022/5"]
	HasBalanceE   []string         `json:"has_balance_e"`  // 哪几个月有余额字段但余额错误 ["2022/3", "2022/4"]
	NoneBalance   []string         `json:"none_balance"`   // 那几个月无流水["2022/2"]
	AbnormalTrade []AbnormalTradeI `json:"abnormal_trade"` // 异常记录

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
	HasSummary    []string          `json:"has_summary"`    // 哪几个月有摘要 ["2022/1", "2022/5"]
	NoneSummary   []string          `json:"none_summary"`   // 哪几个月没有摘要 ["2022/2"]
	AbnormalTrade []AbnormalTradeI1 `json:"abnormal_trade"` // 异常记录
}

type AbnormalTradeI1 struct {
	TradeTime     string `json:"trade_time"`     // 交易时间
	TradeAmount   int64  `json:"trade_amount"`   // 交易金额
	EndingBalance int64  `json:"ending_balance"` // 期末余额
	Summary       string `json:"summary"`        // 交易摘要
	Name          string `json:"name"`           // 交易户名
	Account       string `json:"account"`        // 交易账号
}
