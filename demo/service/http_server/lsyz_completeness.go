/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  lsyz_completeness
 * @Date: 2023/07/19 15:02
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LsyzCompletenessAccount  @Summary 流水验证
// @Description 流水验证-数据完整性-账号信息
// @Param lsyz query rest.Lsyz true "账号信息筛选项目"
// @Success 200 {object} rest.Result1{data=rest.LsyzCAccountResp}
// @Router /edapi/bankbills/analyse/lsyz/completeness/account [get]
func LsyzCompletenessAccount(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.LsyzCAccountResp{},
	})
}

// LsyzCompletenessBalance  @Summary 流水验证
// @Description 流水验证-数据完整性-余额完整
// @Param lsyzC query rest.LsyzC true "账号信息筛选项目"
// @Success 200 {object} rest.Result1{data=rest.LsyzBalanceResp}
// @Router /edapi/bankbills/analyse/lsyz/completeness/balance [get]
func LsyzCompletenessBalance(context *gin.Context) {
	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.LsyzBalanceResp{},
	})
}

// LsyzCompletenessSummary  @Summary 流水验证
// @Description 流水验证-数据完整性-摘要完整
// @Param lsyzC query rest.LsyzC true "账号信息筛选项目"
// @Success 200 {object} rest.Result1{data=rest.LsyzSummaryResp}
// @Router /edapi/bankbills/analyse/lsyz/completeness/summary [get]
func LsyzCompletenessSummary(context *gin.Context) {
	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.LsyzSummaryResp{},
	})
}
