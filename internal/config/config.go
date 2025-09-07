package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDsn 		string `env:"DB_DSN"`
	JWTSecret 	string `env:"JWT_SECRET"`
	JWTRefresh 	string `env:"JWT_REFRESH"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	cfg := &Config{
		DBDsn: 			os.Getenv("DB_DSN"),
		JWTSecret: 		os.Getenv("JWT_SECRET"),
		JWTRefresh: 	os.Getenv("JWT_REFRESH"),
	}

	return cfg
}