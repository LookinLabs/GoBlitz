package repository

import (
	"database/sql"
	"fmt"
	"log"
	"web/model"

	_ "github.com/lib/pq"
)

func NewDBConnection(env model.PostgresConfig) (*sql.DB, error) {
    connStr := "postgres://" + env.DBUser + ":" + env.DBPassword + "@" + env.DBHost + ":" + env.DBPort + "/" + env.DBDatabase + "?sslmode=disable"

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