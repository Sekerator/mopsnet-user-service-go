package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user/internal/db/entities"
	"user/internal/db/repositories"
	tokenGenerator "user/pkg/token-generator"
)

type UserServ struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserServices {
	return &UserServ{userRepo}
}

func (us *UserServ) Login(userData entities.UserJson) (*entities.UserModel, error) {
	var user *entities.UserModel

	exist, err := us.userRepo.IssetUser(entities.UserFilter{Username: userData.Username})
	if err != nil {
		return nil, err
	}

	if exist {
		user, err = us.userRepo.GetUserByUsername(userData.Username)
		if err != nil {
			return nil, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userData.Password))
		if err != nil {
			return nil, errors.New("неверный пароль")
		}

		var authToken string
		authToken, err = tokenGenerator.GenerateAuthToken()
		if err != nil {
			return nil, err
		}

		user.AuthToken = authToken
		user, err = us.userRepo.UpdateUser(user)
		if err != nil {
			return nil, err
		}
	} else {
		var hashedPassword []byte
		var authToken string

		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		authToken, err = tokenGenerator.GenerateAuthToken()
		if err != nil {
			return nil, err
		}

		user = &entities.UserModel{
			Username:     userData.Username,
			PasswordHash: string(hashedPassword),
			AuthToken:    authToken,
			Email:        userData.Email,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		user, err = us.userRepo.CreateUser(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
