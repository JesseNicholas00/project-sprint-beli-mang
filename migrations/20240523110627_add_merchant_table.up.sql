CREATE TABLE merchants (
    merchant_id TEXT PRIMARY KEY,
    name TEXT,
    category TEXT,
    image_url TEXT,
    -- location GEOGRAPHY(Point),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
