package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/user"
)

type PgxQuerier interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewPool() (*pgxpool.Pool, error) {
	config := config.LoadConfig()

	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}

func MigrateSchema(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS products (
			id UUID PRIMARY KEY,
			sku TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			description TEXT,
			price_cents INTEGER NOT NULL,
			is_active BOOLEAN NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);

		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL,
			is_active BOOLEAN NOT NULL,
			email TEXT,
			first_name TEXT,
			last_name TEXT,
			mobile TEXT,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);

		CREATE OR REPLACE FUNCTION trigger_set_updated_at()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.updated_at = NOW();
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_trigger WHERE tgname = 'set_updated_at_products'
			) THEN
				CREATE TRIGGER set_updated_at_products
				BEFORE UPDATE ON products
				FOR EACH ROW
				EXECUTE FUNCTION trigger_set_updated_at();
			END IF;

			IF NOT EXISTS (
				SELECT 1 FROM pg_trigger WHERE tgname = 'set_updated_at_users'
			) THEN
				CREATE TRIGGER set_updated_at_users
				BEFORE UPDATE ON users
				FOR EACH ROW
				EXECUTE FUNCTION trigger_set_updated_at();
			END IF;
		END;
		$$;
	`)
	return err
}

func PopulateInitialData(db *pgxpool.Pool) error {
	defaultAdmin, err := user.NewUser(
		"theman",
		"qwe123!",
		user.AdminRole,
		true,
		config.StringPointer("johndoe@email.com"),
		config.StringPointer("Serrah"),
		config.StringPointer("Kerrigan"),
		config.StringPointer("09279488654"),
	)
	if err != nil {
		return err
	}

	// Set a 15-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// GET, if not existing, add
	repo := NewUserRepository(db)
	existingSuperAdmin, err := repo.FindByUsername(ctx, "theman")
	if err != nil {
		return err
	}

	if existingSuperAdmin != nil {
		// Super admin already exists
		return nil
	}

	if _, err := repo.Save(ctx, defaultAdmin); err != nil {
		return err
	}

	return nil
}
