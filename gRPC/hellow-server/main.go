/*
* @Time ： 2023-02-08 16:55
* @Auth ： 张齐林
* @File ：main.go
* @IDE ：GoLand
 */
package main

import (
	pb "Golang-Micro-service/gRPC/hellow-client/proto"
	"context"
	"fmt"
	"net"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

// SayHello 重写SayHello方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb. HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello"+req.RequestName},nil
}

func main() {
	// SSL/TSL加密验证
	creds, _ := credentials.NewServerTLSFromFile("D:\\种子\\github.com\\Golang-Micro-service\\gRPC\\key\\test.pem", "D:\\种子\\github.com\\Golang-Micro-service\\gRPC\\key\\test.key")
	
	// 开启端口
	listen, _ := net.Listen("tcp","127.0.0.1:9090")
	// 创建gRPC服务
	// grpcServer := grpc.NewServer
	
	// 添加SSL/TSL加密验证机制
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	
	// 在gRPC服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer,&server{})
	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("Failed To Server: %v\n\n", err)
		return
	}
	
}