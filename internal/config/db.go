package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDb() (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s",
		"admin",
		"secret",
		"notes_management_db")

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}
