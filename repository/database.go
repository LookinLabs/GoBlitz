package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	model "web/model"

	// Import pq to register the Postgres driver.
	_ "github.com/lib/pq"
)

func NewDBConnection() (*model.Database, error) {
	if os.Getenv("PSQL_ENABLED") != "true" && os.Getenv("APP_ENV") == "development" {
		log.Println("Warning: PSQL is not enabled. Database queries will fail.")
		return nil, nil
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
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &model.Database{DB: db}, nil
}
