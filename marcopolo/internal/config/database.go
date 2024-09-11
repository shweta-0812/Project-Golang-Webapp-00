package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// CloseDatabaseConnection closes the database connection pool
func CloseDatabaseConnection(db *gorm.DB) {
	// Retrieve the generic database object sql.DB to use its standard library functions
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get the generic database object: %v", err)
	}

	// Close the database connection
	err = dbSQL.Close()
	if err != nil {
		log.Fatalf("Failed to close the database connection: %v", err)
	}

	log.Println("Database connection closed successfully.")
}

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// data source name
	dsn := os.Getenv("POSTGRES_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// connect to db
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to get database: %w", err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	return DB
}
