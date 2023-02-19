/*
* @Time ： 2023-02-19 14:18
* @Auth ： 张齐林
* @File ：grpc_client.go
* @IDE ：GoLand
 */
package main

import (
	"context"
	"fmt"
	"log"
	"ms-proto/service"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// gRPC客户端实现(无任何加密)
func main() {
	// 建立连接，端口是服务端开放的端口(8002)
	conn, err := grpc.Dial(":8002",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("客户端连接不到服务端,",err)
		return
	}
	// 退出时关闭连接
	defer conn.Close()
	// 调用Product.pd.go中的NewProdServiceClient方法
	prodClient := service.NewProdServiceClient(conn)
	request := &service.ProductRequest{ProdId: 123456}
	// 直接调用本地方法一样调用GerProductStock方法
	stockResponse, err := prodClient.GetProductStock(context.Background(),request)
	if err != nil {
		log.Fatal("查询库存失败")
		return
	}
	fmt.Println("查询成功,",stockResponse.ProdStock)
}