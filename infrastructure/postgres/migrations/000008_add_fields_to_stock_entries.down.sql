ALTER TABLE stock_entries
    DROP COLUMN supplier_price_cents,
    DROP COLUMN store_price_cents,
    DROP COLUMN expiry_date,
    DROP COLUMN reorder_date;