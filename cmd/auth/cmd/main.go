package main

import (
	"log"
	"my-project/cmd/auth/internal/config"
	"my-project/cmd/auth/internal/database"
	"my-project/cmd/auth/internal/proto"
	"my-project/cmd/auth/internal/repository"
	"my-project/cmd/auth/internal/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	lis, err := net.Listen("tcp", cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)

	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("Starting gRPC server on %s", cfg.ServerAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}