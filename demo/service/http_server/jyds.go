/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  jyds
 * @Date: 2023/07/31 17:13
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JydsCounterpartyOverview  @Summary 交易对手
// @Description 交易对手-对手方概览
// @Param jydsOverviewRequest query rest.JydsOverviewRequest true "对手方概览筛选项目"
// @Success 200 {object} rest.Result1{data=rest.JydsOverviewResp}
// @Router /edapi/bankbills/analyse/jyds/counterparty/overview [get]
func JydsCounterpartyOverview(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.JydsOverviewResp{},
	})
}

// JydsCounterpartyClass  @Summary 交易对手
// @Description 交易对手-对手方分类
// @Param jydsClassRequest query rest.JydsClassRequest true "对手方分类筛选项目"
// @Success 200 {object} rest.Result1{data=rest.JydsClassResp}
// @Router /edapi/bankbills/analyse/jyds/counterparty/class [get]
func JydsCounterpartyClass(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.JydsClassResp{},
	})
}

// JydsMonitorObject  @Summary 交易对手
// @Description 交易对手-重点监测对象
// @Param jydsMonitorObjectRequest query rest.JydsMonitorObjectRequest true "重点监测对象筛选项目"
// @Success 200 {object} rest.Result1{data=rest.JydsMonitorObjectResp}
// @Router /edapi/bankbills/analyse/jyds/monitor_object [get]
func JydsMonitorObject(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.LsyzCAccountResp{},
	})
}
