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
	server := &http.Server{Addr: ":10001", Handler: r}
	ln, err := net.Listen("tcp4", ":10001")
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.BasePath = "/edapi/bankbills/"
	v1 := r.Group("/edapi/bankbills/analyse")
	addLsyzRoutes(v1)
	v2 := r.Group("/edapi/bankbills/project")
	addProjectRoute(v2)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//r.Run(":8080")
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}
	server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}

//func addUserRoutes(v1 *gin.RouterGroup) {
//	v1.POST("/user/register", http_server.UserRegister)
//	v1.GET("/users", http_server.UserList)
//	v1.GET("/:id/user", http_server.UserDetail)
//}

func addLsyzRoutes(v1 *gin.RouterGroup) {
	v1.GET("/lsyz/authenticity", http_server.LsyzAuthenticity)
	v1.GET("/lsyz/completeness/account", http_server.LsyzCompletenessAccount)
	v1.GET("/lsyz/completeness/balance", http_server.LsyzCompletenessBalance)
	v1.GET("/lsyz/completeness/summary", http_server.LsyzCompletenessSummary)
}

//func addJyfxRoute(v *gin.RouterGroup) {
//
//}

func addProjectRoute(v *gin.RouterGroup) {
	v.GET("/add", http_server.AddProject)
	v.GET("/query", http_server.AddProject)
	v.GET("/update", http_server.AddProject)
}
