//snowflake rpc server
package server

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tanus-co/snowflake/rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	IdWorkerServerPortKey = "snowflake.server-port"
)

//RPC Server
type Server struct {
	Port string
}

//创建新的RPC Server
func NewServer() *Server {
	port := viper.GetInt(IdWorkerServerPortKey)
	return &Server{
		Port: fmt.Sprintf(":%d", port),
	}
}

//开启RPC监听
func (s *Server) Listen() {
	//  创建server端监听端口
	list, err := net.Listen("tcp", s.Port)
	if err != nil {
		log.Panic(err.Error())
	}
	//  创建grpc的server
	server := grpc.NewServer()
	//  注册服务
	proto.RegisterGreeterServer(server, &SnowflakeServer{})
	//  启动grpc服务
	log.Println("the rpc service has started listen, on", s.Port)
	if err := server.Serve(list); err != nil {
		panic(err.Error())
	}
}
