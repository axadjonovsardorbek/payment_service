CREATE TYPE IF NOT EXISTS payment_method AS ENUM('naqd', 'karta', 'online');
CREATE TYPE IF NOT EXISTS payment_status AS ENUM('tasdiqlangan', 'bekor qilingan', "to'langan");

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,
    reservation_id UUID,
    amount DECIMAL,
    payment_method payment_method,
    payment_status payment_status,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);