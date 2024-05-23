CREATE TABLE orders (
    order_id TEXT PRIMARY KEY,
    user_id TEXT,
    merchant_id TEXT,
    completed BOOLEAN,
    -- location GEOGRAPHY(Point),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
