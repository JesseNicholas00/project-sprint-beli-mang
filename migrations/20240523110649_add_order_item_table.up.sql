CREATE TABLE estimate_items (
    estimate_item_id BIGSERIAL PRIMARY KEY,
    estimate_id TEXT,
    merchant_id TEXT,
    merchant_item_id TEXT,
    quantity INT
);
