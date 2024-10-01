package main

import (
	"database/sql"
	"log" // Standard library package for logging errors and messages

	// Importing your custom packages
	config "github.com/Bakin357/PiVault/configs" // Config package for environment variables
	"github.com/Bakin357/PiVault/db"             // DB package for database interactions
	"github.com/Bakin357/PiVault/src/api"        // API package for handling server logic
	"github.com/go-sql-driver/mysql"             // MySQL driver package for database connection
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	InitStorage(db)

	defer db.Close()

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}
