CREATE TABLE orders (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    order_status VARCHAR(50) NOT NULL,
    total_amount FLOAT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT fk_status_id FOREIGN KEY (status_id) REFERENCES order_status(id) ON DELETE CASCADE
);

CREATE TABLE order_items (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity INT NOT NULL,
    unit_price FLOAT NOT NULL,
    total_price FLOAT NOT NULL,
    created_at TIMESTAMPTZ
);
