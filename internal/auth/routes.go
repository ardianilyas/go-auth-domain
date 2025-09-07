package auth

import (
	"github.com/ardianilyas/go-auth-domain/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *Authhandler, middleware *middlewares.AuthMiddleware) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/refresh", handler.Refresh)
		auth.POST("/logout", handler.Logout)
		auth.GET("/me", middleware.RequireAuth(), handler.Me)
	}
}