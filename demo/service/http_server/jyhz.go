package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JyhzBalanceFluctuation  @Summary 交易汇总
// @Description 交易汇总-余额波动
// @Param jyhzBalanceFluctuationRequest query rest.JyhzBalanceFluctuationRequest true "余额波动筛选项目"
// @Success 200 {object} rest.Result1{data=rest.BalanceFluctuationResp}
// @Router /edapi/bankbills/analyse/jyhz/balance/fluctuation [get]
func JyhzBalanceFluctuation(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.BalanceFluctuationResp{},
	})
}

// JyhzTransactionDistribution  @Summary 交易汇总
// @Description 交易汇总-交易分布
// @Param jyhzTransactionDistributionRequest query rest.JyhzTransactionDistributionRequest true "交易分布筛选项目"
// @Success 200 {object} rest.Result1{data=rest.TransactionDistribution}
// @Router /edapi/bankbills/analyse/jyhz/transaction/distribution [get]
func JyhzTransactionDistribution(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.TransactionDistribution{},
	})
}

// JyhzAbnormalTransactionHighFrequency  @Summary 交易汇总
// @Description 交易汇总-异常交易-高频交易
// @Param jydsAbnormalTransactionHighFrequencyRequest query rest.JydsAbnormalTransactionHighFrequencyRequest true "高频交易筛选项目"
// @Success 200 {object} rest.Result1{data=rest.AbnormalTransactionHighFrequencyResp}
// @Router /edapi/bankbills/analyse/jyhz/abnormal/transaction/high_frequency [get]
func JyhzAbnormalTransactionHighFrequency(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.AbnormalTransactionHighFrequencyResp{},
	})
}

// JyhzAbnormalTransactionSuspicious  @Summary 交易汇总
// @Description 交易汇总-异常交易-可疑交易
// @Param jydsAbnormalTransactionSuspiciousRequest query rest.JydsAbnormalTransactionSuspiciousRequest true "可疑交易筛选项目"
// @Success 200 {object} rest.Result1{data=rest.AbnormalTransactionSuspiciousResp}
// @Router /edapi/bankbills/analyse/jyhz/abnormal/transaction/suspicious [get]
func JyhzAbnormalTransactionSuspicious(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.AbnormalTransactionSuspiciousResp{},
	})
}

// JyhzAbnormalTransactionLarge  @Summary 交易汇总
// @Description 交易汇总-异常交易-大额交易
// @Param jydsAbnormalTransactionLargeRequest query rest.JydsAbnormalTransactionLargeRequest true "大额交易筛选项目"
// @Success 200 {object} rest.Result1{data=rest.AbnormalTransactionLargeResp}
// @Router /edapi/bankbills/analyse/jyhz/abnormal/transaction/large [get]
func JyhzAbnormalTransactionLarge(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.AbnormalTransactionLargeResp{},
	})
}
