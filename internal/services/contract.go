package services

import (
	"user-service/internal/db/entities"
)

type UserServices interface {
	Login(userData entities.UserJson) (*entities.UserModel, error)
}
