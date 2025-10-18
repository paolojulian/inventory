package warehouse

import (
	"time"

	"paolojulian.dev/inventory/pkg/id"
)

type Warehouse struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
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