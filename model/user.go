package model

import (
	"errors"
	"log"

	helper "web/helpers"

	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Email    string `gorm:"size64, index"`
	Username string `gorm:"size64, index"`
	Password string `gorm:"size255"`
}

func CheckUserExistance(email string) bool {
	var user User

	if err := DB.Where("email = ?", email).First(&user).Error; err == nil {
		return true
	}

	if err := DB.Where("username = ?", email).First(&user).Error; err == nil {
		return true
	}

	return false
}

func CreateUser(email, username, password string) *User {
	encryptedPassword, err := helper.EncryptPassword(password)
	if err != nil {
		log.Fatalf("Failed to encrypt password: %v", err)
	}

	user := User{
		Email:    email,
		Username: username,
		Password: encryptedPassword,
	}
	DB.Create(&user)
	return &user
}

func CheckPasswordMatch(username, password string) (*User, error) {
	var user User
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, errors.New("invalid username")
	}

	err := helper.CheckPasswordMatch(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

func GetUserByID(userID uint) (User, error) {
	var user User
	if err := DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}
	return user, nil
}
