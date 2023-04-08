package repository

import (
	entity "kayn.ooo/api/src/Entity"
)

func GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUsers() ([]entity.User, error) {
	var users []entity.User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
