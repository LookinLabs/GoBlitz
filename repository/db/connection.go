package sql

import (
	"fmt"
	"os"
	"web/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	dbConnAttrs := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dbConnAttrs), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if db.AutoMigrate(&model.User{}) != nil {
		return nil, err
	}

	return db, nil
}
