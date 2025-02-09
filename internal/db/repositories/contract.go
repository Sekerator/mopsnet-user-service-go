package repositories

import (
	"user-service/internal/db/entities"
)

type UserRepository interface {
	GetUserByID(id string) (*entities.UserModel, error)
	GetUserByUsername(username string) (*entities.UserModel, error)
	GetUserByAuthCode(authCode string) (*entities.UserModel, error)
	CreateUser(user *entities.UserModel) (*entities.UserModel, error)
	UpdateUser(user *entities.UserModel) (*entities.UserModel, error)
	IssetUser(user entities.UserFilter) (bool, error)
}
