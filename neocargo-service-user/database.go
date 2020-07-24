// Create the database connection

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewConnection returns a new database connection instance.
func NewConnection() (*sqlx.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DBNAME")
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	log.Printf("conn: %v\n", conn)

	db, err := sqlx.Connect("postgres", conn)
	log.Println("sqlx connect done")

	if err != nil {
		log.Printf("sqlx connect err: %v\n", err)
		return nil, err
	}

	log.Printf("DB connection OK: %v\n", db)
	return db, nil
}
