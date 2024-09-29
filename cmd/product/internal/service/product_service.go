package service

import (
	"context"
	"my-project/cmd/product/internal/proto"
)

type ProductService struct {
	proto.UnimplementedProductServiceServer
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	// Implement the creation logic here
	return &proto.CreateProductResponse{}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	// Implement the get product logic here
	return &proto.GetProductResponse{}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*proto.ListProductsResponse, error) {
	// Implement the list products logic here
	return &proto.ListProductsResponse{}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.UpdateProductResponse, error) {
	// Implement the update product logic here
	return &proto.UpdateProductResponse{}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	// Implement the delete product logic here
	return &proto.DeleteProductResponse{}, nil
}