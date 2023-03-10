/*
* @Time ： 2023-02-19 13:47
* @Auth ： 张齐林
* @File ：grpc_service.go
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

// gRPC服务端实现(无任何加密)
func main() {
	rpcServer := grpc.NewServer()
	
	
	service.RegisterProdServiceServer(rpcServer,service.ProductService)
	
	listener, err := net.Listen("tcp",":8002")
	if err != nil {
		log.Fatal("启动监听出错",err)
		return
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务出错",err)
		return
	}
	fmt.Println("启动RPC成功...")
}
