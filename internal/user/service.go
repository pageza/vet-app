// internal/user/service.go

package user

// UserService handles business logic for users.
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUserByID gets a user by their ID using the UserRepository.
func (s *UserService) GetUserByID(id int) (*User, error) {
	return s.repo.FindByID(id)
}
