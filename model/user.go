package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size64, index"`
	Username string `gorm:"size64, index"`
	Password string `gorm:"size255"`
}
