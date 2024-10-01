package user

import (
	"net/http" // Standard library for handling HTTP requests and responses

	"github.com/gorilla/mux" // Gorilla Mux is a powerful HTTP router for Go
)

// Handler struct is responsible for handling user-related HTTP requests (login, registration, etc.)
type Handler struct {
}

// NewHandler creates and returns a new instance of the Handler struct
// This will be used to register and handle user-related routes
func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRoutes adds user-related routes to the provided Gorilla Mux router
// It registers the login and register routes as POST requests
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Register the login route
	// When a POST request is made to "/login", the handleLogin method will be invoked
	router.HandleFunc("/login", h.handleLogin).Methods("POST")

	// Register the register route
	// When a POST request is made to "/register", the handleRegister method will be invoked
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handleLogin handles the login logic when a POST request is made to "/login"
// Currently, it does not perform any actions (empty function body)
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// This is where you would add your login logic
	// For example, checking credentials, generating a session or JWT token, etc.
}

// handleRegister handles the registration logic when a POST request is made to "/register"
// Currently, it does not perform any actions (empty function body)
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// This is where you would add your registration logic
	// For example, creating a new user, hashing a password, storing user info in the database, etc.
}
