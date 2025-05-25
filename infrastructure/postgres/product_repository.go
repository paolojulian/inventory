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

func (r *ProductRepository) Delete(ctx context.Context, productId string) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM products WHERE id = $1
	`, productId)

	return err // will only be non-nil if query fails
}

func (r *ProductRepository) GetByID(ctx context.Context, productID string) (*product.Product, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, sku, name, description, price_cents, is_active
		FROM products
		WHERE id = $1
	`, productID)

	var found product.Product
	var priceCents int

	if err := row.Scan(
		&found.ID,
		&found.SKU,
		&found.Name,
		&found.Description,
		&priceCents,
		&found.IsActive,
	); err != nil {
		return nil, err
	}

	found.Price = product.Money{Cents: priceCents}

	return &found, nil
}

func (r *ProductRepository) UpdateByID(ctx context.Context, productID string, p *product.Product) (*product.Product, error) {
	row := r.db.QueryRow(ctx, `
		UPDATE products
		SET sku = $1, name = $2, description = $3, price_cents = $4
		WHERE id = $5
		RETURNING id, sku, name, description, price_cents, is_active
	`, p.SKU, p.Name, p.Description, p.Price.Cents, productID)

	var updated product.Product
	var priceCents int

	if err := row.Scan(
		&updated.ID,
		&updated.SKU,
		&updated.Name,
		&updated.Description,
		&priceCents,
		&updated.IsActive,
	); err != nil {
		return nil, err
	}

	updated.Price = product.Money{Cents: priceCents}

	return &updated, nil
}

func (r *ProductRepository) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM products WHERE sku = $1)
	`, sku).Scan(&exists)

	return exists, err
}
