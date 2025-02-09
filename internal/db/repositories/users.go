package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"user-service/internal/db/entities"
)

type UserRepo struct {
	conn *sqlx.DB
}

func NewUserRepo(conn *sqlx.DB) UserRepository {
	return &UserRepo{conn: conn}
}

func (ur *UserRepo) GetUserByID(id string) (*entities.UserModel, error) {
	var user entities.UserModel

	sql := "SELECT * FROM users WHERE id=$1"

	err := ur.conn.Get(&user, sql, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) GetUserByUsername(username string) (*entities.UserModel, error) {
	var user entities.UserModel

	sql := "SELECT * FROM users WHERE username=$1"

	err := ur.conn.Get(&user, sql, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) GetUserByAuthCode(authCode string) (*entities.UserModel, error) {
	return nil, nil
}

func (ur *UserRepo) CreateUser(user *entities.UserModel) (*entities.UserModel, error) {
	sql := `INSERT INTO users (username, password_hash, auth_token, email) 
            VALUES (:username, :password_hash, :auth_token, :email)`

	_, err := ur.conn.NamedExec(sql, user)
	if err != nil {
		return nil, err
	}

	sql = "SELECT * FROM users WHERE username=$1"

	err = ur.conn.Get(user, sql, user.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepo) UpdateUser(user *entities.UserModel) (*entities.UserModel, error) {
	sql := `UPDATE users 
	        SET username = :username, password_hash = :password_hash, auth_token = :auth_token, email = :email
	        WHERE id = :id`

	_, err := ur.conn.NamedExec(sql, user)
	if err != nil {
		return nil, err
	}

	sql = "SELECT * FROM users WHERE id=$1"

	err = ur.conn.Get(user, sql, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepo) IssetUser(user entities.UserFilter) (bool, error) {
	var exist bool
	var filterField, sql string

	if user.Username != "" {
		filterField = user.Username
		sql = "SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)"
	} else if user.ID != "" {
		filterField = user.ID
		sql = "SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)"
	} else if user.AuthToken != "" {
		filterField = user.AuthToken
		sql = "SELECT EXISTS (SELECT 1 FROM users WHERE auth_token = $1)"
	} else if user.Email != "" {
		filterField = user.Email
		sql = "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"
	}

	err := ur.conn.Get(&exist, sql, filterField)
	if err != nil {
		return false, err
	}

	return exist, nil
}
