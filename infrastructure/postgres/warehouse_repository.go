package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	warehouseDomain "paolojulian.dev/inventory/domain/warehouse"
)

type WarehouseRepository struct {
	db *pgxpool.Pool
}

func NewWarehouseRepository(db *pgxpool.Pool) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

func (r *WarehouseRepository) GetSummary(ctx context.Context, warehouseID string) (*warehouseDomain.WarehouseSummary, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, name
		FROM warehouses
		WHERE id = $1
	`, warehouseID)

	var summary warehouseDomain.WarehouseSummary
	if err := row.Scan(
		&summary.ID,
		&summary.Name,
	); err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *WarehouseRepository) GetDefaultWarehouse(ctx context.Context) (*warehouseDomain.WarehouseSummary, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, name
		FROM warehouses
		WHERE name = 'Default'
		LIMIT 1
	`)

	var summary warehouseDomain.WarehouseSummary
	if err := row.Scan(
		&summary.ID,
		&summary.Name,
	); err != nil {
		return nil, err
	}

	return &summary, nil
}

