package postgres

import (
	"context"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/domain/product"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
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

func (r *ProductRepository) ActivateProductByID(ctx context.Context, productID string) (*product.Product, error) {
	row := r.db.QueryRow(ctx, `
		UPDATE products
		SET is_active = $1
		WHERE id = $2
		RETURNING id, sku, name, description, price_cents, is_active
	`, true, productID)

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

func (r *ProductRepository) DeactivateProductByID(ctx context.Context, productID string) (*product.Product, error) {
	row := r.db.QueryRow(ctx, `
		UPDATE products
		SET is_active = $1
		WHERE id = $2
		RETURNING id, sku, name, description, price_cents, is_active
	`, false, productID)

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

func (r *ProductRepository) UpdateByID(ctx context.Context, productID string, p *product.ProductPatch) (*product.Product, error) {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if p.SKU != nil {
		setClauses = append(setClauses, "sku = $"+strconv.Itoa(argPos))
		args = append(args, *p.SKU)
		argPos++
	}
	if p.Name != nil {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, *p.Name)
		argPos++
	}
	if p.Description != nil {
		setClauses = append(setClauses, "description = $"+strconv.Itoa(argPos))
		args = append(args, *p.Description)
		argPos++
	}
	if p.Price != nil {
		setClauses = append(setClauses, "price_cents = $"+strconv.Itoa(argPos))
		args = append(args, p.Price.Cents)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil, ErrNoFieldsToUpdate
	}

	args = append(args, productID)
	query := `
		UPDATE products
		SET ` + strings.Join(setClauses, ", ") + `
		WHERE id = $` + strconv.Itoa(argPos) + `
		RETURNING id, sku, name, description, price_cents, is_active
	`

	row := r.db.QueryRow(ctx, query, args...)

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

func (r *ProductRepository) GetList(
	ctx context.Context,
	pager paginationShared.PagerInput,
	filter *product.ProductFilter,
	sort *product.ProductSort,
) (*product.GetListOutput, error) {
	query := `
		SELECT
			id,
			sku,
			name,
			description,
			price_cents,
			is_active,
			COUNT(*) OVER() as total_count
		FROM products
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	if filter != nil {
		if filter.SearchText != nil {
			query += ` AND (LOWER(name) LIKE $` + strconv.Itoa(argPos) + ` OR LOWER(sku) LIKE $` + strconv.Itoa(argPos) + `)`
			args = append(args, "%"+strings.ToLower(*filter.SearchText)+"%")
			argPos++
		}

		if filter.IsActive != nil {
			query += ` AND is_active = $` + strconv.Itoa(argPos)
			args = append(args, *filter.IsActive)
			argPos++
		}
	}

	if sort != nil && sort.Order.IsValid() && sort.Field.IsValid() {
		if *sort.Field == product.ProductSortFieldName {
			query += ` ORDER BY name ` + string(*sort.Order)
		}

		if *sort.Field == product.ProductSortFieldSKU {
			query += ` ORDER BY sku ` + string(*sort.Order)
		}

		if *sort.Field == product.ProductSortFieldPrice {
			query += ` ORDER BY price_cents ` + string(*sort.Order)
		}
	}

	query += ` LIMIT $` + strconv.Itoa(argPos)
	argPos++
	query += ` OFFSET $` + strconv.Itoa(argPos)

	offset := (pager.PageNumber - 1) * pager.PageSize
	args = append(args, pager.PageSize, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []product.Product{}
	var totalItems int

	for rows.Next() {
		var p product.Product
		var priceCents int

		if err := rows.Scan(
			&p.ID,
			&p.SKU,
			&p.Name,
			&p.Description,
			&priceCents,
			&p.IsActive,
			&totalItems,
		); err != nil {
			return nil, err
		}

		p.Price = product.Money{Cents: priceCents}
		products = append(products, p)
	}

	totalPages := (totalItems + pager.PageSize - 1) / pager.PageSize // Ceiling division
	pagerResults := paginationShared.PagerOutput{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pager.PageNumber,
		PageSize:    pager.PageSize,
	}

	return &product.GetListOutput{Products: &products, Pager: &pagerResults}, nil
}
