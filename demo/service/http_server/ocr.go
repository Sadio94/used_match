/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  Ocr
 * @Date: 2023/08/11 16:37
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// OCR  @Summary OCR
// @Description 文件OCR识别
// @Tags         OCR
// @Param ocrRequest query rest.OCRRequest true "文件OCR识别参数"
// @Success 200 {object} rest.Result1{}
// @Router /edapi/bankbills/ocr [get]
func OCR(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{})
}
