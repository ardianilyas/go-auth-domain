package routes

import (
	"github.com/ardianilyas/go-auth-domain/internal/auth"
	"github.com/ardianilyas/go-auth-domain/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	AuthHandler 	*auth.Authhandler
	AuthMiddleware 	*middlewares.AuthMiddleware
}

func SetupRoutes(r *gin.Engine, deps *Dependencies) {
	api := r.Group("/api")

	auth.RegisterAuthRoutes(api, deps.AuthHandler, deps.AuthMiddleware)
}