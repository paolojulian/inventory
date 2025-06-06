package e2e

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
)

func cleanupTables(ctx context.Context, db *pgxpool.Pool) error {
	if !config.IsTestEnv() {
		return errors.New("You are not using a test env, cleanup failed.")
	}

	_, err := db.Exec(ctx, `
		DO
		$$
		DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
				EXECUTE 'TRUNCATE TABLE public.' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE';
			END LOOP;
		END
		$$;
	`)
	return err
}
