CREATE TABLE order_status (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT
);

CREATE TABLE orders (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    status_id VARCHAR(36),
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