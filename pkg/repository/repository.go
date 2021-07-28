package repository

import "ads/models"

type Authorization interface {
	CreateUser(user models.User) (int, error)
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
