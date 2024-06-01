package sql

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDBConnection() error {
	dbConnAttrs := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dbConnAttrs), &gorm.Config{})
	if err != nil {
		return err
	}

	// assign the *gorm.DB instance to DB
	DB = db
	return nil
}
