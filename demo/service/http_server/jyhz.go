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

// JyhzAbnormalTransaction  @Summary 交易汇总
// @Description 交易汇总-异常交易
// @Param jydsAbnormalTransactionRequest query rest.JydsAbnormalTransactionRequest true "异常交易筛选项目"
// @Success 200 {object} rest.Result1{data=rest.AbnormalTransactionResp}
// @Router /edapi/bankbills/analyse/jyhz/abnormal/transaction [get]
func JyhzAbnormalTransaction(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.AbnormalTransactionResp{},
	})
}
