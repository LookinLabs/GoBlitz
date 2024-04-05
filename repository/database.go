package repository

import (
	"database/sql"
	"fmt"
	"log"
	envModel "web/model/config"
)

func NewDBConnection(env envModel.PostgresEnv) (*sql.DB, error) {
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
