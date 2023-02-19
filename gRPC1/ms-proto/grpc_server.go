/*
* @Time ： 2023-02-12 14:34
* @Auth ： 张齐林
* @File ：grpc_server.go
* @IDE ：GoLand
 */
package main

import (
	"fmt"
	"log"
	"ms-proto/service"
	"net"
	
	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	
	service.RegisterProdServiceServer(rpcServer,service.ProductService)
	
	listener, err := net.Listen("tcp",":8002")
	if err != nil {
		log.Fatal("启动监听出错：",err )
		return
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务出错：",err )
		return
	}
	fmt.Println("启动gRPC服务成功")
}
