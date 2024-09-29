package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"my-project/cmd/auth/internal/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(c.Request.Context(), input.Username, input.Password, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(c.Request.Context(), input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
    // Get the token from the context (set by the AuthMiddleware)
    tokenInterface, exists := c.Get("token")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
        return
    }

    token, ok := tokenInterface.(*jwt.Token)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims"})
        return
    }

    username, ok := claims["username"].(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username in token"})
        return
    }

    // Fetch user profile from the database
    user, err := h.service.GetUserByUsername(username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
        return
    }

    // Return the user profile (excluding sensitive information like password)
    c.JSON(http.StatusOK, gin.H{
        "id":       user.ID,
        "username": user.Username,
        "email":    user.Email,
    })
}