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

// @host      http://192.168.6.50:10006
// @BasePath  /edapi/bankbills/
// @contact.email  yinjie_luo@intsig.net
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

	//docs.SwaggerInfo.Host = "http://192.168.6.50:10006"
	//docs.SwaggerInfo.BasePath = "/edapi/bankbills/"
	docs.SwaggerInfo.Title = "Bank Statement API"
	docs.SwaggerInfo.Description = "API Conversation For Bank Statement."

	// definition API
	v1 := r.Group("/edapi/bankbills/analyse")
	addLsyzRoutes(v1)
	addJyfxRoute(v1)
	addJydsRoutes(v1)
	addJyhzRoute(v1)
	v2 := r.Group("/edapi/bankbills/project")
	addProjectRoute(v2)
	v3 := r.Group("/edapi/bankbills/doc")
	addDocRoute(v3)
	v4 := r.Group("/edapi/bankbills/ocr")
	addOcrRoute(v4)
	r.GET("/edapi/bankbills/api/day/type", http_server.DayType)
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

func addJyfxRoute(v1 *gin.RouterGroup) {
	v1.GET("/jyfx/summary", http_server.JyfxSummary)
	v1.GET("/jyfx/top_user", http_server.JyfxTopUser)
	v1.GET("/jyfx/top_user/detail", http_server.JyfxTopUserDetail)
	v1.GET("/jyfx/top_summary", http_server.JyfxTopSummary)
}

func addProjectRoute(v *gin.RouterGroup) {
	v.GET("/add", http_server.AddProject)
	v.GET("/query", http_server.ProjectList)
	v.GET("/update", http_server.UpdateProject)
	v.GET("/delete", http_server.DeleteProject)
}

func addJydsRoutes(v1 *gin.RouterGroup) {
	v1.GET("/jyds/counterparty/overview", http_server.JydsCounterpartyOverview)
	v1.GET("/jyds/counterparty/class", http_server.JydsCounterpartyClass)
	v1.GET("/jyds/monitor_object", http_server.JydsMonitorObject)
}

func addJyhzRoute(v1 *gin.RouterGroup) {
	v1.GET("/jyhz/balance/fluctuation", http_server.JyhzBalanceFluctuation)
	v1.GET("/jyhz/transaction/distribution", http_server.JyhzTransactionDistribution)
	v1.GET("/jyhz/abnormal/transaction/high_frequency", http_server.JyhzAbnormalTransactionHighFrequency)
	v1.GET("/jyhz/abnormal/transaction/suspicious", http_server.JyhzAbnormalTransactionSuspicious)
	v1.GET("/jyhz/abnormal/transaction/large", http_server.JyhzAbnormalTransactionLarge)
}

func addDocRoute(v *gin.RouterGroup) {
	v.GET("/query", http_server.DocList)
	v.GET("/update", http_server.UpdateDoc)
	v.GET("/detail", http_server.DocDetail)
	v.GET("/delete", http_server.DeleteDoc)
	v.GET("/upload", http_server.UploadFile)
}

func addOcrRoute(v *gin.RouterGroup) {
	v.GET("", http_server.OCR)
}
