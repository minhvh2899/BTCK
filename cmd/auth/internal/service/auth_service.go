package service

import (
	"context"
	"fmt"
	"my-project/cmd/auth/internal/models"
	"my-project/cmd/auth/internal/proto"
	"my-project/cmd/auth/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var jwtKey = []byte("secret") // Replace with a secure secret key

// type AuthService interface {
// 	Register(username, password, email string) (*models.User, error)
// 	Login(username, password string) (string, error)
// 	ValidateToken(tokenString string) (*jwt.Token, error)
// 	GetUserByUsername(username string) (*models.User, error)
// }

type AuthService interface {
	Register(ctx context.Context, username, password, email string) (*proto.RegisterResponse, error)
	Login(ctx context.Context, username, password string) (*proto.LoginResponse, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserByUsername(username string) (*models.User, error)
	GetProfile(ctx context.Context, username string) (*proto.ProfileResponse, error)
}



type authService struct {
	proto.UnimplementedAuthServiceServer
	repo repository.UserRepository
}
func NewAuthService(repo repository.UserRepository) proto.AuthServiceServer {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	user, err := s.createUser(ctx, req)
	if err != nil {
		return &proto.RegisterResponse{Success: false, Message: err.Error()}, nil
	}
	return &proto.RegisterResponse{Success: true, Message: user.Email}, nil
}

func (s *authService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return &proto.LoginResponse{Success: false, Message: err.Error()}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &proto.LoginResponse{Success: false, Message: err.Error()}, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return &proto.LoginResponse{Success: false, Message: err.Error()}, nil
	}
	return &proto.LoginResponse{Success: true, Token: tokenString, Message: "Login successful"}, nil
}

func (s *authService) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
    token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })

    if err != nil {
        return &proto.ValidateTokenResponse{Valid: false}, nil
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        username, ok := claims["username"].(string)
        if !ok {
            return &proto.ValidateTokenResponse{Valid: false}, nil
        }
        return &proto.ValidateTokenResponse{Valid: true, Username: username}, nil
    }

    return &proto.ValidateTokenResponse{Valid: false}, nil
}

func (s *authService) GetUserByUsername(username string) (*models.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *authService) createUser(ctx context.Context, req *proto.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	return s.repo.Create(user)
}

func (s *authService) GetProfile(ctx context.Context, req *proto.ProfileRequest) (*proto.ProfileResponse, error) {
	user, err := s.GetUserByUsername(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v", err)
    }

    if user == nil {
        return nil, status.Errorf(codes.NotFound, "User not found")
    }

    return &proto.ProfileResponse{
        Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}
