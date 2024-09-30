package main

import (
	"context"
	"log"
	"my-project/cmd/api_gateway/internal/api/handlers"
	"my-project/cmd/api_gateway/internal/api/middleware"
	"my-project/cmd/api_gateway/internal/config"
	"my-project/cmd/api_gateway/internal/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	conn, err := grpc.DialContext(context.Background(), cfg.AuthServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer conn.Close()

	authService := service.NewAuthService(conn)
	authHandler := handlers.NewAuthHandler(authService)
	authMiddleware := middleware.AuthMiddleware(authService)
	router := gin.Default()

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.GET("/profile", authMiddleware, authHandler.GetProfile)


	productConn, err := grpc.DialContext(context.Background(), cfg.ProductServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	productService := service.NewProductService(productConn)
	productHandler := handlers.NewProductHandler(productService)
    
	productRouter := router.Group("/products", authMiddleware)
	{
		productRouter.POST("", productHandler.CreateProduct)
		productRouter.GET("/:id", productHandler.GetProduct)
		productRouter.GET("", productHandler.ListProducts)
		productRouter.PUT("/:id", productHandler.UpdateProduct)
		productRouter.DELETE("/:id", productHandler.DeleteProduct)
	}

	log.Printf("Starting API Gateway on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}