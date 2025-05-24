package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
)

func NewPool() (*pgxpool.Pool, error) {
	config := config.LoadConfig()

	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}
