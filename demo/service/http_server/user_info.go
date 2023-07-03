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

func UserList(context *gin.Context) {
	var userList rest.UserList
	// 这里增加翻译器
	if err := context.ShouldBind(&userList); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 表单校验
		context.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Translate(initialize.Trans),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg": rest.Paging{},
	})
}

//func UserDetail(context *gin.Context) {
//
//}
