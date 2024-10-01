package api

import (
	"database/sql" // Go's standard package for working with SQL databases
	"log"          // Standard package for logging messages
	"net/http"     // Package for handling HTTP requests

	"github.com/Bakin357/PiVault/src/user" // Importing user package to handle user-related routes
	"github.com/gorilla/mux"               // Gorilla Mux is a powerful HTTP router for Go
)

// APIServer represents the API server with an address and a database connection
type APIServer struct {
	addr string  // The address (host:port) the API server listens on
	db   *sql.DB // A pointer to the SQL database connection
}

// NewAPIServer creates and returns a new instance of APIServer
// Takes an address string and a database connection as parameters
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	// Initializes a new APIServer struct with the provided address and database connection
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server and listens on the specified address
func (s *APIServer) Run() error {
	// Initialize a new Gorilla Mux router
	router := mux.NewRouter()

	// Create a subrouter with a path prefix of "/api/v1"
	// All routes starting with "/api/v1" will be handled by this subrouter
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new user handler from the user package and register the user routes
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter) // Register user-related routes with the subrouter

	// Log the address the server is listening on
	log.Println("Listening on", s.addr)

	// Start the HTTP server and listen on the specified address (s.addr)
	// Pass the router as the handler for HTTP requests
	return http.ListenAndServe(s.addr, router)
}
