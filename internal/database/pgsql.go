package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPgSQL(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	return pool, nil
}