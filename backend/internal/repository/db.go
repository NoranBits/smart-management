// ////////////////////////////////////////////////////////////////////////////
// src: ./internal/repository/db.go											//
// desc: Establishes and returns a new GORM DB connection using PostgreSQL.//
// /////////////////////////////////////////////////////////////////////////
package repository

import (
	"log"

	config "backend_server/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB initializes and returns a new GORM DB connection.
func NewDB(cfg *config.Config) *gorm.DB {
	dsn := cfg.DSN
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set or is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}
