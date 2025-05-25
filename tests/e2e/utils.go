package e2e

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func cleanupTables(ctx context.Context, db *pgxpool.Pool) error {
	_, err := db.Exec(ctx, `
		DO
		$$
		DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
				EXECUTE 'DROP TABLE IF EXISTS public.' || quote_ident(r.tablename) || ' CASCADE';
			END LOOP;
		END
		$$;
	`)
	return err
}
