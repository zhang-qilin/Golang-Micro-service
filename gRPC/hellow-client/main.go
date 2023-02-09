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


// ClientTokenAuth 客户端定义的Token验证
type ClientTokenAuth struct {
}

// GetRequestMetadata
func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":"zhangqili",
		"appKey":"123456",
	},nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool{
	return false
}

func main() {
	// creds, _ := credentials.NewClientTLSFromFile("D:\\种子\\github.com\\Golang-Micro-service\\gRPC\\key\\test.pem", "*.zhangqilin.com") // *.zhangqilin.com这个域一般不会写死，一般都是通过url去获取的

	// 连接到server端，此处禁用安全传输，没有加密和验证
	// grpc.Dial("127.0.0.1:9090",grpc.WithTransportCredentials(grpc.WithInsecure()))  // grpc.WithInsecure()  已被遗弃，可使用以下的insecure.NewCredentials()则代表不使用安全加密
	// conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	// 使用SSL/TLS加密传输
	// conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	
	// token验证
	var opts []grpc.DialOption
	// 如果上面的RequireTransportSecurity返回true的话则需要把insecure.NewCredentials()替换成creds
	// opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	
	conn, err := grpc.Dial("127.0.0.1:9090",opts...)
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
