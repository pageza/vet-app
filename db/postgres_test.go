package db

import (
	"fmt"
	"sync"
	"testing"

	"github.com/pageza/vet-app/config"
	"github.com/pageza/vet-app/models"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) {
	SetupDB(t, &models.User{})
}

func TestInitDB(t *testing.T) {
	setup(t)
	// Additional assertions or checks can be added here if needed
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
	setup(t)
	// No specific assertions needed here as setup runs AutoMigrate
}

func TestDataInsertionAndRetrieval(t *testing.T) {
	setup(t)

	// Insert a user
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Retrieve the user
	var retrievedUser models.User
	result = DB.First(&retrievedUser, user.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, user.Name, retrievedUser.Name)
	assert.Equal(t, user.Email, retrievedUser.Email)
}

func TestDataUpdateAndDelete(t *testing.T) {
	setup(t)

	// Insert a user
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Update the user
	user.Name = "Jane Doe"
	result = DB.Save(&user)
	assert.NoError(t, result.Error)

	// Retrieve the updated user
	var updatedUser models.User
	result = DB.First(&updatedUser, user.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Jane Doe", updatedUser.Name)

	// Delete the user
	result = DB.Delete(&user)
	assert.NoError(t, result.Error)

	// Ensure the user is deleted
	var deletedUser models.User
	result = DB.First(&deletedUser, user.ID)
	assert.Error(t, result.Error)
	assert.Empty(t, deletedUser)
}

func TestDatabaseTransactions(t *testing.T) {
	setup(t)

	tx := DB.Begin()
	assert.NoError(t, tx.Error)

	// Insert a user within the transaction
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	result := tx.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Rollback the transaction
	tx.Rollback()

	// Ensure the user was not committed to the database
	var rolledBackUser models.User
	result = DB.First(&rolledBackUser, user.ID)
	assert.Error(t, result.Error)
	assert.Empty(t, rolledBackUser)

	// Start a new transaction and commit
	tx = DB.Begin()
	assert.NoError(t, tx.Error)

	user = models.User{Name: "Jane Doe", Email: "jane@example.com"}
	result = tx.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	tx.Commit()

	// Ensure the user was committed to the database
	var committedUser models.User
	result = DB.First(&committedUser, user.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Jane Doe", committedUser.Name)
}

func TestDatabaseConstraints(t *testing.T) {
	setup(t)

	// Insert a user
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Attempt to insert a user with a duplicate email
	duplicateUser := models.User{Name: "Jane Doe", Email: "john@example.com"}
	result = DB.Create(&duplicateUser)
	assert.Error(t, result.Error)
}

func TestConcurrentPostgresAccess(t *testing.T) {
	setup(t)

	user := models.User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	var wg sync.WaitGroup
	const numGoroutines = 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				// Simulate read operation
				var retrievedUser models.User
				DB.First(&retrievedUser, user.ID)
				assert.Equal(t, user.Email, retrievedUser.Email)

				// Simulate write operation
				retrievedUser.Name = fmt.Sprintf("John Doe %d-%d", i, j)
				DB.Save(&retrievedUser)
			}
		}(i)
	}

	wg.Wait()

	// Verify final state
	var finalUser models.User
	DB.First(&finalUser, user.ID)
	assert.Contains(t, finalUser.Name, "John Doe")
}
