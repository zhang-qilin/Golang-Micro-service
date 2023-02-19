/*
* @Time ： 2023-02-19 13:50
* @Auth ： 张齐林
* @File ：product.go
* @IDE ：GoLand
 */
package service

import (
	"context"
)

var ProductService = &productService{}

type productService struct {
	// UnimplementedProdServiceServer

}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	// 实现具体的业务逻辑
	stock := p.GetStockById(request.ProdId)
	return &ProductResponse{ProdStock:stock},nil
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {

}

func (p *productService) GetStockById(id int32) int32 {
	return id
}
