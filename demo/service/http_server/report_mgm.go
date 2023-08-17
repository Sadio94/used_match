/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  report
 * @Date: 2023/08/17 10:17
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReportList  @Summary 分析报告
// @Description 分析报告-报告列表主页
// @Tags         分析报告
// @Param reportListRequest query rest.ReportListRequest true "获取主页列表参数"
// @Success 200 {object} rest.Result1{data=rest.ReportListResp}
// @Router /edapi/bankbills/report/query [get]
func ReportList(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.ReportListResp{},
	})
}

// ReportDetail  @Summary 分析报告
// @Description 分析报告-项目报告详情
// @Tags         分析报告
// @Param reportDetailRequest query rest.ReportDetailRequest true "项目报告详情传参"
// @Success 200 {object} rest.Result1{data=rest.ReportDetailResp}
// @Router /edapi/bankbills/report/detail [get]
func ReportDetail(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.ReportDetailResp{},
	})
}

// DeleteReport  @Summary 分析报告
// @Description 分析报告-删除指定报告
// @Tags         分析报告
// @Param deleteReportRequest query rest.DeleteReportRequest true "删除指定报告"
// @Success 200 {object} rest.Result1{}
// @Router /edapi/bankbills/report/delete [get]
func DeleteReport(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{})
}
