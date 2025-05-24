package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/domain/product"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(ctx context.Context, p *product.Product) (*product.Product, error) {
	row := r.db.QueryRow(ctx, `
		INSERT INTO products (id, sku, name, description, price_cents, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, sku, name, description, price_cents, is_active
	`, p.ID, p.SKU, p.Name, p.Description, p.Price.Cents, p.IsActive)

	var created product.Product
	var priceCents int

	if err := row.Scan(
		&created.ID,
		&created.SKU,
		&created.Name,
		&created.Description,
		&priceCents,
		&created.IsActive,
	); err != nil {
		return nil, err
	}

	created.Price = product.Money{Cents: priceCents}

	return &created, nil
}

func (r *ProductRepository) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM products WHERE sku = $1)
	`, sku).Scan(&exists)

	return exists, err
}
