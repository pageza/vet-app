package db

import (
	"log"
	"testing"

	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/pageza/vet-app/config"
)

// ClearDB is a helper function to clear all tables in the database.
func ClearDB(db *gorm.DB) {
	tables := []string{"users"}
	for _, table := range tables {
		if err := db.Exec("DELETE FROM " + table).Error; err != nil {
			log.Fatalf("Failed to clear table %s: %v", table, err)
		}
	}
}

// SetupDB is a helper function to initialize the database and run migrations.
func SetupDB(t *testing.T, models ...interface{}) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	// Clear the database before each test
	ClearDB(DB)

	// Run migrations
	err = DB.AutoMigrate(models...)
	assert.NoError(t, err)
}