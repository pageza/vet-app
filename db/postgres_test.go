package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/pageza/vet-app/config"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255;unique"`
}

func setup(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	// Clear the database before each test
	ClearDB(DB)

	// Run migrations
	err = DB.AutoMigrate(&User{})
	assert.NoError(t, err)
}

func TestInitDB(t *testing.T) {
	setup(t)
	
}

func TestInitDBConnectionFailure(t *testing.T) {
	invalidConfig := config.DBConfig{
		Host:     "invalid_host",
		Port:     5432,
		User:     "invalid_user",
		Password: "invalid_password",
		Name:     "invalid_db",
	}

	err := InitDB(invalidConfig)
	assert.Error(t, err)
}

func TestDatabaseMigrations(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	// Example migration test
	type User struct {
		ID    uint   `gorm:"primaryKey"`
		Name  string `gorm:"size:255"`
		Email string `gorm:"size:255;unique"`
	}

	err = DB.AutoMigrate(&User{})
	assert.NoError(t, err)
}

func TestDataInsertionAndRetrieval(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	err = DB.AutoMigrate(&User{})
	assert.NoError(t, err)

	// Insert a user
	user := User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Retrieve the user
	var retrievedUser User
	result = DB.First(&retrievedUser, user.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, user.Name, retrievedUser.Name)
	assert.Equal(t, user.Email, retrievedUser.Email)
}

func TestDataUpdateAndDelete(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	type User struct {
		ID    uint   `gorm:"primaryKey"`
		Name  string `gorm:"size:255"`
		Email string `gorm:"size:255;unique"`
	}

	err = DB.AutoMigrate(&User{})
	assert.NoError(t, err)

	// Insert a user
	user := User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Update the user
	user.Name = "Jane Doe"
	result = DB.Save(&user)
	assert.NoError(t, result.Error)

	// Retrieve the updated user
	var updatedUser User
	result = DB.First(&updatedUser, user.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Jane Doe", updatedUser.Name)

	// Delete the user
	result = DB.Delete(&user)
	assert.NoError(t, result.Error)

	// Ensure the user is deleted
	var deletedUser User
	result = DB.First(&deletedUser, user.ID)
	assert.Error(t, result.Error)
	assert.Empty(t, deletedUser)
}

func TestDatabaseTransactions(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	type User struct {
		ID    uint   `gorm:"primaryKey"`
		Name  string `gorm:"size:255"`
		Email string `gorm:"size:255;unique"`
	}

	err = DB.AutoMigrate(&User{})
	assert.NoError(t, err)

	tx := DB.Begin()
	assert.NoError(t, tx.Error)

	// Insert a user within the transaction
	user := User{Name: "John Doe", Email: "john@example.com"}
	result := tx.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Rollback the transaction
	tx.Rollback()

	// Ensure the user was not committed to the database
	var rolledBackUser User
	result = DB.First(&rolledBackUser, user.ID)
	assert.Error(t, result.Error)
	assert.Empty(t, rolledBackUser)

	// Start a new transaction and commit
	tx = DB.Begin()
	assert.NoError(t, tx.Error)

	user = User{Name: "Jane Doe", Email: "jane@example.com"}
	result = tx.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	tx.Commit()

	// Ensure the user was committed to the database
	var committedUser User
	result = DB.First(&committedUser, user.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Jane Doe", committedUser.Name)
}

func TestDatabaseConstraints(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitDB(config.DB)
	assert.NoError(t, err)

	type User struct {
		ID    uint   `gorm:"primaryKey"`
		Name  string `gorm:"size:255"`
		Email string `gorm:"size:255;unique"`
	}

	err = DB.AutoMigrate(&User{})
	assert.NoError(t, err)

	// Insert a user
	user := User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Attempt to insert a user with a duplicate email
	duplicateUser := User{Name: "Jane Doe", Email: "john@example.com"}
	result = DB.Create(&duplicateUser)
	assert.Error(t, result.Error)
}