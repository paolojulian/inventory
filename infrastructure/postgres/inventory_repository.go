package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/domain/inventory"
	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/warehouse"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
)

type InventoryRepository struct {
	db *pgxpool.Pool
}

func NewInventoryRepository(db *pgxpool.Pool) *InventoryRepository {
	return &InventoryRepository{db: db}
}

// GetCurrentStock calculates the current stock level for a specific product and warehouse
func (r *InventoryRepository) GetCurrentStock(ctx context.Context, productID, warehouseID string) (*inventory.InventoryItem, error) {
	// First get the product details
	productQuery := `
		SELECT id, sku, name, price_cents, description, is_active
		FROM products 
		WHERE id = $1
	`

	productRow := r.db.QueryRow(ctx, productQuery, productID)
	var p product.Product
	var priceCents int
	err := productRow.Scan(
		&p.ID, &p.SKU, &p.Name, &priceCents, &p.Description, &p.IsActive,
	)
	if err != nil {
		return nil, err
	}
	p.Price = product.Money{Cents: priceCents}

	// Get warehouse details (using users table temporarily)
	warehouseQuery := `
		SELECT id, first_name
		FROM users 
		WHERE id = $1
	`

	warehouseRow := r.db.QueryRow(ctx, warehouseQuery, warehouseID)
	var w warehouse.Warehouse
	err = warehouseRow.Scan(&w.ID, &w.Name)
	if err != nil {
		return nil, err
	}

	// Calculate current stock
	stockQuery := `
		SELECT COALESCE(SUM(quantity_delta), 0) as current_stock, MAX(created_at) as last_updated
		FROM stock_entries 
		WHERE product_id = $1 AND warehouse_id = $2
	`

	stockRow := r.db.QueryRow(ctx, stockQuery, productID, warehouseID)
	var currentStock int
	var lastUpdated *string
	err = stockRow.Scan(&currentStock, &lastUpdated)
	if err != nil {
		return nil, err
	}

	return inventory.NewInventoryItem(&p, &w, currentStock), nil
}

// GetAllCurrentStock calculates current stock levels for all products
func (r *InventoryRepository) GetAllCurrentStock(ctx context.Context, warehouseID string, pager paginationShared.PagerInput) (*inventory.GetAllStockOutput, error) {
	query := `
		WITH stock_data AS (
			SELECT 
				p.id, p.sku, p.name, p.price_cents, p.description, p.is_active,
				COALESCE(SUM(se.quantity_delta), 0) as current_stock,
				w.id as warehouse_id, w.name as warehouse_name
			FROM products p
			LEFT JOIN stock_entries se ON p.id = se.product_id AND se.warehouse_id = $1
			LEFT JOIN warehouses w ON w.id = $1
			GROUP BY p.id, p.sku, p.name, p.price_cents, p.description, p.is_active, w.id, w.name
		)
		SELECT 
			id, sku, name, price_cents, description, is_active,
			current_stock,
			warehouse_id, warehouse_name,
			COUNT(*) OVER() as total_count
		FROM stock_data
		ORDER BY name
		LIMIT $2 OFFSET $3
	`

	offset := (pager.PageNumber - 1) * pager.PageSize
	rows, err := r.db.Query(ctx, query, warehouseID, pager.PageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []*inventory.InventoryItem
	var totalItems int

	for rows.Next() {
		var p product.Product
		var w warehouse.Warehouse
		var priceCents int
		var currentStock int

		err := rows.Scan(
			&p.ID, &p.SKU, &p.Name, &priceCents, &p.Description, &p.IsActive,
			&currentStock,
			&w.ID, &w.Name,
			&totalItems,
		)
		if err != nil {
			return nil, err
		}

		p.Price = product.Money{Cents: priceCents}
		stocks = append(stocks, inventory.NewInventoryItem(&p, &w, currentStock))
	}

	totalPages := (totalItems + pager.PageSize - 1) / pager.PageSize // Ceiling division
	pagerResults := paginationShared.PagerOutput{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pager.PageNumber,
		PageSize:    pager.PageSize,
	}

	return &inventory.GetAllStockOutput{Stocks: stocks, Pager: pagerResults}, nil
}

// GetInventorySummary calculates overall inventory statistics
func (r *InventoryRepository) GetInventorySummary(ctx context.Context, warehouseID string) (*inventory.InventorySummary, error) {
	query := `
		WITH stock_levels AS (
			SELECT 
				p.id,
				p.name,
				p.sku,
				p.price_cents,
				COALESCE(SUM(se.quantity_delta), 0) as current_stock
			FROM products p
			LEFT JOIN stock_entries se ON p.id = se.product_id AND se.warehouse_id = $1
			GROUP BY p.id, p.name, p.sku, p.price_cents
		)
		SELECT 
			COUNT(*) as total_products,
			SUM(current_stock * price_cents) as total_stock_value,
			COUNT(CASE WHEN current_stock > 0 AND current_stock < 10 THEN 1 END) as low_stock_items,
			COUNT(CASE WHEN current_stock <= 0 THEN 1 END) as out_of_stock_items,
			NOW() as last_updated
		FROM stock_levels
	`

	row := r.db.QueryRow(ctx, query, warehouseID)

	var summary inventory.InventorySummary
	err := row.Scan(
		&summary.TotalProducts,
		&summary.TotalStockValue,
		&summary.LowStockItems,
		&summary.OutOfStockItems,
		&summary.LastUpdated,
	)

	if err != nil {
		return nil, err
	}

	return &summary, nil
}

// GetLowStockProducts returns products with low stock levels
func (r *InventoryRepository) GetLowStockProducts(ctx context.Context, warehouseID string, threshold int) ([]*inventory.InventoryItem, error) {
	query := `
		SELECT 
			p.id, p.sku, p.name, p.price_cents, p.description, p.is_active,
			COALESCE(SUM(se.quantity_delta), 0) as current_stock,
			u.id as warehouse_id, u.first_name as warehouse_name
		FROM products p
		LEFT JOIN stock_entries se ON p.id = se.product_id AND se.warehouse_id = $1
		LEFT JOIN users u ON u.id = $1
		GROUP BY p.id, p.sku, p.name, p.price_cents, p.description, p.is_active, u.id, u.first_name
		HAVING COALESCE(SUM(se.quantity_delta), 0) < $2
		ORDER BY current_stock ASC
	`

	rows, err := r.db.Query(ctx, query, warehouseID, threshold)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []*inventory.InventoryItem
	for rows.Next() {
		var p product.Product
		var w warehouse.Warehouse
		var priceCents int
		var currentStock int

		err := rows.Scan(
			&p.ID, &p.SKU, &p.Name, &priceCents, &p.Description, &p.IsActive,
			&currentStock,
			&w.ID, &w.Name,
		)
		if err != nil {
			return nil, err
		}

		p.Price = product.Money{Cents: priceCents}
		stocks = append(stocks, inventory.NewInventoryItem(&p, &w, currentStock))
	}

	return stocks, nil
}

// GetOutOfStockProducts returns products that are out of stock
func (r *InventoryRepository) GetOutOfStockProducts(ctx context.Context, warehouseID string) ([]*inventory.InventoryItem, error) {
	query := `
		SELECT 
			p.id, p.sku, p.name, p.price_cents, p.description, p.is_active,
			COALESCE(SUM(se.quantity_delta), 0) as current_stock,
			u.id as warehouse_id, u.first_name as warehouse_name
		FROM products p
		LEFT JOIN stock_entries se ON p.id = se.product_id AND se.warehouse_id = $1
		LEFT JOIN users u ON u.id = $1
		GROUP BY p.id, p.sku, p.name, p.price_cents, p.description, p.is_active, u.id, u.first_name
		HAVING COALESCE(SUM(se.quantity_delta), 0) <= 0
		ORDER BY p.name
	`

	rows, err := r.db.Query(ctx, query, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []*inventory.InventoryItem
	for rows.Next() {
		var p product.Product
		var w warehouse.Warehouse
		var priceCents int
		var currentStock int

		err := rows.Scan(
			&p.ID, &p.SKU, &p.Name, &priceCents, &p.Description, &p.IsActive,
			&currentStock,
			&w.ID, &w.Name,
		)
		if err != nil {
			return nil, err
		}

		p.Price = product.Money{Cents: priceCents}
		stocks = append(stocks, inventory.NewInventoryItem(&p, &w, currentStock))
	}

	return stocks, nil
}
