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
	log.Printf("DB_HOST: %v\n", os.Getenv("DB_HOST"))

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
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
