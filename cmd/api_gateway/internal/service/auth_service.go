package service

import (
	"context"
	"my-project/cmd/api_gateway/internal/proto"

	"google.golang.org/grpc"
)

type AuthService struct {
	client proto.AuthServiceClient
}

func NewAuthService(conn *grpc.ClientConn) *AuthService {
	return &AuthService{
		client: proto.NewAuthServiceClient(conn),
	}
}

func (s *AuthService) Register(ctx context.Context, username, password, email string) (*proto.RegisterResponse, error) {
	return s.client.Register(ctx, &proto.RegisterRequest{
		Username: username,
		Password: password,
		Email:    email,
	})
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*proto.LoginResponse, error) {
	return s.client.Login(ctx, &proto.LoginRequest{
		Username: username,
		Password: password,
	})
}

func (s *AuthService) GetProfile(ctx context.Context, username string) (*proto.ProfileResponse, error) {
    return s.client.GetProfile(ctx, &proto.ProfileRequest{
        Username: username,
    })
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (*proto.ValidateTokenResponse, error) {
    return s.client.ValidateToken(ctx, &proto.ValidateTokenRequest{
        Token: token,
	})
}
