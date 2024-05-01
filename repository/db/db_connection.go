package sql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Import pq to register the Postgres driver.
	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDBConnection() error {
	if os.Getenv("PSQL_ENABLED") != "true" && os.Getenv("APP_ENV") == "development" {
		log.Println("Warning: PSQL is not enabled. Database queries will fail.")
		return nil
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	// assign the *sql.DB instance to DB
	DB = db
	return nil
}