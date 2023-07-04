/**
* @Author: Sadio94
* @Description:
* @File:  main
* @Date: 2023/06/30 19:15
 */

package main

import (
	"demo/docs"
	"demo/internal/initialize"
	"demo/service/http_server"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net"
	"net/http"
)

func main() {
	// param locale by config
	if err := initialize.TransInitialize("zh"); err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	// 强制ipv4监听
	server := &http.Server{Addr: ":8080", Handler: r}
	ln, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("v1/api")
	addUserRoutes(v1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//r.Run(":8080")
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}
	server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}

func addUserRoutes(v1 *gin.RouterGroup) {
	v1.POST("/user/register", http_server.UserRegister)
	v1.GET("/users", http_server.UserList)
	v1.GET("/:id/user", http_server.UserDetail)
}
