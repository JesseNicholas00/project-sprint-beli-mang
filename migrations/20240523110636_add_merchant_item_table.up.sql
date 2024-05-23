CREATE TABLE merchant_items (
    merchant_item_id TEXT PRIMARY KEY,
    merchant_id TEXT,
    name TEXT,
    category TEXT,
    price INT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
