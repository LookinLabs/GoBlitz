package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func NewDBConnection() (*sql.DB, error) {
	psqlHost := os.Getenv("POSTGRES_HOST")
	psqlPort := os.Getenv("POSTGRES_PORT")
	psqlUser := os.Getenv("POSTGRES_USER")
	psqlPassword := os.Getenv("POSTGRES_PASSWORD")
	psqlDB := os.Getenv("POSTGRES_DB")

	connStr := "postgres://" + psqlUser + ":" + psqlPassword + "@" + psqlHost + ":" + psqlPort + "/" + psqlDB + "?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("Connected to the PostgreSQL database!")
	return db, nil
}
