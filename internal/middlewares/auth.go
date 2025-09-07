package middlewares

import (
	"strings"

	"github.com/ardianilyas/go-auth-domain/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: secret,
	}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil || strings.TrimSpace(accessToken) == "" {
			utils.Unauthorized(c, "access token missing")
			c.Abort()
			return 
		}

		token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(m.jwtSecret), nil
		})
		if err != nil || !token.Valid {
			utils.Unauthorized(c, "invalid or expired access token")
			c.Abort()
			return 
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.Unauthorized(c, "invalid token claims")
			c.Abort()
			return 
		}

		if uid, ok := claims["user_id"].(string); ok {
			c.Set("user_id", uid)
		}
		if role, ok := claims["role"].(string); ok {
			c.Set("role", role)
		}
		
		c.Next()
	}
}