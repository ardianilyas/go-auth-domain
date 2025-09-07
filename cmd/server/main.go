package main

import (
	"github.com/ardianilyas/go-auth-domain/internal/bootstrap"
	"github.com/ardianilyas/go-auth-domain/internal/config"
	"github.com/ardianilyas/go-auth-domain/internal/database"
	"github.com/ardianilyas/go-auth-domain/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db := database.ConnectDB(cfg)

	deps := bootstrap.InitDependencies(db, cfg)

	r := gin.Default()

	routes.SetupRoutes(r, deps)

	r.Run(":8000")
}