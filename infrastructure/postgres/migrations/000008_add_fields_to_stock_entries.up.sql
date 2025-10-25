ALTER TABLE stock_entries
    ADD COLUMN supplier_price_cents INTEGER,
    ADD COLUMN store_price_cents INTEGER,
    ADD COLUMN expiry_date DATE,
    ADD COLUMN reorder_date DATE;

COMMENT ON COLUMN stock_entries.supplier_price_cents IS 'Optional, for margin tracking';
COMMENT ON COLUMN stock_entries.store_price_cents IS 'Optional, can override product price';
COMMENT ON COLUMN stock_entries.expiry_date IS 'For perishable items';
COMMENT ON COLUMN stock_entries.reorder_date IS 'Reminder for restocking';