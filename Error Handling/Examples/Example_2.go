// Handling Database Errors

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// _ "github.com/lib/pq" // Import the PostgreSQL driver
)

func getUser(db *sql.DB, id int) (string, error) {
	var name string
	err := db.QueryRow("SELECT name FROM users WHERE id = $1", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("getUser: user with id %d not found", id)
		}
		return "", fmt.Errorf("getUser: failed to query user: %w", err)
	}
	return name, nil
}

func main() {
	// Connection string (replace with your actual credentials)
	connStr := "user=postgres password=password dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	user, err := getUser(db, 123)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("User:", user)
}
