/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  user_register
 * @Date: 2023/07/03 16:17
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister @Summary 用户
// @Description 用户注册
// @Param registerInfo body rest.RegisterInfo true "用户注册基本信息"
// @Accept json
// @Produce json
// @Success 200 {object} json "{"msg": ""}"
// @Router /user/register [post]
func UserRegister(context *gin.Context) {
	var registerInfo rest.RegisterInfo
	// 这里是最简单的bind标签的使用
	if err := context.ShouldBind(&registerInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
