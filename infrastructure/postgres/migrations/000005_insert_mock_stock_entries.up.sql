-- Insert mock stock entries for testing
-- Using existing product IDs and getting the first available user ID
INSERT INTO stock_entries (id, product_id, warehouse_id, user_id, quantity_delta, reason, created_at) 
SELECT 
    '550e8400-e29b-41d4-a716-446655440100'::uuid,
    '550e8400-e29b-41d4-a716-446655440001'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    50,
    'restock',
    NOW() - INTERVAL '5 days'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440101'::uuid,
    '550e8400-e29b-41d4-a716-446655440001'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    -5,
    'sale',
    NOW() - INTERVAL '4 days'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440102'::uuid,
    '550e8400-e29b-41d4-a716-446655440002'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    100,
    'restock',
    NOW() - INTERVAL '3 days'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440103'::uuid,
    '550e8400-e29b-41d4-a716-446655440002'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    -10,
    'sale',
    NOW() - INTERVAL '2 days'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440104'::uuid,
    '550e8400-e29b-41d4-a716-446655440003'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    25,
    'restock',
    NOW() - INTERVAL '1 day'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440105'::uuid,
    '550e8400-e29b-41d4-a716-446655440003'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    -3,
    'sale',
    NOW() - INTERVAL '12 hours'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440106'::uuid,
    '550e8400-e29b-41d4-a716-446655440004'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    75,
    'restock',
    NOW() - INTERVAL '6 hours'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440107'::uuid,
    '550e8400-e29b-41d4-a716-446655440004'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    -2,
    'damage',
    NOW() - INTERVAL '3 hours'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440108'::uuid,
    '550e8400-e29b-41d4-a716-446655440005'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    30,
    'restock',
    NOW() - INTERVAL '2 hours'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440109'::uuid,
    '550e8400-e29b-41d4-a716-446655440005'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    -1,
    'sale',
    NOW() - INTERVAL '1 hour'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440110'::uuid,
    '550e8400-e29b-41d4-a716-446655440006'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    20,
    'restock',
    NOW() - INTERVAL '30 minutes'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440111'::uuid,
    '550e8400-e29b-41d4-a716-446655440007'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    15,
    'restock',
    NOW() - INTERVAL '15 minutes'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440112'::uuid,
    '550e8400-e29b-41d4-a716-446655440008'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    40,
    'restock',
    NOW() - INTERVAL '10 minutes'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440113'::uuid,
    '550e8400-e29b-41d4-a716-446655440009'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    60,
    'restock',
    NOW() - INTERVAL '5 minutes'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440114'::uuid,
    '550e8400-e29b-41d4-a716-446655440010'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    200,
    'restock',
    NOW() - INTERVAL '2 minutes'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440115'::uuid,
    '550e8400-e29b-41d4-a716-446655440001'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    10,
    'transfer_in',
    NOW() - INTERVAL '1 day'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440116'::uuid,
    '550e8400-e29b-41d4-a716-446655440002'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    -5,
    'transfer_out',
    NOW() - INTERVAL '18 hours'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440117'::uuid,
    '550e8400-e29b-41d4-a716-446655440003'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    2,
    'adjustment',
    NOW() - INTERVAL '8 hours'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440118'::uuid,
    '550e8400-e29b-41d4-a716-446655440004'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    -1,
    'adjustment',
    NOW() - INTERVAL '4 hours'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440119'::uuid,
    '550e8400-e29b-41d4-a716-446655440011'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    80,
    'restock',
    NOW() - INTERVAL '1 minute'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1)
UNION ALL
SELECT 
    '550e8400-e29b-41d4-a716-446655440120'::uuid,
    '550e8400-e29b-41d4-a716-446655440012'::uuid,
    (SELECT id FROM users LIMIT 1),
    (SELECT id FROM users LIMIT 1),
    35,
    'restock',
    NOW() - INTERVAL '30 seconds'
WHERE EXISTS (SELECT 1 FROM users LIMIT 1);
