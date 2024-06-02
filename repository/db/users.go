package sql

import (
	"errors"
	"log"
	helper "web/helpers"
	"web/model"

	"gorm.io/gorm"
)

var DB *gorm.DB

func GetUser() (model.User, error) {
	var user model.User

	transaction := DB.First(&user)
	if transaction.Error != nil {
		if errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}

		return model.User{}, transaction.Error
	}

	return user, nil
}

func GetUserByID(userID uint) (model.User, error) {
	var user model.User

	transaction := DB.First(&user, userID)
	if transaction.Error != nil {
		if errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}

		return model.User{}, transaction.Error
	}

	return user, nil
}

func GetUserIDByUsername(username string) (uint, error) {
	var user model.User

	transaction := DB.Where("username = ?", username).First(&user)
	if transaction.Error != nil {
		if errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("user not found")
		}

		return 0, transaction.Error
	}

	return user.ID, nil
}

func CheckUserExistenceByEmail(email string) (bool, error) {
	var user model.User

	transaction := DB.Where("email = ?", email).First(&user)
	if transaction.Error != nil {
		if errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, transaction.Error
	}

	return true, nil
}

func CheckUserExistenceByUsername(username string) (bool, error) {
	var user model.User

	transaction := DB.Where("username = ?", username).First(&user)
	if transaction.Error != nil {
		if errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, transaction.Error
	}

	return true, nil
}

func CheckPasswordMatch(username, password string) (*model.User, error) {
	var user model.User

	transaction := DB.Where("username = ?", username).First(&user)
	if transaction.Error != nil {
		if errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username")
		}

		return nil, transaction.Error
	}

	if err := helper.CheckPasswordMatch(user.Password, password); err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func CreateUser(email, username, password string) *model.User {
	encryptedPassword, err := helper.EncryptPassword(password)
	if err != nil {
		log.Fatalf("Failed to encrypt password: %v", err)
	}

	user := model.User{
		Email:    email,
		Username: username,
		Password: encryptedPassword,
	}

	if transaction := DB.Create(&user); transaction.Error != nil {
		log.Printf("Failed to create user: %v", transaction.Error)
	}

	log.Printf("User created: %v", user)

	return &user
}
