package postgres

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/domain/warehouse"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
)

type StockRepository struct {
	db *pgxpool.Pool
}

func NewStockRepository(db *pgxpool.Pool) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) CreateStockEntry(ctx context.Context, stockEntry *stock.StockEntry) (*stock.StockEntry, error) {
	row := r.db.QueryRow(ctx, `
		INSERT INTO stock_entries (
			id,
			product_id,
			warehouse_id,
			user_id,
			quantity_delta,
			reason,
			supplier_price_cents,
			store_price_cents,
			expiry_date,
			reorder_date,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, product_id, warehouse_id, user_id, quantity_delta, reason, supplier_price_cents, store_price_cents, expiry_date, reorder_date, created_at
	`,
		stockEntry.ID,
		stockEntry.ProductID,
		stockEntry.WarehouseID,
		stockEntry.UserID,
		stockEntry.QuantityDelta,
		stockEntry.Reason,
		stockEntry.SupplierPriceCents,
		stockEntry.StorePriceCents,
		stockEntry.ExpiryDate,
		stockEntry.ReorderDate,
		stockEntry.CreatedAt,
	)

	var created stock.StockEntry
	if err := row.Scan(
		&created.ID,
		&created.ProductID,
		&created.WarehouseID,
		&created.UserID,
		&created.QuantityDelta,
		&created.Reason,
		&created.SupplierPriceCents,
		&created.StorePriceCents,
		&created.ExpiryDate,
		&created.ReorderDate,
		&created.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *StockRepository) GetByID(ctx context.Context, stockEntryID string) (*stock.StockEntry, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, product_id, warehouse_id, user_id, quantity_delta, reason, supplier_price_cents, store_price_cents, expiry_date, reorder_date, created_at
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
		&found.SupplierPriceCents,
		&found.StorePriceCents,
		&found.ExpiryDate,
		&found.ReorderDate,
		&found.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &found, nil
}

func (r *StockRepository) GetList(ctx context.Context, pager *paginationShared.PagerInput) ([]*stock.StockEntryWithRelations, int, error) {
	// First get the total count
	var totalCount int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM stock_entries`).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Then get the entries
	offset := (pager.PageNumber - 1) * pager.PageSize
	rows, err := r.db.Query(ctx, `
		WITH stock_data AS (
			SELECT
				se.id, se.product_id, se.warehouse_id, se.user_id, se.quantity_delta, se.reason, se.supplier_price_cents, se.store_price_cents, se.expiry_date, se.reorder_date, se.created_at,
				p.id as p_id, p.sku, p.name as product_name, p.price_cents, p.description, p.is_active,
				w.id as w_id, w.name as warehouse_name,
				u.id as u_id, u.first_name as user_first_name, u.last_name as user_last_name
			FROM stock_entries se
			LEFT JOIN products p ON se.product_id = p.id
			LEFT JOIN warehouses w ON se.warehouse_id = w.id
			LEFT JOIN users u ON se.user_id = u.id
		)
		SELECT
			id, product_id, warehouse_id, user_id, quantity_delta, reason, supplier_price_cents, store_price_cents, expiry_date, reorder_date, created_at,
			p_id, sku, product_name, price_cents, description, is_active,
			w_id, warehouse_name,
			u_id, user_first_name, user_last_name
		FROM stock_data
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, pager.PageSize, offset)

	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var stockEntries []*stock.StockEntryWithRelations

	for rows.Next() {
		var entry stock.StockEntry
		var p product.Product
		var w warehouse.Warehouse
		var u user.User
		var priceCents int

		if err := rows.Scan(
			&entry.ID, &entry.ProductID, &entry.WarehouseID, &entry.UserID, &entry.QuantityDelta, &entry.Reason, &entry.SupplierPriceCents, &entry.StorePriceCents, &entry.ExpiryDate, &entry.ReorderDate, &entry.CreatedAt,
			&p.ID, &p.SKU, &p.Name, &priceCents, &p.Description, &p.IsActive,
			&w.ID, &w.Name,
			&u.ID, &u.FirstName, &u.LastName,
		); err != nil {
			return nil, 0, err
		}

		p.Price = product.Money{Cents: priceCents}

		stockEntries = append(stockEntries, &stock.StockEntryWithRelations{
			StockEntry: entry,
			Product:    &p,
			Warehouse:  &w,
			User:       &u,
		})
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
		RETURNING id, product_id, warehouse_id, user_id, quantity_delta, reason, supplier_price_cents, store_price_cents, expiry_date, reorder_date, created_at
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
