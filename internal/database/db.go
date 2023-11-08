// internal/database/db.go

package database

import (
	"log"
	"os"

	"github.com/pageza/vet-app/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
