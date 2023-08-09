/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  jyfx
 * @Date: 2023/08/08 18:26
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JyfxSummary  @Summary 经营分析
// @Description 经营分析-经营概览
// @Tags 经营分析
// @Param jyfx query rest.Lsyz true "经营概览筛选项"
// @Success 200 {object} rest.Result1{data=rest.JyfxSummaryResp}
// @Router /edapi/bankbills/analyse/jyfx/summary [get]
func JyfxSummary(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.JyfxSummaryResp{},
	})
}

// JyfxTopUser  @Summary 经营分析
// @Description 经营分析-交易集中度
// @Tags 经营分析
// @Param jyfxTopUser query rest.JyfxTopUserRequest true "交易集中度筛选项"
// @Success 200 {object} rest.Result1{data=rest.JyfxTopUserResp}
// @Router /edapi/bankbills/analyse/jyfx/top_user [get]
func JyfxTopUser(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.JyfxTopUserResp{},
	})
}

// JyfxTopUserDetail  @Summary 经营分析
// @Description 经营分析-交易集中度详情
// @Tags 经营分析
// @Param jyfxTopUserDetail query rest.JyfxTopUserDetailRequest true "交易集中度详情筛选项"
// @Success 200 {object} rest.Result1{data=rest.JyfxTopUserDetailResp}
// @Router /edapi/bankbills/analyse/jyfx/top_user/detail [get]
func JyfxTopUserDetail(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.JyfxTopUserDetailResp{},
	})
}

// JyfxTopSummary  @Summary 经营分析
// @Description 经营分析-流水构成
// @Tags 经营分析
// @Param jyfxTopSummary query rest.JyfxTopSummaryRequest true "流水构成筛选项"
// @Success 200 {object} rest.Result1{data=rest.JyfxTopSummaryResp}
// @Router /edapi/bankbills/analyse/jyfx/top_summary [get]
func JyfxTopSummary(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.JyfxTopSummaryResp{},
	})
}
