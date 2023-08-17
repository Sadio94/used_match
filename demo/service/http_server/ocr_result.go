/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  ocr_result
 * @Date: 2023/08/17 15:51
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResultQuery  @Summary OCR识别结果
// @Description 文件OCR识别结果
// @Tags         OCR识别结果
// @Param ocrResultQueryRequest query rest.OCRResultQueryRequest true "文件OCR识别结果请求参数"
// @Success 200 {object} rest.Result1{data=rest.OCRResultQueryResp}
// @Router /edapi/bankbills/result/query [get]
func ResultQuery(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.OCRResultQueryResp{},
	})
}

// UpdateHeader  @Summary OCR识别结果
// @Description 表头错误处理更新
// @Tags         OCR识别结果
// @Param updateHeaderQuery query rest.UpdateHeaderQuery true "表头错误处理更新请求参数"
// @Param updateHeaderBody  body rest.UpdateHeaderBody true "表头错误处理更新请求参数"
// @Success 200 {object} rest.Result1{data=rest.UpdateHeaderResp}
// @Router /edapi/bankbills/result/update_header [post]
func UpdateHeader(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.UpdateHeaderResp{},
	})
}
