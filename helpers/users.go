package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encryptedPassword), nil
}

func CheckPasswordMatch(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckUserExists(checkUserExists func(string) (bool, error), value string) (bool, error) {
	userExists, err := checkUserExists(value)
	if err != nil {
		return false, err
	}

	return userExists, nil
}
