// internal/user/user.go

package user

// User defines the structure for the user model
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	// Add other fields as necessary
}
