package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DBConnect = "user=postgres dbname=postgres password=4444 host=localhost port=5432 sslmode=disable pool_max_conns=50"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_lists"
	userListsTable  = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type PostgresDB struct {
	pool *pgxpool.Pool
}

func NewPostgresDB() (*PostgresDB, error) {
	pool, err := pgxpool.Connect(context.Background(), DBConnect)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &PostgresDB{pool: pool}, nil
}
