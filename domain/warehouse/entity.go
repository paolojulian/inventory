package warehouse

import (
	"time"

	"paolojulian.dev/inventory/pkg/id"
)

const (
	DefaultWarehouseID = "550e8400-e29b-41d4-a716-446655440000"
)

type Warehouse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WarehouseSummary struct {
	ID   string
	Name string
}

func NewWarehouse(name string) *Warehouse {
	now := time.Now()
	return &Warehouse{
		ID:        id.NewUUID(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
