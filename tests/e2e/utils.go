package e2e

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
)

func cleanupTables(ctx context.Context, db *pgxpool.Pool) error {
	if !config.IsTestEnv() {
		return errors.New("You are not using a test env, cleanup failed.")
	}

	// Get all table names first
	rows, err := db.Query(ctx, "SELECT tablename FROM pg_tables WHERE schemaname = 'public'")
	if err != nil {
		return err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			return err
		}
		tables = append(tables, tablename)
	}

	// Truncate each table
	for _, table := range tables {
		_, err := db.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			return fmt.Errorf("failed to truncate table %s: %w", table, err)
		}
	}

	return nil
}
