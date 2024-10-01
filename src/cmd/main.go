package main

import (
	"log" // Standard library package for logging errors and messages

	// Importing your custom packages
	config "github.com/Bakin357/PiVault/configs" // Config package for environment variables
	"github.com/Bakin357/PiVault/db"             // DB package for database interactions
	"github.com/Bakin357/PiVault/src/api"        // API package for handling server logic
	"github.com/go-sql-driver/mysql"             // MySQL driver package for database connection
)

func main() {
	// Create the MySQL connection using configuration values from the config package
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,     // Database user from environment variables
		Passwd:               config.Envs.DBPassword, // Database password from environment variables
		Addr:                 config.Envs.DBAddress,  // Database address (host:port) from environment variables
		DBName:               config.Envs.DBName,     // Database name from environment variables
		Net:                  "tcp",                  // Network type (TCP) to connect to MySQL
		AllowNativePasswords: true,                   // Enables native password authentication in MySQL
		ParseTime:            true,                   // Parses SQL DATETIME fields into Go's time.Time format
	})

	// If there is an error in the database connection, log it and terminate the program
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ensure the database connection is closed after the main function finishes execution
	defer db.Close()

	// Initialize the API server, passing the database connection and setting it to listen on port 8080
	server := api.NewAPIServer(":8080", db)

	// Start the API server and log an error if it fails to start
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
