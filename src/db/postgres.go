package db

import (
	"fmt"
	"log"
	"os"

	"github.com/pageza/vet-app/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Automatically migrate the schema
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}

	// Manually create index if it does not exist
	result := DB.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);")
	if result.Error != nil {
		log.Fatalf("Error creating unique index on email: %v", result.Error)
	}
}
