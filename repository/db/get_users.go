package sql

import (
	"web/model"
)

func GetUsers() ([]model.User, error) {
	var users []model.User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
