package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/stock"
	userDomain "paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/pkg/id"
)

// setupTestData creates necessary test data (warehouse, product, user) for stock entry tests
func setupTestData(t *testing.T, ctx context.Context, db *pgxpool.Pool) (warehouseID, productID, userID string) {
	t.Helper()

	// Use the default warehouse from migrations
	warehouseRepo := postgres.NewWarehouseRepository(db)
	defaultWarehouse, err := warehouseRepo.GetDefaultWarehouse(ctx)
	require.NoError(t, err, "failed to get default warehouse")

	// Create test product
	productRepo := postgres.NewProductRepository(db)
	sku, err := product.NewSKU("TEST-" + id.NewUUID()[:8])
	require.NoError(t, err, "failed to create SKU")
	
	desc, err := product.NewDescription("Test product for stock entry")
	require.NoError(t, err, "failed to create description")

	testProduct := &product.Product{
		ID:          id.NewUUID(),
		SKU:         sku,
		Name:        "Test Product",
		Description: desc,
		Price:       product.Money{Cents: 1000},
		IsActive:    true,
	}
	createdProduct, err := productRepo.Save(ctx, testProduct)
	require.NoError(t, err, "failed to create test product")

	// Create test user
	userRepo := postgres.NewUserRepository(db)
	email := "test@example.com"
	firstName := "Test"
	lastName := "User"
	mobile := "1234567890"
	
	testUser, err := userDomain.NewUser(
		"testuser-"+id.NewUUID()[:8],
		"password123",
		userDomain.AdminRole,
		true,
		&email,
		&firstName,
		&lastName,
		&mobile,
	)
	require.NoError(t, err, "failed to create user domain object")
	createdUser, err := userRepo.Save(ctx, testUser)
	require.NoError(t, err, "failed to create test user")

	return defaultWarehouse.ID, createdProduct.ID, createdUser.ID
}

// cleanupTestData removes test data created during the test
func cleanupTestData(t *testing.T, ctx context.Context, db *pgxpool.Pool, stockEntryID, productID, userID string) {
	t.Helper()

	// Delete stock entry
	if stockEntryID != "" {
		stockRepo := postgres.NewStockRepository(db)
		_ = stockRepo.Delete(ctx, stockEntryID)
	}

	// Delete product (cascade should handle stock entries)
	if productID != "" {
		productRepo := postgres.NewProductRepository(db)
		_ = productRepo.Delete(ctx, productID)
	}

	// Delete user via direct SQL (no Delete method in repository)
	if userID != "" {
		_, _ = db.Exec(ctx, "DELETE FROM users WHERE id = $1", userID)
	}
}

func TestStockRepository_CreateStockEntry_Success(t *testing.T) {
	// Setup
	db, err := postgres.NewPool()
	require.NoError(t, err, "failed to connect to database")
	defer db.Close()

	ctx := context.Background()
	warehouseID, productID, userID := setupTestData(t, ctx, db)

	repo := postgres.NewStockRepository(db)

	// Create stock entry with all fields
	supplierPrice := 800
	storePrice := 1200
	expiryDate := time.Now().AddDate(0, 6, 0)
	reorderDate := time.Now().AddDate(0, 5, 0)

	stockEntry := &stock.StockEntry{
		ID:                 id.NewUUID(),
		ProductID:          productID,
		WarehouseID:        warehouseID,
		UserID:             userID,
		QuantityDelta:      50,
		Reason:             stock.ReasonRestock,
		SupplierPriceCents: &supplierPrice,
		StorePriceCents:    &storePrice,
		ExpiryDate:         &expiryDate,
		ReorderDate:        &reorderDate,
		CreatedAt:          time.Now(),
	}

	// Execute
	created, err := repo.CreateStockEntry(ctx, stockEntry)

	// Cleanup
	defer cleanupTestData(t, ctx, db, created.ID, productID, userID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, stockEntry.ID, created.ID)
	assert.Equal(t, stockEntry.ProductID, created.ProductID)
	assert.Equal(t, stockEntry.WarehouseID, created.WarehouseID)
	assert.Equal(t, stockEntry.UserID, created.UserID)
	assert.Equal(t, stockEntry.QuantityDelta, created.QuantityDelta)
	assert.Equal(t, stockEntry.Reason, created.Reason)
	assert.NotNil(t, created.SupplierPriceCents)
	assert.Equal(t, *stockEntry.SupplierPriceCents, *created.SupplierPriceCents)
	assert.NotNil(t, created.StorePriceCents)
	assert.Equal(t, *stockEntry.StorePriceCents, *created.StorePriceCents)
	assert.NotNil(t, created.ExpiryDate)
	assert.NotNil(t, created.ReorderDate)
	assert.False(t, created.CreatedAt.IsZero())
}

func TestStockRepository_CreateStockEntry_MinimalFields(t *testing.T) {
	// Setup
	db, err := postgres.NewPool()
	require.NoError(t, err, "failed to connect to database")
	defer db.Close()

	ctx := context.Background()
	warehouseID, productID, userID := setupTestData(t, ctx, db)

	repo := postgres.NewStockRepository(db)

	// Create stock entry with only required fields
	stockEntry := &stock.StockEntry{
		ID:            id.NewUUID(),
		ProductID:     productID,
		WarehouseID:   warehouseID,
		UserID:        userID,
		QuantityDelta: -10,
		Reason:        stock.ReasonSale,
		CreatedAt:     time.Now(),
	}

	// Execute
	created, err := repo.CreateStockEntry(ctx, stockEntry)

	// Cleanup
	defer cleanupTestData(t, ctx, db, created.ID, productID, userID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, stockEntry.ID, created.ID)
	assert.Equal(t, stockEntry.ProductID, created.ProductID)
	assert.Equal(t, stockEntry.WarehouseID, created.WarehouseID)
	assert.Equal(t, stockEntry.UserID, created.UserID)
	assert.Equal(t, stockEntry.QuantityDelta, created.QuantityDelta)
	assert.Equal(t, stockEntry.Reason, created.Reason)
	assert.Nil(t, created.SupplierPriceCents)
	assert.Nil(t, created.StorePriceCents)
	assert.Nil(t, created.ExpiryDate)
	assert.Nil(t, created.ReorderDate)
}

func TestStockRepository_CreateStockEntry_InvalidProductID(t *testing.T) {
	// Setup
	db, err := postgres.NewPool()
	require.NoError(t, err, "failed to connect to database")
	defer db.Close()

	ctx := context.Background()
	warehouseID, productID, userID := setupTestData(t, ctx, db)
	defer cleanupTestData(t, ctx, db, "", productID, userID)

	repo := postgres.NewStockRepository(db)

	// Create stock entry with non-existent product ID
	stockEntry := &stock.StockEntry{
		ID:            id.NewUUID(),
		ProductID:     id.NewUUID(), // Non-existent product
		WarehouseID:   warehouseID,
		UserID:        userID,
		QuantityDelta: 10,
		Reason:        stock.ReasonRestock,
		CreatedAt:     time.Now(),
	}

	// Execute
	created, err := repo.CreateStockEntry(ctx, stockEntry)

	// Assert - should fail due to foreign key constraint
	assert.Error(t, err)
	assert.Nil(t, created)
}

func TestStockRepository_CreateStockEntry_InvalidWarehouseID(t *testing.T) {
	// Setup
	db, err := postgres.NewPool()
	require.NoError(t, err, "failed to connect to database")
	defer db.Close()

	ctx := context.Background()
	_, productID, userID := setupTestData(t, ctx, db)
	defer cleanupTestData(t, ctx, db, "", productID, userID)

	repo := postgres.NewStockRepository(db)

	// Create stock entry with non-existent warehouse ID
	stockEntry := &stock.StockEntry{
		ID:            id.NewUUID(),
		ProductID:     productID,
		WarehouseID:   id.NewUUID(), // Non-existent warehouse
		UserID:        userID,
		QuantityDelta: 10,
		Reason:        stock.ReasonRestock,
		CreatedAt:     time.Now(),
	}

	// Execute
	created, err := repo.CreateStockEntry(ctx, stockEntry)

	// Assert - should fail due to foreign key constraint
	assert.Error(t, err)
	assert.Nil(t, created)
}

func TestStockRepository_CreateStockEntry_InvalidUserID(t *testing.T) {
	// Setup
	db, err := postgres.NewPool()
	require.NoError(t, err, "failed to connect to database")
	defer db.Close()

	ctx := context.Background()
	warehouseID, productID, userID := setupTestData(t, ctx, db)
	defer cleanupTestData(t, ctx, db, "", productID, userID)

	repo := postgres.NewStockRepository(db)

	// Create stock entry with non-existent user ID
	stockEntry := &stock.StockEntry{
		ID:            id.NewUUID(),
		ProductID:     productID,
		WarehouseID:   warehouseID,
		UserID:        id.NewUUID(), // Non-existent user
		QuantityDelta: 10,
		Reason:        stock.ReasonRestock,
		CreatedAt:     time.Now(),
	}

	// Execute
	created, err := repo.CreateStockEntry(ctx, stockEntry)

	// Assert - should fail due to foreign key constraint
	assert.Error(t, err)
	assert.Nil(t, created)
}

func TestStockRepository_CreateStockEntry_DifferentReasons(t *testing.T) {
	// Setup
	db, err := postgres.NewPool()
	require.NoError(t, err, "failed to connect to database")
	defer db.Close()

	ctx := context.Background()
	warehouseID, productID, userID := setupTestData(t, ctx, db)
	defer cleanupTestData(t, ctx, db, "", productID, userID)

	repo := postgres.NewStockRepository(db)

	// Test all valid reasons
	testCases := []struct {
		name   string
		reason stock.StockReason
	}{
		{"Restock", stock.ReasonRestock},
		{"Sale", stock.ReasonSale},
		{"Damage", stock.ReasonDamage},
		{"Transfer In", stock.ReasonTransferIn},
		{"Transfer Out", stock.ReasonTransferOut},
		{"Adjustment", stock.ReasonAdjustment},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stockEntry := &stock.StockEntry{
				ID:            id.NewUUID(),
				ProductID:     productID,
				WarehouseID:   warehouseID,
				UserID:        userID,
				QuantityDelta: 5,
				Reason:        tc.reason,
				CreatedAt:     time.Now(),
			}

			created, err := repo.CreateStockEntry(ctx, stockEntry)

			// Cleanup
			if created != nil {
				defer repo.Delete(ctx, created.ID)
			}

			// Assert
			require.NoError(t, err)
			assert.NotNil(t, created)
			assert.Equal(t, tc.reason, created.Reason)
		})
	}
}

func TestStockRepository_CreateStockEntry_NegativeQuantity(t *testing.T) {
	// Setup
	db, err := postgres.NewPool()
	require.NoError(t, err, "failed to connect to database")
	defer db.Close()

	ctx := context.Background()
	warehouseID, productID, userID := setupTestData(t, ctx, db)

	repo := postgres.NewStockRepository(db)

	// Create stock entry with negative quantity (valid for sales/damage)
	stockEntry := &stock.StockEntry{
		ID:            id.NewUUID(),
		ProductID:     productID,
		WarehouseID:   warehouseID,
		UserID:        userID,
		QuantityDelta: -15,
		Reason:        stock.ReasonDamage,
		CreatedAt:     time.Now(),
	}

	// Execute
	created, err := repo.CreateStockEntry(ctx, stockEntry)

	// Cleanup
	defer cleanupTestData(t, ctx, db, created.ID, productID, userID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, -15, created.QuantityDelta)
	assert.Equal(t, stock.ReasonDamage, created.Reason)
}
