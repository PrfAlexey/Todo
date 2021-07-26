package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DBConnect = "user=postgres dbname=TodoDB password=1234 host=localhost port=5432 sslmode=disable"
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
