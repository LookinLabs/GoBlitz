package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func NewDBConnection() (*sql.DB, error) {
	conn := fmt.Sprintf(
		// postgres://username:password@host:port/database?sslmode=disable
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("the database refused the connection - %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("the database connection timed out - %v", err)
	}

	log.Println("Connected to the PostgreSQL database!")
	return db, nil
}
