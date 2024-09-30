package db

import (
	"fmt"
	"log"

	"github.com/jeauchter/adjutor-auth-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB initializes the database connection
func InitDB(config *config.Config) (db *gorm.DB, err error) {

	if config == nil {
		log.Fatal("Config is nil")
	}

	// Get the database connection components from environment variables
	dbUser := config.Database.User
	dbPassword := config.Database.Password
	dbHost := config.Database.Host
	dbPort := config.Database.Port
	dbName := config.Database.Name

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Database connection details are missing in environment variables")
	}

	// Build the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Initialize the database connection
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
