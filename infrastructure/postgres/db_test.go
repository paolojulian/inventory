package postgres_test

import (
	"context"
	"log"
	"testing"

	"paolojulian.dev/inventory/infrastructure/postgres"
)

func TestInitDB_ConnectedSuccessfully(t *testing.T) {
	pool, err := postgres.NewPool()
	if err != nil {
		log.Fatalf("postgress failed to connect with error: %v", err)
	}

	// Make sure to close pool after this func
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("ping failed. err: %v", err)
	}
}