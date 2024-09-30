package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// LoadConfig loads the environment variables from the .env file
func LoadConfig() (*Config, error) {
	config := Config{}
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Printf("Loaded .env file, value of DB_USER: %s", os.Getenv("DB_USER"))

	// Get the database connection components from environment variables
	DatabaseConfig := DatabaseConfig{}
	DatabaseConfig.User = os.Getenv("DB_USER")
	DatabaseConfig.Password = os.Getenv("DB_PASSWORD")
	DatabaseConfig.Host = os.Getenv("DB_HOST")
	DatabaseConfig.Port = os.Getenv("DB_PORT")
	DatabaseConfig.Name = os.Getenv("DB_NAME")

	// Set the database connection components in the Config struct

	config.Database = DatabaseConfig

	// Check if the .env file was loaded successfully
	log.Println("Loading .env file")

	return &config, nil
}
