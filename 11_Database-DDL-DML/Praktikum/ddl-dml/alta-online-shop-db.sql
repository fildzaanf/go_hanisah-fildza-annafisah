-- 1
-- Creating Database Alta Online Shop
DROP DATABASE IF EXISTS alta_online_shop;
CREATE DATABASE alta_online_shop;

-- Creating Alta Online Shop Schema

USE alta_online_shop;

-- 2A

DROP TABLE IF EXISTS users;
CREATE TABLE users(
     user_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
     user_name VARCHAR(50) NOT NULL,
     user_address VARCHAR(100) NOT NULL,
     user_birthdate DATE,
     user_status VARCHAR(10) NOT NULL,
     user_gender VARCHAR(20) NOT NULL,
     user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     user_update_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2B
DROP TABLE IF EXISTS operators;
CREATE TABLE operators(
    operator_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    operator_name VARCHAR(50) NOT NULL,
    operator_contact VARCHAR(50) NOT NULL,
    operator_address VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS product_types;
CREATE TABLE product_types(
    product_type_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    product_name VARCHAR(50) NOT NULL,
    product_type_category VARCHAR(100) NOT NULL,
    product_stock INT(100) NOT NULL,
    product_price FLOAT NOT NULL,
);

DROP TABLE IF EXISTS product_descriptions;
CREATE TABLE product_descriptions(
    product_desc_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    product_type_id INT(10) NOT NULL,
    product_description VARCHAR(250) NOT NULL,

    FOREIGN KEY (product_type_id) REFERENCES product_types(product_type_id) ON DELETE CASCADE  
    ON UPDATE CASCADE
);

DROP TABLE IF EXISTS products;
CREATE TABLE products(
    product_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    operator_id INT(10) NOT NULL,
    product_type_id INT(10),
    product_desc_id INT(10),
    product_name VARCHAR(50) NOT NULL,

    FOREIGN KEY (operator_id) REFERENCES operators(operator_id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (product_type_id) REFERENCES product_types(product_type_id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (product_desc_id) REFERENCES product_descriptions(product_desc_id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
); 


DROP TABLE IF EXISTS payment_methods;
CREATE TABLE payment_methods(
    payment_method_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    payment_method_name VARCHAR(50) NOT NULL,
    payment_method_status VARCHAR(50) NOT NULL,
    payment_method_fee FLOAT NOT NULL
);

-- 2C

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions(
    transaction_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id INT(10) NOT NULL,
    payment_method_id INT(10) NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    transaction_status VARCHAR(10) NOT NULL,
    transaction_total_price INT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id) ON DELETE CASCADE  
    ON UPDATE CASCADE
);

DROP TABLE IF EXISTS transaction_details;
CREATE TABLE transaction_details(
    transaction_detail_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    transaction_id INT(10) NOT NULL,
    product_id INT(10) NOT NULL,
    product_type_id INT(10) NOT NULL,
    transaction_detail_price_product FLOAT NOT NULL,
    transaction_detail_quantity_product INT NOT NULL,

    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (product_type_id) REFERENCES product_types(product_type_id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id) ON DELETE CASCADE  
    ON UPDATE CASCADE

);

-- 3

DROP TABLE IF EXISTS couriers;
CREATE TABLE couriers(
    courier_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    courier_name VARCHAR(50) NOT NULL,
    courier_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    courier_update_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4 
ALTER TABLE couriers
ADD ongkos_dasar FLOAT NOT NULL
AFTER courier_name;

-- 5
ALTER TABLE couriers
RENAME TO shipping;

-- 6
DROP TABLE shipping;

-- 7A

DROP TABLE IF EXISTS payment_method_descriptions;
CREATE TABLE payment_method_descriptions(
    payment_method_desc_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    payment_method_id INT(10) NOT NULL,
    payment_method_desc_name VARCHAR(50) NOT NULL,
    payment_method_descriptions VARCHAR(200) NOT NULL,
    UNIQUE (payment_method_id),

    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id) ON DELETE CASCADE  
    ON UPDATE CASCADE
);

-- 7B

DROP TABLE IF EXISTS addresses;
CREATE TABLE addresses(
    address_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id INT(10) NOT NULL,
    address_street VARCHAR(50) NOT NULL,
    address_city VARCHAR(50) NOT NULL,
    address_state VARCHAR(50) NOT NULL,
    address_country VARCHAR(50) NOT NULL,
    address_postal_code INT(50) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE  
    ON UPDATE CASCADE
);

-- 7C

DROP TABLE IF EXISTS user_payment_method_details;
CREATE TABLE user_payment_method_details(
    user_payment_method_detail_id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    payment_method_id INT(10) NOT NULL,
    user_id INT(10) NOT NULL,  

    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE  
    ON UPDATE CASCADE
);

-- Desc Table

DESC users;
DESC products;
DESC product_types;
DESC operators;
DESC product_descriptions;
DESC payment_methods;
DESC transactions;
DESC transaction_details;
DESC payment_method_descriptions;
DESC addresses;
DESC user_payment_method_details;


