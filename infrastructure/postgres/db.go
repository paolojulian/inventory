package postgres

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
)

func NewPool() (*pgxpool.Pool, error) {
	config := config.LoadConfig()

	poolConfig, err := pgxpool.ParseConfig(config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database config: %w", err)
	}

	// Force IPv4
	poolConfig.ConnConfig.LookupFunc = func(ctx context.Context, host string) ([]string, error) {
		return []string{host}, nil
	}
	
	// Add connection timeout
	poolConfig.ConnConfig.ConnectTimeout = 10 * time.Second
	
	// Configure pool
	poolConfig.MaxConns = 5
	poolConfig.MinConns = 1
	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}

func MigrateSchema(db *pgxpool.Pool) error {
	// Get the directory of the caller (this file)
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filepath.Dir(filepath.Dir(b))) // go 3 levels up from db.go

	// Now point to the migrations folder
	migrationsPath := filepath.Join(basePath, "infrastructure", "postgres", "migrations")

	m, err := migrate.New(
		"file://"+migrationsPath,
		config.LoadConfig().DatabaseNoSSLURL,
	)
	if err != nil {
		return err
	}

	if config.IsTestEnv() {
		m.Down()
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

