package factory

import (
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/warehouse"
	"paolojulian.dev/inventory/pkg/id"
	"paolojulian.dev/inventory/tests/middleware"
)

func NewTestWarehouse() *warehouse.Warehouse {
	return &warehouse.Warehouse{
		ID: id.NewUUID(),
		Location: warehouse.WarehouseLocation{
			Address: "123 Test St",
			City:    config.StringPointer("Test City"),
			Region:  config.StringPointer("Test Region"),
			Country: config.StringPointer("Testland"),
			Coordinates: &warehouse.Coordinates{
				Latitude:  12.345678,
				Longitude: 98.765432,
			},
		},
		IsActive:      true,
		CreatedBy:     middleware.FakeUserID,
		LastUpdatedBy: middleware.FakeUserID,
	}
}

func NewTestWarehouseSummary() *warehouse.WarehouseSummary {
	return &warehouse.WarehouseSummary{
		ID: id.NewUUID(),
		Location: warehouse.WarehouseLocation{
			Address: "123 Test St",
			City:    config.StringPointer("Test City"),
			Region:  config.StringPointer("Test Region"),
			Country: config.StringPointer("Testland"),
			Coordinates: &warehouse.Coordinates{
				Latitude:  12.345678,
				Longitude: 98.765432,
			},
		},
	}
}
