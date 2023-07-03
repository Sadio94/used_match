/**
* @Author: Sadio94
* @Description:
* @File:  main
* @Date: 2023/06/30 19:15
 */

package main

import (
	"demo/internal/initialize"
	"demo/internal/rest"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	// param locale by config
	if err := initialize.TransInitialize("zh"); err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	v1 := r.Group("v1/api")
	addUserRoutes(v1)
	r.Run(":8080")
}

func addUserRoutes(v1 *gin.RouterGroup) {
	v1.POST("/user/register", func(context *gin.Context) {
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
	})

	v1.GET("/users", func(context *gin.Context) {
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
	})
}
