/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  doc
 * @Date: 2023/08/09 15:12
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DocList  @Summary 文件管理
// @Description 文件管理-项目文件列表
// @Tags         文件管理
// @Param docListRequest query rest.DocListRequest true "获取文件列表参数"
// @Success 200 {object} rest.Result1{data=rest.DocListResp}
// @Router /edapi/bankbills/doc/query [get]
func DocList(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.DocListResp{},
	})
}

// UpdateDoc  @Summary 文件管理
// @Description 文件管理-修改某个文档
// @Tags         文件管理
// @Param updateDocRequest query rest.UpdateDocRequest true "编辑指定项目传参"
// @Success 200 {object} rest.Result1{}
// @Router /edapi/bankbills/doc/update [get]
func UpdateDoc(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{})
}

// DocDetail  @Summary 文件管理
// @Description 文件管理-文档详情
// @Tags         文件管理
// @Param docDetailRequest query rest.DocDetailRequest true "文档详情传参"
// @Success 200 {object} rest.Result1{data=rest.DocDetailResp}
// @Router /edapi/bankbills/doc/detail [get]
func DocDetail(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{
		Data: rest.DocDetailResp{},
	})
}

// DeleteDoc  @Summary 文件管理
// @Description 文件管理-删除指定文档
// @Tags         文件管理
// @Param deleteDocRequest query rest.DeleteDocRequest true "删除指定项目"
// @Success 200 {object} rest.Result1{}
// @Router /edapi/bankbills/doc/delete [get]
func DeleteDoc(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result1{})
}
