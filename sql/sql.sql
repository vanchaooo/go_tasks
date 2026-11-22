-- 1
    SELECT id, email, name
    FROM users
    WHERE email IS NOT NULL
          AND name IS NOT NULL
    ORDER BY id;


-- 2
    SELECT id, name, price_kopecks
    FROM products
    WHERE is_active = true
          AND stock > 0
    ORDER BY id;

-- 3
    SELECT id, name, price_kopecks
    FROM products
    WHERE price_kopecks BETWEEN 50000 and 100000
    ORDER BY price_kopecks;

-- 4
    SELECT id, email, city
    FROM users
    WHERE name = 'Анна'
    ORDER BY id;

-- 5
    SELECT id, name
    FROM users
    WHERE city IS NULL
    ORDER BY id

-- 6
    SELECT id, name, price_kopecks
    FROM products
    ORDER by price_kopecks DESC, id ASC
    LIMIT 3;

-- 7
    SELECT product_id, 
           quantity, 
           unit_price_kopecks, 
           quantity * unit_price_kopecks AS line_total_kopecks
    FROM order_items
    WHERE order_id = 1
    ORDER BY product_id;

-- 8
    SELECT id, created_at, status
    FROM orders
    WHERE created_at >= '2024-01-01 00:00:00+00' AND created_at < '2024-02-01 00:00:00+00'
    ORDER BY created_at, id;

-- 9
    SELECT id, created_at, canceled_at
    FROM orders
    WHERE canceled_at IS NOT NULL
    ORDER BY canceled_at, id;

-- 10
    SELECT id, name, city
    FROM users
    WHERE city <> 'Москва' OR city IS NULL
    ORDER BY id;

-- 11
    SELECT p.id AS product_id, p.name AS product_name, c.name AS category_name
    FROM products p
    JOIN categories c ON c.id = p.category_id
    ORDER BY p.id;

-- 12
    SELECT o.id AS order_id, u.id AS user_id, u.name AS user_name, o.created_at AS created_at
    FROM orders o
    JOIN users u ON u.id = o.user_id
    WHERE o.status = 'completed'
    ORDER BY o.created_at, o.id;

-- 13
    SELECT c.id AS category_id, c.name
    FROM categories C
    LEFT JOIN products p ON c.id = p.category_id
    WHERE p.id IS NULL
    ORDER BY c.id

-- 14
    SELECT o.id AS order_id, 
		count(i.id) AS line_count,
		coalesce(sum(i.quantity)) AS unit_count
    FROM orders o
    LEFT JOIN order_items i ON i.order_id = o.id
    GROUP BY o.id
    ORDER BY o.id

-- 15
    SELECT product_id, 
		count(*) AS review_count, 
		round(avg(rating)) AS avg_rating
    FROM reviews
    GROUP BY product_id
    ORDER BY avg(rating)