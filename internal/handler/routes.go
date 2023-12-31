package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pageza/vet-app/internal/user"
)

// SetupRoutes initializes and returns a new router with all the routes.
func SetupRoutes(authHandler *AuthHandler, userService *user.UserService) *mux.Router {
	r := mux.NewRouter()

	// Define Routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/profile", authHandler.Profile).Methods("GET")

	r.HandleFunc("/profile/update", authHandler.UpdateProfile).Methods("PUT")
	r.HandleFunc("/profile/change-password", authHandler.ChangePassword).Methods("POST")
	r.HandleFunc("/user/delete", authHandler.DeleteAccount).Methods("DELETE")
	r.HandleFunc("/users", authHandler.UserList).Methods("GET")
	r.HandleFunc("/user/{id}/role", authHandler.ManageUserRole).Methods("PUT")
	r.HandleFunc("/password-reset-request", authHandler.PasswordResetRequest).Methods("POST")
	r.HandleFunc("/password-reset", authHandler.PasswordReset).Methods("POST")
	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Route logic here
	})

	return r
}
