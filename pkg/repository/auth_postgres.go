package repository

import (
	"Todo/models"
	"context"
	"fmt"
)

type AuthPostgres struct {
	db *PostgresDB
}

func NewAuthPostgres(db *PostgresDB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.pool.QueryRow(context.Background(), query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	err := r.db.pool.QueryRow(context.Background(), `SELECT * FROM users WHERE username=$1 AND password_hash=$2`, username, password).Scan(&user.Id, &user.Username, &user.Password, &user.Name)

	return user, err
}
