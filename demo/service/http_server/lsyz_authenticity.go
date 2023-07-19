/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  lsyz_authenticity
 * @Date: 2023/07/19 14:59
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LsyzAuthenticity  @Summary 流水验证
// @Description 流水验证-数据真实性
// @Param lsyz query rest.Lsyz true "数据真实性报表筛选项"
// @Success 200 {object} rest.Result1{data=rest.LsyzAResp}
// @Router /edapi/bankbills/analyse/lsyz/authenticity [get]
func LsyzAuthenticity(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.LsyzAResp{},
	})
}
