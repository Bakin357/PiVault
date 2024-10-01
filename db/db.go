package db

import (
	"database/sql" // Go's standard package for working with SQL databases
	"log"          // Standard package for logging errors and other information

	"github.com/go-sql-driver/mysql" // MySQL driver for Go, provides connection to MySQL databases
)

// NewMySQLStorage creates and returns a new MySQL database connection
// cfg is the configuration struct that holds the MySQL connection details
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	// sql.Open initializes a new database connection using the MySQL driver and the provided Data Source Name (DSN)
	// FormatDSN formats the configuration into a DSN string
	db, err := sql.Open("mysql", cfg.FormatDSN())

	// If there is an error while opening the database connection, log the error and terminate the program
	if err != nil {
		log.Fatal(err)
	}

	// Return the database connection object (db) and no error (nil)
	// In case of successful connection, err will be nil
	return db, nil
}
