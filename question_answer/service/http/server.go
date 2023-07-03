package http

import (
	"github.com/Sadio94/question_answer/service/http/http_handler"
	"github.com/gin-gonic/gin"
)

func GinStart(){
	r := initRoute()


	r.Run(":8080")
}


func initRoute() *gin.Engine{
	route := gin.New()
	v1 := route.Group("v1/api")
	addUserRoutes(v1)


	return route
}


func addUserRoutes(v1 *gin.RouterGroup) {
	v1.POST("/user/register", http_handler.UserRegister)
	v1.POST("/user/login", http_handler.UserLogin)
}
