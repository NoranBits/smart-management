package main

import (
	"log"
	"net/http"

	config "backend_server/internal/config"
	model "backend_server/internal/model"
	repository "backend_server/internal/repository"
	router "backend_server/internal/router"
	logger "backend_server/pkg/logger"

	"github.com/joho/godotenv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Load configuration from environment variables or a file.
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	// cfg now contains values from .env
	log.Println("Configuration loaded:", cfg)

	// Initialize custom logger.
	logger.Init(cfg.LogLevel)

	// Initialize the database connection using GORM.
	db := repository.NewDB(cfg)

	// Optionally run migrations.
	db.AutoMigrate(&model.User{})

	// Create repository instance for additional repository methods.
	repo := repository.NewRepository(db)

	// Initialize the router, passing in any required dependencies (e.g., repo).
	r := router.NewRouter(repo)

	// Start the HTTP server.
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
