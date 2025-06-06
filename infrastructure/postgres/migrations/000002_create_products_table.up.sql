CREATE TABLE
    IF NOT EXISTS products (
        id UUID PRIMARY KEY,
        sku TEXT UNIQUE NOT NULL,
        name TEXT NOT NULL,
        description TEXT,
        price_cents INTEGER NOT NULL,
        is_active BOOLEAN NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

-- Create the trigger function once
CREATE OR REPLACE FUNCTION trigger_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger only if it doesn't exist for products
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'set_updated_at_products'
    ) THEN
        CREATE TRIGGER set_updated_at_products
        BEFORE UPDATE ON products
        FOR EACH ROW
        EXECUTE FUNCTION trigger_set_updated_at();
    END IF;
END;
$$;