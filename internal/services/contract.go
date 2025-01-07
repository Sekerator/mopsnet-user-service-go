package services

import (
	"user/internal/db/entities"
)

type UserServices interface {
	Login(userData entities.UserJson) (*entities.UserModel, error)
}
