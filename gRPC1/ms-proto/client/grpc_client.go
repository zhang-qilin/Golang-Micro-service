/*
* @Time ： 2023-02-12 16:47
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

func main() {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("客户端连接不是服务端，", err)
		return
	}
	defer conn.Close()
	prodClient := service.NewProdServiceClient(conn)
	request := &service.ProductRequest{
		ProdId: 1234,
	}
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错", err)
		return
	}
	fmt.Println("查询成功",stockResponse.ProdStock)

}
