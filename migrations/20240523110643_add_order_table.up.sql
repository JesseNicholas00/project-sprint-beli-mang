CREATE TABLE estimates (
    estimate_id TEXT PRIMARY KEY,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    order_id TEXT PRIMARY KEY,
    estimate_id TEXT
)