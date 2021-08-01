package repository

import "Todo/models"

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}
type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *PostgresDB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
