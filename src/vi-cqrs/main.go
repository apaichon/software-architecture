package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"vi-cqrs/api"
	"vi-cqrs/commands"
	"vi-cqrs/queries"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(dbPath string) (*sql.DB, error) {
	// Create database file if it doesn't exist
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Initialize schema
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Initialize database
	db, err := initDB("concert.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize handlers
	commandHandler := commands.NewCommandHandler(db)
	queryHandler := queries.NewQueryHandler(db)
	apiHandler := api.NewHandler(commandHandler, queryHandler)

	// Setup routes
	http.HandleFunc("/api/concerts", apiHandler.GetAvailableConcerts)
	http.HandleFunc("/api/concert", apiHandler.GetConcert)
	http.HandleFunc("/api/purchase", apiHandler.PurchaseTicket)
	http.HandleFunc("/api/create-concert", apiHandler.CreateConcert)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
