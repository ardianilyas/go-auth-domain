package database

import (
	"log"

	"github.com/ardianilyas/go-auth-domain/internal/auth/models"
	"github.com/ardianilyas/go-auth-domain/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := cfg.DBDsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.RefreshToken{})

	log.Println("Connected to database")

	return db
}