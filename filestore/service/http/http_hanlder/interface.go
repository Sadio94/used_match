package http_hanlder

import "github.com/gin-gonic/gin"

type IHandler interface {
	CorpList(context *gin.Context)
	CorpCreate(context *gin.Context)
	CorpEdit(context *gin.Context)
}

type HttpServer struct{}

func NewHttpServer() IHandler {
	return &HttpServer{}
}
