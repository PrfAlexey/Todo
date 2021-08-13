package repository

import (
	"Todo/microservice_auth/user"
	"Todo/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserDatabase struct {
	pool *pgxpool.Pool
}

func NewUserDatabase(conn *pgxpool.Pool) user.Repository {
	return &UserDatabase{pool: conn}
}

func (ud UserDatabase) GetUser(login, password string) (models.User, error) {
	var user models.User
	err := ud.pool.QueryRow(context.Background(), `SELECT * FROM users WHERE username=$1 AND password_hash=$2`, login, password).Scan(&user.Id, &user.Username, &user.Password, &user.Name)

	return user, err
}
