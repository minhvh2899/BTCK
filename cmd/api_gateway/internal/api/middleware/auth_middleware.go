package middleware

import (
	"my-project/cmd/api_gateway/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 {
            c.JSON(401, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        token := bearerToken[1]
        claims, err := authService.ValidateToken(c.Request.Context(), token)
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        username := claims.Username
        if username == "" {
            c.JSON(401, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        c.Set("username", username)
        c.Next()
    }
}