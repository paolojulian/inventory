-- Insert initial admin user
-- Password: qwe123! (hashed with bcrypt cost 12)
INSERT INTO users (
    id, 
    username, 
    password, 
    role, 
    is_active, 
    email, 
    first_name, 
    last_name, 
    mobile
)
VALUES (
    gen_random_uuid(),
    'theman',
    '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYFxNDxFKXe', -- bcrypt hash of 'qwe123!'
    'admin',
    true,
    'paolojulian.personal@gmail.com',
    'Paolo Vincent',
    'Julian',
    '09279488654'
)
ON CONFLICT (username) DO NOTHING;