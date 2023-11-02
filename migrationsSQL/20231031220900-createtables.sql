-- +migrate Up
-- Users info table, can be a buyer or a seller.
-- Type for buyer / seller (not implemented)
CREATE TYPE user_type AS ENUM ('buyer', 'seller');
CREATE TABLE IF NOT EXISTS user_data (
    user_data_id SERIAL PRIMARY KEY NOT NULL,
    current_type user_type,
    username VARCHAR(50) UNIQUE,
    email VARCHAR(255),
    password_hash VARCHAR(255),
    full_name VARCHAR(100),
    phone_number VARCHAR(20),
    registration_date TIMESTAMP,
    street_address VARCHAR(255),
    place_number VARCHAR(10),
    city VARCHAR(100),
    state_province VARCHAR(100),
    postal_code VARCHAR(20)
    );

-- Buyer table, ppl who buy products, registered in purchase table
CREATE TABLE IF NOT EXISTS buyers (
    buyer_id SERIAL PRIMARY KEY NOT NULL,
    user_data_id INT,
    has_purchased BOOLEAN
    --FOREIGN KEY (user_data_id) REFERENCES user_data (user_data_id)
    );


-- Products table
CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255),
    product_value DECIMAL(10, 2),
    description TEXT
    );

-- Total value of a buy table
-- Each register on this table represent a purchase made by a buyer
-- 1 buyer can have more than one order
CREATE TABLE IF NOT EXISTS buy_order (
    buy_order_id UUID PRIMARY KEY,
    buyer_id INT,
    order_date TIMESTAMP,
    total_value DECIMAL(10, 2),
    payment_method VARCHAR(255)
    --FOREIGN KEY (buyer_id) REFERENCES buyers (buyer_id)
    );

-- Each register on this table represent a product that was purchased, with the quantity.
-- You can see some register in the column purchase_id repeat sometimes.
-- This is because a purchase can have more than one product from different types, like a TV and a DVD.
-- Probably, you will see register with same purchase_id but different product_id.
CREATE TABLE IF NOT EXISTS purchased_products (
    purchased_product_id SERIAL PRIMARY KEY,
    buy_order_id INT,
    product_id INT,
    quantity INT,
    value_per_unit DECIMAL(10, 2)
    --FOREIGN KEY (buy_order_id) REFERENCES buy_order (buy_order_id),
    --FOREIGN KEY (product_id) REFERENCES products (product_id)
    );


-- Seller table
-- Need some improvements, so right now i will not implement.
--CREATE TABLE IF NOT EXISTS sellers (
--    seller_id SERIAL PRIMARY KEY NOT NULL,
--    user_id INT,
--    FOREIGN KEY (user_id) REFERENCES USER_DATA(user_id)
--    );


-- Just for testing
CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name VARCHAR(50) NOT NULL,
       email VARCHAR(50) NOT NULL
);