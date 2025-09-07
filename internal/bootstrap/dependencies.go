package bootstrap

import (
	"github.com/ardianilyas/go-auth-domain/internal/auth"
	"github.com/ardianilyas/go-auth-domain/internal/config"
	"github.com/ardianilyas/go-auth-domain/internal/middlewares"
	"github.com/ardianilyas/go-auth-domain/internal/routes"
	"gorm.io/gorm"
)

func InitDependencies(db *gorm.DB, cfg *config.Config) *routes.Dependencies {
	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo, cfg)
	authHandler := auth.NewAuthHandler(authService)
	authMiddleware := middlewares.NewAuthMiddleware(cfg.JWTSecret)

	return &routes.Dependencies{
		AuthHandler: 	authHandler,
		AuthMiddleware: authMiddleware,
	}
}