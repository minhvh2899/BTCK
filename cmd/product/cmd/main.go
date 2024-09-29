package main

import (
	"log"
	"my-project/cmd/product/internal/api/proto"
	"my-project/cmd/product/internal/config"
	"my-project/cmd/product/internal/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	lis, err := net.Listen("tcp", cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	productService := service.NewProductService()
	grpcServer := grpc.NewServer()
	proto.RegisterProductServiceServer(grpcServer, productService)

	log.Printf("Starting Product gRPC server on %s", cfg.ServerAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}