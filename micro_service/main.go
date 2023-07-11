package main

import (
	"context"
	"flag"
	"fmt"
	. "github.com/Sadio94/micro_service/internal/intialize"
	"github.com/Sadio94/micro_service/internal/models"
	"github.com/Sadio94/micro_service/internal/registry"
	"github.com/Sadio94/micro_service/proto"
	"github.com/Sadio94/micro_service/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var cfn string
	flag.StringVar(&cfn, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()

	// parse config
	configErr := Init(cfn)
	if configErr != nil {
		fmt.Println(configErr.Error())
		os.Exit(-1)
	}

	// init db
	modelErr := models.InitDb(Conf.MySQLConfig)
	if modelErr != nil {
		Logger.Error(modelErr.Error())
		os.Exit(-1)
	}

	// init consul
	initConsulErr := registry.Init(Conf.ConsulConfig.Addr)
	if initConsulErr != nil {
		Logger.Error(initConsulErr.Error())
		os.Exit(-1)
	}

	// listen
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Conf.IP, Conf.Port))
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(-1)
	}

	// grpc server
	s := grpc.NewServer()
	proto.RegisterGoodsServer(s, &service.GoodSvc{})
	go func() {
		startErr := s.Serve(lis)
		if startErr != nil {
			Logger.Error(startErr.Error())
			os.Exit(-1)
		}
	}()

	// register to consul
	go registry.Reg.RegisterService(Conf.Name, Conf.IP, Conf.Port, nil)

	Logger.Info("service start...")

	// gateway start
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext( // RPC客户端
		context.Background(),
		fmt.Sprintf("%s:%d", Conf.IP, Conf.Port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(-1)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = proto.RegisterGoodsHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: gwmux,
	}

	Logger.Info("Serving gRPC-Gateway on http://0.0.0.0:8091")
	go func() {
		gwListenErr := gwServer.ListenAndServe()
		if gwListenErr != nil {
			Logger.Error(gwListenErr.Error())
			os.Exit(-1)
		}
	}()

	// smooth exit
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// block until get message
	<-quit
	// 退出时注销服务
	serviceId := fmt.Sprintf("%s-%s-%d", Conf.Name, Conf.IP, Conf.Port)
	registry.Reg.Deregister(serviceId)
}
