/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  lsgl
 * @Date: 2023/08/16 14:27
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LsglList  @Summary 流水管理
// @Description 流水管理-流水详情列表
// @Tags 流水管理
// @Param lsgl query rest.LsglListRequest true "流水管理筛选项"
// @Success 200 {object} rest.Result1{data=rest.LsglListResp}
// @Router /edapi/bankbills/analyse/lsgl/list [get]
func LsglList(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.LsglListResp{},
	})
}

// LsglExport  @Summary 流水管理
// @Description 流水管理-导出
// @Tags 流水管理
// @Param lsglExport query rest.LsglExportRequest true "流水管理导出请求参数"
// @Success 200 {object} rest.Result1{}
// @Router /edapi/bankbills/analyse/lsgl/export [get]
func LsglExport(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{})
}
