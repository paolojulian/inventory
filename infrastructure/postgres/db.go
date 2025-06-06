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
	"paolojulian.dev/inventory/domain/user"
)

func NewPool() (*pgxpool.Pool, error) {
	config := config.LoadConfig()

	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
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
