package user

// UserRepository defines the interface for user persistence operations.
type UserRepository interface {
	FindByID(id int) (*User, error)
	// Other methods like Create, Update, Delete, etc.
}
