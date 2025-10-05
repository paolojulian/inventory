CREATE TABLE
    IF NOT EXISTS stock_entries (
        id UUID PRIMARY KEY,
        product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
        warehouse_id UUID NOT NULL,
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        quantity_delta INTEGER NOT NULL,
        reason TEXT NOT NULL CHECK (reason IN ('restock', 'sale', 'damage', 'transfer_in', 'transfer_out', 'adjustment')),
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

-- Create index for better query performance
CREATE INDEX IF NOT EXISTS idx_stock_entries_product_id ON stock_entries(product_id);
CREATE INDEX IF NOT EXISTS idx_stock_entries_warehouse_id ON stock_entries(warehouse_id);
CREATE INDEX IF NOT EXISTS idx_stock_entries_user_id ON stock_entries(user_id);
CREATE INDEX IF NOT EXISTS idx_stock_entries_created_at ON stock_entries(created_at);

-- Create the trigger for stock_entries
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'set_updated_at_stock_entries'
    ) THEN
        CREATE TRIGGER set_updated_at_stock_entries
        BEFORE UPDATE ON stock_entries
        FOR EACH ROW
        EXECUTE FUNCTION trigger_set_updated_at();
    END IF;
END;
$$;
