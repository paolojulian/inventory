-- Insert default warehouse
INSERT INTO warehouses (id, name, created_at, updated_at)
VALUES (
    '550e8400-e29b-41d4-a716-446655440000'::uuid,
    'Default',
    NOW(),
    NOW()
);

