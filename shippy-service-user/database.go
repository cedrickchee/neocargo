// Create the database connection

package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewConnection returns a new database connection instance.
func NewConnection() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
