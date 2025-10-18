CREATE TABLE
    IF NOT EXISTS warehouses (
        id UUID PRIMARY KEY,
        name TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

-- Create index for better query performance
CREATE INDEX IF NOT EXISTS idx_warehouses_name ON warehouses(name);

-- Create the trigger for warehouses
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'set_updated_at_warehouses'
    ) THEN
        CREATE TRIGGER set_updated_at_warehouses
        BEFORE UPDATE ON warehouses
        FOR EACH ROW
        EXECUTE FUNCTION trigger_set_updated_at();
    END IF;
END;
$$;

