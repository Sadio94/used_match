/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  user_info
 * @Date: 2023/07/03 16:18
 */

package http_server

import (
	"demo/internal/initialize"
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// UserList  @Summary 用户
// @Description 用户列表
// @Param userList query rest.UserList true "用户列表过滤项"
// @Success 200 {object} rest.Result{data=rest.Paging}
// @Router /api/v1/users [get]
func UserList(context *gin.Context) {
	var userList rest.UserList
	// 这里增加翻译器
	if err := context.ShouldBind(&userList); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusBadRequest, rest.Result{
				Msg: err.Error(),
			})
			return
		}

		// 表单校验
		context.JSON(http.StatusBadRequest, rest.Result{
			Msg: errors.Translate(initialize.Trans),
		})
		return
	}

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.Paging{},
	})
}

// UserDetail  @Summary 用户
// @Description 用户详情
// @Param id path string true "用户id"
// @Success 200 {object} rest.Result{data=rest.UserInfo}
// @Router /api/v1/:id/user [get]
func UserDetail(context *gin.Context) {
	var idReq rest.IdRequest
	if err := context.ShouldBindUri(&idReq); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusBadRequest, rest.Result{
				Msg: err.Error(),
			})
			return
		}

		// 表单校验
		context.JSON(http.StatusBadRequest, rest.Result{
			Msg: errors.Translate(initialize.Trans),
		})
		return
	}

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.UserInfo{},
	})
}
