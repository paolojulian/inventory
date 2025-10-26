-- Insert initial admin user
-- Password: qwe123! (hashed with bcrypt cost 12)
INSERT INTO
    users (
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
VALUES
    (
        gen_random_uuid (),
        'admin',
        '$2a$12$cJN9BiLwsrRxlJkpurKnje.zOGz.7kNHqjdxh65euyLdaZ7auLiJG', -- bcrypt hash of 'qwe123!'
        'admin',
        true,
        'paolojulian.personal@gmail.com',
        'Paolo Vincent',
        'Julian',
        '09279488654'
    ) ON CONFLICT (username) DO NOTHING;