package service

import (
	"github.com/gabrielmartins0312/crud-go-with-api-and-db/model"
	"github.com/gabrielmartins0312/crud-go-with-api-and-db/repository"
)

func CreateUser(user model.User) error {
	return repository.CreateUser(user)
}

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}

func UpdateUser(user model.User) error {
	return repository.UpdateUser(user)
}
