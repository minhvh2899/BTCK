// package service

// import (
// 	"my-project/cmd/product/internal/models"
// 	"my-project/cmd/product/internal/repository"
// )

// type ProductService interface {
//     CreateProduct(title, desc string) (*models.Product, error)
//     GetProductByID(id uint) (*models.Product, error)
//     UpdateProduct(id uint, title, desc string) (*models.Product, error)
//     DeleteProduct(id uint) error
//     ListProducts() ([]models.Product, error)
// }

// type productService struct {
//     repo repository.ProductRepository
// }

// func NewProductService(repo repository.ProductRepository) ProductService {
//     return &productService{repo: repo}
// }

// func (s *productService) CreateProduct(title, desc string) (*models.Product, error) {
//     product := &models.Product{Title: title, Desc: desc}
//     return s.repo.Create(product)
// }

// func (s *productService) GetProductByID(id uint) (*models.Product, error) {
//     return s.repo.FindByID(id)
// }

// func (s *productService) UpdateProduct(id uint, title, desc string) (*models.Product, error) {
//     product, err := s.repo.FindByID(id)
//     if err != nil {
//         return nil, err
//     }
//     product.Title = title
//     product.Desc = desc
//     return s.repo.Update(product)
// }

// func (s *productService) DeleteProduct(id uint) error {
//     return s.repo.Delete(id)
// }

//	func (s *productService) ListProducts() ([]models.Product, error) {
//	    return s.repo.FindAll()
//	}
package service

import (
	"context"
	"fmt"
	"my-project/cmd/product/internal/api/proto"
	"my-project/cmd/product/internal/models"
	"my-project/cmd/product/internal/repository"
	"strconv"
)

type ProductService struct {
	proto.UnimplementedProductServiceServer
	repo repository.ProductRepository
    }

func NewProductService(repo repository.ProductRepository) proto.ProductServiceServer {
	return &ProductService{
        repo: repo,
    }
}

func (s *ProductService) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	product := &models.Product{
		Name:  req.Name,
		Description:  req.Description,
		Price: req.Price,
	}
    
	createdProduct, err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}
	
	return &proto.CreateProductResponse{Product: &proto.Product{
		Id: strconv.FormatUint(uint64(createdProduct.ID), 10),
		Name:  createdProduct.Name,
		Description:  createdProduct.Description,
		Price: createdProduct.Price,
	}}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	// Implement the get product logic here
	id, err := strconv.ParseUint(req.Id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	product, err := s.repo.FindByID(uint(id))
	if err != nil {
		return nil, err
	}
	return &proto.GetProductResponse{Product: &proto.Product{
		Id: strconv.FormatUint(uint64(product.ID), 10),
		Name:  product.Name,
		Description:  product.Description,
		Price: product.Price,
	}}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*proto.ListProductsResponse, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	protoProducts := make([]*proto.Product, len(products))
	for i, product := range products {
		protoProducts[i] = &proto.Product{
			Id: strconv.FormatUint(uint64(product.ID), 10),
			Name:  product.Name,
			Description:  product.Description,
			Price: product.Price,
		}
	}
	return &proto.ListProductsResponse{Products: protoProducts}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.UpdateProductResponse, error) {
	id, err := strconv.ParseUint(req.Id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	product, err := s.repo.FindByID(uint(id))
	if err != nil {
		return nil, err
	}
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price

	updatedProduct, err := s.repo.Update(product)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateProductResponse{
		Product: &proto.Product{
			Id: strconv.FormatUint(uint64(updatedProduct.ID), 10),
			Name:  updatedProduct.Name,
			Description:  updatedProduct.Description,
			Price: updatedProduct.Price,
		},
	}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	id, err := strconv.ParseUint(req.Id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	err = s.repo.Delete(uint(id))
	if err != nil {
		return nil, err
	}
	return &proto.DeleteProductResponse{
        Success: true,
    }, nil
}