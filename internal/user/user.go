// internal/user/user.go

package user

import "time"

// User defines the structure for the user model
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`

	Role        string    `json:"role"`        // User role (e.g., admin, user, moderator)
	FirstName   string    `json:"firstName"`   // User's first name
	LastName    string    `json:"lastName"`    // User's last name
	Bio         string    `json:"bio"`         // Short biography or description
	ResetToken  string    `json:"resetToken"`  // Token for password reset
	TokenExpiry time.Time `json:"tokenExpiry"` // Expiry time for the reset token
	Status      string    `json:"status"`      // Account status (active, suspended, etc.)
	CreatedAt   time.Time `json:"createdAt"`   // Timestamp for account creation
	UpdatedAt   time.Time `json:"updatedAt"`   // Timestamp for last update
	LastLogin   time.Time `json:"lastLogin"`   // Timestamp for last login
	// Add other fields as necessary
}
