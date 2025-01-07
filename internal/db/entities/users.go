package entities

import "time"

type UserModel struct {
	ID           string    `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	AuthToken    string    `json:"auth_token" db:"auth_token"`
	Email        string    `json:"email" db:"email"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserFilter struct {
	ID        string `db:"id"`
	Username  string `db:"username"`
	AuthToken string `db:"auth_token"`
	Email     string `db:"email"`
}

type UserReturnJson struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	AuthToken string    `json:"auth_token"`
	CreatedAt time.Time `json:"created_at"`
}
