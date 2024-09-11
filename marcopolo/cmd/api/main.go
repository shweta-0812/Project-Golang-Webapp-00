package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Database driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // File source driver
	"github.com/joho/godotenv"
	"log"
	"marcopolo/internal/config"
	"marcopolo/internal/routes"
	"os"
)

func RunMigrations() {
	// Define the database URL and source URL

	dbUrl := os.Getenv("POSTGRES_DB_URL")
	//fmt.Println(dbUrl)

	// Create a new migrate instance with the correct file URL scheme
	m, err := migrate.New(
		"file://./database/migrations", // Correctly formatted URL with 'file://' scheme
		dbUrl,
	)

	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}

}

//func FixAndRunMigrations() {
//	// Define the database URL and source URL
//
//	dbUrl := os.Getenv("POSTGRES_DB_URL")
//	//fmt.Println(dbUrl)
//
//	// Create a new migrate instance with the correct file URL scheme
//	m, err := migrate.New(
//		"file://./database/migrations", // Correctly formatted URL with 'file://' scheme
//		dbUrl,
//	)
//
//	if err != nil {
//		log.Fatalf("Failed to create migrate instance: %v", err)
//	}
//
//	// Run the migrations
//	if err = m.Up(); err != nil {
//		if strings.Contains(err.Error(), "Dirty database") {
//			log.Println("Database is in a dirty state. Attempting to clean...")
//
//			//  Roll Back the Failed Migration
//			if err := m.Steps(-1); err != nil { // Roll back 1 step
//				log.Println("Failed to roll back migration. Attempting to force version")
//
//				// Force the Version to the Last Successful Migration
//				if forceErr := m.Force(1); forceErr != nil {
//					log.Fatalf("Failed to force migration version: %v", forceErr)
//				}
//			}
//
//			// Retry running migrations
//			log.Println("Retrying migrations...")
//			if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
//				log.Fatalf("Failed to rerun migrations: %v", err)
//			}
//		} else if !errors.Is(err, migrate.ErrNoChange) {
//			log.Fatalf("Failed to run migrations: %v. Fix the errors and retry.", err)
//		}
//	} else {
//		log.Println("Migrations applied successfully!")
//	}
//
//}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db := config.SetupDatabaseConnection()

	// Ensure the database connection is closed when the main function returns
	defer config.CloseDatabaseConnection(db)

	// Run database migrations
	RunMigrations()

	// Set up Gin router
	router := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(router)

	// Start the server
	log.Fatal(router.Run(":8080"))

}
