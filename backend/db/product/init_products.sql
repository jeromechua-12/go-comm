CREATE TABLE IF NOT EXISTS products (
    id              INTEGER PRIMARY KEY AUTO_INCREMENT,
    sku             VARCHAR(64) NOT NULL,
    category_id     INTEGER NOT NULL,
    brand_id        INTEGER NOT NULL,
    name            VARCHAR(255) NOT NULL,
    description     TEXT NOT NULL,
    specs           JSON NOT NULL,
    price_cents_usd INTEGER NOT NULL,
    quantity        INTEGER NOT NULL,
    image_url       TEXT NOT NULL,
    created_at      DATETIME NOT NULL,
    updated_at      DATETIME NOT NULL,
    CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES categories(id),
    CONSTRAINT fk_products_brand FOREIGN KEY (brand_id) REFERENCES brands(id)
);

ALTER TABLE products ADD CONSTRAINT uq_products_sku UNIQUE (sku);
ALTER TABLE products ADD CONSTRAINT uq_products_imageURL UNIQUE (image_url);

ALTER TABLE products ADD CONSTRAINT chk_products_price_positive CHECK (price_cents_usd > 0); 
ALTER TABLE products ADD CONSTRAINT chk_products_quantity_nonNegative CHECK (quantity >= 0); 
