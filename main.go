
package main

import (
	"evaDemo/conf"
	"evaDemo/handler"
	"evaDemo/proto/evaDemo"
	"github.com/Gitforxuyang/eva/server"
	"google.golang.org/grpc"
)

func main(){
	server.Init()
	conf.Registry()
	server.RegisterGRpcService(func(server *grpc.Server) {
		evaDemo.RegisterEvaDemoServer(server,&handler.HandlerService{})
	},evaDemo.GetServerDesc())
	server.Run()
}
