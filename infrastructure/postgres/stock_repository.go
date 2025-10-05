package postgres

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/domain/stock"
)

type StockRepository struct {
	db *pgxpool.Pool
}

func NewStockRepository(db *pgxpool.Pool) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) CreateStockEntry(ctx context.Context, stockEntry *stock.StockEntry) (*stock.StockEntry, error) {
	row := r.db.QueryRow(ctx, `
		INSERT INTO stock_entries (id, product_id, warehouse_id, user_id, quantity_delta, reason, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, product_id, warehouse_id, user_id, quantity_delta, reason, created_at
	`, stockEntry.ID, stockEntry.ProductID, stockEntry.WarehouseID, stockEntry.UserID, stockEntry.QuantityDelta, stockEntry.Reason, stockEntry.CreatedAt)

	var created stock.StockEntry
	if err := row.Scan(
		&created.ID,
		&created.ProductID,
		&created.WarehouseID,
		&created.UserID,
		&created.QuantityDelta,
		&created.Reason,
		&created.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *StockRepository) GetByID(ctx context.Context, stockEntryID string) (*stock.StockEntry, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, product_id, warehouse_id, user_id, quantity_delta, reason, created_at
		FROM stock_entries
		WHERE id = $1
	`, stockEntryID)

	var found stock.StockEntry
	if err := row.Scan(
		&found.ID,
		&found.ProductID,
		&found.WarehouseID,
		&found.UserID,
		&found.QuantityDelta,
		&found.Reason,
		&found.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &found, nil
}

func (r *StockRepository) GetList(ctx context.Context, limit int) ([]*stock.StockEntry, int, error) {
	// First get the total count
	var totalCount int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM stock_entries`).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Then get the entries
	rows, err := r.db.Query(ctx, `
		SELECT id, product_id, warehouse_id, user_id, quantity_delta, reason, created_at
		FROM stock_entries
		ORDER BY created_at DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	stockEntries := []*stock.StockEntry{}
	for rows.Next() {
		var entry stock.StockEntry
		if err := rows.Scan(
			&entry.ID,
			&entry.ProductID,
			&entry.WarehouseID,
			&entry.UserID,
			&entry.QuantityDelta,
			&entry.Reason,
			&entry.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		stockEntries = append(stockEntries, &entry)
	}

	return stockEntries, totalCount, nil
}

func (r *StockRepository) UpdateByID(ctx context.Context, stockEntryID string, patch *stock.StockEntryPatch) (*stock.StockEntry, error) {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if patch.QuantityDelta != nil {
		setClauses = append(setClauses, "quantity_delta = $"+strconv.Itoa(argPos))
		args = append(args, *patch.QuantityDelta)
		argPos++
	}
	if patch.Reason != nil {
		setClauses = append(setClauses, "reason = $"+strconv.Itoa(argPos))
		args = append(args, *patch.Reason)
		argPos++
	}
	if patch.ProductID != nil {
		setClauses = append(setClauses, "product_id = $"+strconv.Itoa(argPos))
		args = append(args, *patch.ProductID)
		argPos++
	}
	if patch.WarehouseID != nil {
		setClauses = append(setClauses, "warehouse_id = $"+strconv.Itoa(argPos))
		args = append(args, *patch.WarehouseID)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil, errors.New("no fields to update")
	}

	args = append(args, stockEntryID)
	query := `
		UPDATE stock_entries
		SET ` + strings.Join(setClauses, ", ") + `
		WHERE id = $` + strconv.Itoa(argPos) + `
		RETURNING id, product_id, warehouse_id, user_id, quantity_delta, reason, created_at
	`

	row := r.db.QueryRow(ctx, query, args...)

	var updated stock.StockEntry
	if err := row.Scan(
		&updated.ID,
		&updated.ProductID,
		&updated.WarehouseID,
		&updated.UserID,
		&updated.QuantityDelta,
		&updated.Reason,
		&updated.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *StockRepository) Delete(ctx context.Context, stockEntryID string) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM stock_entries WHERE id = $1
	`, stockEntryID)

	return err
}
