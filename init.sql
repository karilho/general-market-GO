-- Users info table, can be a buyer or a seller.
CREATE TABLE IF NOT EXISTS USER_DATA (
    user_id SERIAL PRIMARY KEY NOT NULL,
    user_type BOOLEAN,
    username VARCHAR(50) UNIQUE,
    email VARCHAR(255),
    password_hash VARCHAR(255),
    full_name VARCHAR(100),
    phone_number VARCHAR(20),
    registration_date DATE,
    street_address VARCHAR(255),
    place_number VARCHAR(10),
    city VARCHAR(100),
    state_province VARCHAR(100),
    postal_code VARCHAR(20),
    is_buyer BOOLEAN
    );

-- Buyer table, ppl who buy products, registered in purchase table
CREATE TABLE IF NOT EXISTS BUYERS (
    buyer_id SERIAL PRIMARY KEY NOT NULL,
    user_id INT,
    date_of_purchase DATE,
    payment_method VARCHAR(255),
    shipping_address VARCHAR(255),
    has_purchased BOOLEAN,
    FOREIGN KEY (user_id) REFERENCES USER_DATA(user_id)
    );

-- Seller table
-- Need some improvements
CREATE TABLE IF NOT EXISTS SELLERS (
    seller_id SERIAL PRIMARY KEY,
    user_id INT,
    seller_rating DECIMAL(3, 2),
    FOREIGN KEY (user_id) REFERENCES USER_DATA(user_id)
    );

-- Products table
CREATE TABLE IF NOT EXISTS PRODUCT (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255),
    product_value DECIMAL(10, 2),
    description TEXT
    );

-- Total value of a buy table
CREATE TABLE IF NOT EXISTS BUYER_ORDER (
    buyer_order_id SERIAL PRIMARY KEY,
    buyer_id INT,
    purchase_date DATE,
    total_value DECIMAL(10, 2),
    FOREIGN KEY (buyer_id) REFERENCES BUYERS(buyer_id)
    );

-- Each register on this table represent a product that was purchased, with the quantity.
-- You can see some register in the column purchase_id repeat sometimes.
-- This is because a purchase can have more than one product from different types, like a TV and a DVD.
-- Probably, you will see register with same purchase_id but different product_id.
CREATE TABLE IF NOT EXISTS PURCHASED_PRODUCTS (
    purchased_product_id SERIAL PRIMARY KEY,
    buyer_order_id INT,
    product_id INT,
    quantity INT,
    value_per_unit DECIMAL(10, 2),
    FOREIGN KEY (buyer_order_id) REFERENCES BUYER_ORDER(buyer_order_id),
    FOREIGN KEY (product_id) REFERENCES PRODUCT(product_id)
    );
