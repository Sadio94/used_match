/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  gljy
 * @Date: 2023/08/16 17:05
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GljyRelatedOverview  @Summary 关联交易
// @Description 关联交易-关联方概览和关联方交易概览
// @Tags 关联交易
// @Param gljyRelatedOverviewRequest query rest.GljyRelatedOverviewRequest true "关联方概览筛选项"
// @Success 200 {object} rest.Result1{data=rest.GljyRelatedOverviewResp}
// @Router /edapi/bankbills/analyse/gljy/related_party/overview [get]
func GljyRelatedOverview(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.GljyRelatedOverviewResp{},
	})
}

// GljyRelatedInformation  @Summary 关联交易
// @Description 关联交易-关联方信息
// @Tags 关联交易
// @Param gljyRelatedInformationRequest query rest.GljyRelatedInformationRequest true "关联方信息筛选项"
// @Success 200 {object} rest.Result1{data=rest.GljyRelatedInformationResp}
// @Router /edapi/bankbills/analyse/gljy/related_party/information [get]
func GljyRelatedInformation(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.GljyRelatedInformationResp{},
	})
}

// GljyRelatedDetail  @Summary 关联交易
// @Description 关联交易-关联交易详情
// @Tags 关联交易
// @Param gljyRelatedDetailRequest query rest.GljyRelatedDetailRequest true "关联交易详情筛选项"
// @Success 200 {object} rest.Result1{data=rest.GljyRelatedDetailResp}
// @Router /edapi/bankbills/analyse/gljy/related_party/transaction/detail [get]
func GljyRelatedDetail(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.GljyRelatedDetailResp{},
	})
}

// GljyRelatedTransaction  @Summary 关联交易
// @Description 关联交易-关联交易
// @Tags 关联交易
// @Param gljyRelatedTransactionRequest query rest.GljyRelatedTransactionRequest true "关联交易筛选项"
// @Success 200 {object} rest.Result1{data=rest.GljyRelatedTransactionResp}
// @Router /edapi/bankbills/analyse/gljy/related_party/transaction [get]
func GljyRelatedTransaction(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.GljyRelatedTransactionResp{},
	})
}
