package database

import (
	"github.com/pageza/vet-app/internal/user"
	"gorm.io/gorm"
)

// DBUserRepository is the database implementation of the user.UserRepository interface.
type DBUserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of a user repository.
func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &DBUserRepository{db: db}
}

// FindByID finds a user by their ID.
func (r *DBUserRepository) FindByID(id int) (*user.User, error) {
	// Database logic to find a user by ID.
	var u user.User
	result := r.db.First(&u, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}
