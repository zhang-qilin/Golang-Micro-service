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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到server端，此处禁用安全传输，没有加密和验证
	// grpc.Dial("127.0.0.1:9090",grpc.WithTransportCredentials(grpc.WithInsecure()))  // grpc.WithInsecure()  已被遗弃，可使用以下的insecure.NewCredentials()则代表不使用安全加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Sprintf("Did Not Connect: %v\n", err)
		return
	}
	defer conn.Close()

	// 建立经历连接
	client := pb.NewSayHelloClient(conn)

	// 执行RPC调用(这个方法在服务端来实现并返回)
	response, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "张齐林"})
	fmt.Println(response.GetResponseMsg())
}
