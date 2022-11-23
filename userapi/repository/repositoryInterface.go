package repository

import (
	"refactoring/api/request"
	"refactoring/models"
)

type Repository interface {
	CreateUser(request request.CreateUserRequest) string
	UpdateUser(request request.UpdateUserRequest) error
	DeleteUser(id string) error
	GetAllUsers() map[string]models.User
}
