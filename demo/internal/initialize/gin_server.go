package initialize

//import (
//	"crypto/tls"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	swaggerFiles "github.com/swaggo/files"
//	ginSwagger "github.com/swaggo/gin-swagger"
//	. "gitlab2.aishu.cn/Zeus/service/hypersla/sla-mgm/internal/initials"
//	"gitlab2.aishu.cn/Zeus/service/hypersla/sla-mgm/service/http_server/http_handler"
//	tls2 "gitlab2.aishu.cn/flora/greed/tls"
//	"net"
//	"net/http_server"
//	"time"
//)
//
//var handler = http_handler.NewHttpServer()
//
//func GinStart() {
//	// 初始化请求路由
//	route := initRoute()
//
//	cfg := new(tls.Config)
//	// 配置tls
//	if cert, err := tls2.LoadX509KeyPair(Env.HttpsSSLcCfg.ServerCrtFile, Env.HttpsSSLcCfg.ServerKeyFile, Env.HttpsSSLcCfg.ServerCrtPwd); err == nil {
//		cfg.Certificates = append(cfg.Certificates, cert)
//
//	} else {
//		Logger.Errorf("【web-server】LoadX509KeyPair err %v", err)
//	}
//	cfg.InsecureSkipVerify = true
//	cfg.MinVersion = tls.VersionTLS12
//	cfg.MaxVersion = tls.VersionTLS13
//	cfg.CipherSuites = CipherSuites
//	// 设置服务启动参数
//	addr := fmt.Sprintf("%s:%d", AppCfg.Http.Host, AppCfg.Http.Port)
//	server := &http_server.Server{
//		Addr:           addr,
//		TLSConfig:      cfg,
//		Handler:        route,
//		ReadTimeout:    time.Duration(AppCfg.Http.HttpTimeOut) * time.Second,
//		WriteTimeout:   time.Duration(AppCfg.Http.HttpTimeOut) * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//	//启动https服务
//	ln, err := net.Listen("tcp", addr)
//	if err != nil {
//		Logger.Errorf("【web-server】net.Listen %s", err)
//	}
//	listen := tls.NewListener(ln, cfg)
//	err = server.Serve(listen)
//	if err != nil {
//		Logger.Errorf("【web-server】running %s", err)
//	}
//}
//
//func initRoute() *gin.Engine {
//	route := gin.New()
//	route.Use(MW.Inform())
//
//	// 用户鉴权
//	route.Use(MW.JWTWebAuth())
//	port, _ := Env.ServerIni.Int("WebService::wsport")
//	route.Use(MW.Referer(Env.VIP, Env.VIPEx, port, INSTALLDIR))
//
//	// 超时
//	route.Use(MW.TimeoutMiddleware(time.Duration(AppCfg.Http.HttpTimeOut) * time.Second))
//	// 每秒最大请求次数限制
//	route.Use(MW.LimitMiddleware(AppCfg.Http.RateLimit))
//
//	v1 := route.Group("/api/sla/v1")
//	// 注册接口文档路由
//	v1.GET("/docs/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//
//	// 注册路由
//	addGroupRoutes(v1)
//	addCommonRoutes(v1)
//
//	return route
//}
//
//func addGroupRoutes(r *gin.RouterGroup) {
//	// SLA
//	r.POST("/group", handler.AddSLA)                         // 新建SLA
//	r.PUT("/group/:id/base_info", handler.UpdateSLABaseInfo) // 编辑SLA基本信息
//	r.PUT("/group/:id/backup", handler.UpdateSLABaseBackup)  // 编辑SLA 备份策略配置
//	r.GET("/group/:id", handler.GetSLADetailById)            // SLA详情
//	r.GET("/groups", handler.GetSLAList)                     // SLA列表
//	r.DELETE("/groups", handler.BatchDeleteSLA)              // 批量删除SLA
//	//r.POST("/group/save", handler.AddSLAForResource)         // 保护对象下编辑SLA 本质上是另存为
//}
//
//func addCommonRoutes(r *gin.RouterGroup) {
//	// Common
//	r.PUT("/common/modify_status", handler.BatchUpdateSLAStatus)  // 批量启用/禁用SLA
//	r.GET("/common/name/exists", handler.SLANameExists)           // SLA名校验
//	r.GET("/common/:id/bind_objects", handler.GetBindObjectsById) // 单个SLA下绑定的保护对象
//	r.GET("/common/has_bind", handler.HasBind)                    // 显示已绑定保护对象的SLA名数组 主要用于前端页面弹窗
//	r.POST("/common/remove_objects", handler.SLARemoveObjects)
//}
//
////func (this *HttpServer) addStrategyRoutes(r *gin.RouterGroup) {
//// strategy
////}
