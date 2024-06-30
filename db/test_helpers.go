package db

import (
	"log"
	"testing"

	"github.com/pageza/vet-app/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// ClearDB is a helper function to clear all tables in the database.
func ClearDB(db *gorm.DB) {
	tables := []string{"responses", "calls", "users"} // Clear in reverse order of dependencies
	for _, table := range tables {
		if err := db.Exec("TRUNCATE TABLE " + table + " CASCADE").Error; err != nil {
			log.Printf("Failed to clear table %s: %v", table, err)
		}
	}
}

// SetupDB is a helper function to initialize the database, run migrations, and clear the tables.
func SetupDB(t *testing.T, models ...interface{}) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	// Run migrations
	err = DB.AutoMigrate(models...)
	assert.NoError(t, err)

	// Clear the database before each test
	ClearDB(DB)
}
