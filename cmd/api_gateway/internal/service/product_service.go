package service

import (
	"context"
	"my-project/cmd/api_gateway/internal/proto"

	"google.golang.org/grpc"
)

type ProductService struct {
    client proto.ProductServiceClient
}

func NewProductService(conn *grpc.ClientConn) *ProductService {
    return &ProductService{
        client: proto.NewProductServiceClient(conn),
    }
}

func (s *ProductService) CreateProduct(ctx context.Context, name, description string, price float64) (*proto.CreateProductResponse, error) {
    return s.client.CreateProduct(ctx, &proto.CreateProductRequest{
        Name:        name,
        Description: description,
        Price:       price,
    })
}

func (s *ProductService) GetProduct(ctx context.Context, id string) (*proto.GetProductResponse, error) {
    return s.client.GetProduct(ctx, &proto.GetProductRequest{
        Id: id,
    })
}

func (s *ProductService) ListProducts(ctx context.Context, page, limit int32) (*proto.ListProductsResponse, error) {
    return s.client.ListProducts(ctx, &proto.ListProductsRequest{
        Page:  page,
        Limit: limit,
    })
}

func (s *ProductService) UpdateProduct(ctx context.Context, id, name, description string, price float64) (*proto.UpdateProductResponse, error) {
    return s.client.UpdateProduct(ctx, &proto.UpdateProductRequest{
        Id:          id,
        Name:        name,
        Description: description,
        Price:       price,
    })
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) (*proto.DeleteProductResponse, error) {
    return s.client.DeleteProduct(ctx, &proto.DeleteProductRequest{
        Id: id,
    })
}