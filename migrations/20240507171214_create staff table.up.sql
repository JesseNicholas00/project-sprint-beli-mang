CREATE TABLE staffs (
    staff_id TEXT PRIMARY KEY,
    staff_name TEXT,
    staff_phone_number TEXT,
    staff_password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
