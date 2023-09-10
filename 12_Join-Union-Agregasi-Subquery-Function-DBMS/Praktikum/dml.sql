-- Creating Database Alta DB
DROP DATABASE IF EXISTS alta_db;
CREATE DATABASE alta_db;

-- Creating Alta DB Schema

USE alta_db;

-- table users 
DROP TABLE IF EXISTS users;
CREATE TABLE users(
     id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
     name VARCHAR(100),
     status SMALLINT,
     dob DATE,
     gender CHAR(1),
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- table operators
DROP TABLE IF EXISTS operators;
CREATE TABLE operators(
    id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- table product types
DROP TABLE IF EXISTS product_types;
CREATE TABLE product_types(
    id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- table products
DROP TABLE IF EXISTS products;
CREATE TABLE products(
    id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    product_type_id INT(11),
    operator_id INT(11),
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,


    FOREIGN KEY (operator_id) REFERENCES operators(id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (product_type_id) REFERENCES product_types(id) ON DELETE CASCADE  
    ON UPDATE CASCADE
); 

-- table product descriptions
DROP TABLE IF EXISTS product_descriptions;
CREATE TABLE product_descriptions(
    id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- table payment methods
DROP TABLE IF EXISTS payment_methods;
CREATE TABLE payment_methods(
    id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- table transactions
DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions(
    id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id INT(11),
    payment_method_id INT(11),
    status VARCHAR(10),
    total_qty INT(11),
    total_price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE CASCADE  
    ON UPDATE CASCADE
);

-- table transaction details
DROP TABLE IF EXISTS transaction_details;
CREATE TABLE transaction_details(
    id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    transaction_id INT(11),
    product_id INT(11),
    status VARCHAR(10),
    qty INT(11),
    price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE  
    ON UPDATE CASCADE,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE  
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

-- INSERT


-- 1A
-- Insert 5 operators pada table operators
INSERT INTO operators (name) VALUES
    ('Operator 1'),
    ('Operator 2'),
    ('Operator 3'),
    ('Operator 4'),
    ('Operator 5');

-- 1B
-- Insert 3 product type

INSERT INTO product_types (name)
VALUES
    ('Product 1'),
    ('Product 2'),
    ('Product 3');

-- 1C
-- Insert 2 product dengan product type id = 1, dan operators id = 3

INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES
    (1, 3, 'P1-001', 'Product 1', 1),
    (1, 3, 'P1-002', 'Product 2', 1);

-- 1D
-- Insert 3 product dengan product type id = 2, dan operators id = 1

INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES
    (2, 1, 'P2-001', 'Product 1', 1),
    (2, 1, 'P2-002', 'Product 2', 1),
    (2, 1, 'P2-003', 'Product 3', 1);

-- 1E
-- Insert 3 product dengan product type id = 3, dan operators id = 4

INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES
    (3, 4, 'P3-001', 'Product 1', 1),
    (3, 4, 'P3-002', 'Product 2', 1),
    (3, 4, 'P3-003', 'Product 3', 1);

-- 1F
-- Insert product description pada setiap product

INSERT INTO product_descriptions (description)
VALUES
    ('Deskripsi untuk Product 1 dengan ID 1'),
    ('Deskripsi untuk Product 2 dengan ID 2'),
    ('Deskripsi untuk Product 3 dengan ID 3'),
    ('Deskripsi untuk Product 4 dengan ID 4'),
    ('Deskripsi untuk Product 5 dengan ID 5');

-- 1G
-- Insert 3 payment methods
INSERT INTO payment_methods (name, status)
VALUES
    ('Payment Method 1', 1),
    ('Payment Method 2', 1),
    ('Payment Method 3', 1);

-- 1H
-- Insert 5 user pada tabel user
INSERT INTO users (name, status, dob, gender)
VALUES
    ('John Doe', 1, '2001-01-01', 'M'),
    ('Jane Smith', 1, '1995-05-15', 'F'),
    ('Michael Johnson', 1, '1980-11-30', 'M'),
    ('Emily Davis', 1, '1990-07-10', 'F'),
    ('David Wilson', 1, '1988-10-25', 'M');

-- 1I
-- Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES
    -- Transaksi untuk User ID 1
    (1, 1, 'Completed', 3, 1000),
    (1, 2, 'Pending', 2, 750),
    (1, 3, 'Completed', 14, 1525),

    -- Transaksi untuk User ID 2
    (2, 1, 'Pending', 2, 6000),
    (2, 2, 'Completed', 11, 3575),
    (2, 3, 'Pending', 3, 9520),

    -- Transaksi untuk User ID 3
    (3, 1, 'Completed', 5, 18050),
    (3, 2, 'Completed', 2, 7500),
    (3, 3, 'Pending', 4, 1475),

    -- Transaksi untuk User ID 4
    (4, 1, 'Pending', 1, 4525),
    (4, 2, 'Completed', 3, 11000),
    (4, 3, 'Completed', 12, 7040),

    -- Transaksi untuk User ID 5
    (5, 1, 'Completed', 4, 15575),
    (5, 2, 'Pending', 21, 7000),
    (5, 3, 'Completed', 1, 4020);


-- 1J
-- Insert 3 product di masing-masing transaksi

-- Transaksi 1
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (1, 1, 'Completed', 2, 5000),
    (1, 2, 'Completed', 1, 2500),
    (1, 3, 'Completed', 3, 7500);

-- Transaksi 2
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (2, 4, 'Pending', 1, 3000),
    (2, 5, 'Pending', 1, 3000),
    (2, 6, 'Pending', 1, 3000);

-- Transaksi 3
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (3, 1, 'Completed', 2, 5000),
    (3, 2, 'Completed', 1, 2500),
    (3, 3, 'Completed', 2, 5000);

-- Transaksi 4
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (4, 4, 'Pending', 3, 7500),
    (4, 5, 'Pending', 1, 2500),
    (4, 6, 'Pending', 2, 5000);

-- Transaksi 5
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (5, 1, 'Completed', 1, 2500),
    (5, 2, 'Completed', 3, 7500),
    (5, 3, 'Completed', 2, 500);

-- Transaksi 6
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (6, 4, 'Pending', 2, 6000),
    (6, 5, 'Completed', 1, 3500),
    (6, 6, 'Pending', 3, 9500);

-- Transaksi 7
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (7, 1, 'Completed', 4, 18000),
    (7, 2, 'Completed', 1, 7000),
    (7, 3, 'Pending', 2, 1400);

-- Transaksi 8
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (8, 4, 'Pending', 1, 4500),
    (8, 5, 'Completed', 3, 11500),
    (8, 6, 'Completed', 2, 7000);

-- Transaksi 9
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (9, 1, 'Completed', 2, 5500),
    (9, 2, 'Completed', 1, 2400),
    (9, 3, 'Pending', 4, 1600);

-- Transaksi 10
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (10, 4, 'Pending', 3, 9500),
    (10, 5, 'Completed', 1, 3100),
    (10, 6, 'Completed', 2, 4800);

-- Transaksi 11
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (11, 1, 'Pending', 1, 3000),
    (11, 2, 'Completed', 4, 10000),
    (11, 3, 'Pending', 2, 1200);

-- Transaksi 12
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (12, 4, 'Completed', 3, 7500),
    (12, 5, 'Completed', 1, 3200),
    (12, 6, 'Pending', 2, 6000);

-- Transaksi 13
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (13, 1, 'Pending', 1, 5000),
    (13, 2, 'Completed', 3, 9200),
    (13, 3, 'Pending', 2, 4800);

-- Transaksi 14
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (14, 4, 'Completed', 2, 5200),
    (14, 5, 'Pending', 1, 2500),
    (14, 6, 'Completed', 4, 9000);

-- Transaksi 15
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
    (15, 1, 'Pending', 3, 7500),
    (15, 2, 'Completed', 1, 4200),
    (15, 3, 'Completed', 2, 6500);


-- SELECT

-- 2A
--test
SELECT * FROM users;
-- answer
SELECT * FROM users WHERE gender = 'M';

-- 2B
--test
SELECT * FROM products;
-- answer
SELECT * FROM products WHERE id = 3;

-- 2C
--test
SELECT * FROM users;
-- answer
SELECT * FROM users WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND name LIKE '%a%';

-- 2D
--test
SELECT * FROM users;
-- answer
SELECT COUNT(*) FROM users WHERE gender = 'F';

-- 2E
SELECT * FROM users;
-- answer
SELECT * FROM users ORDER BY name ASC;

-- 2F
--test
SELECT * FROM products;
-- answer
SELECT * FROM products LIMIT 5;


-- UPDATE

-- 3A
--test 
SELECT * FROM products;
-- answer
UPDATE products
SET name = 'product dummy'
WHERE id = 1;
--test
SELECT * FROM products;

-- 3B
--test
SELECT * FROM transaction_details WHERE product_id = 1;
-- answer
UPDATE transaction_details
SET qty = 3
WHERE product_id = 1;
--test
SELECT * FROM transaction_details WHERE product_id = 1;

-- DELETE 

-- 4A
--test
SELECT * FROM products;
-- answer
DELETE FROM products
WHERE id = 1;
--test
SELECT * FROM products;

-- 4B
--test
SELECT * FROM products;
-- answer
DELETE FROM products
WHERE product_type_id = 1;
--test
SELECT * FROM products;

-- Join, Union, Sub query, Function

-- 1
-- test
SELECT * FROM transactions;
-- answer
SELECT *
FROM transactions
WHERE user_id = 1
UNION 
SELECT *
FROM transactions
WHERE user_id = 2;

-- 2
-- test
SELECT * FROM transactions;
-- answer
SELECT user_id, SUM(total_price) AS sum_total_price
FROM transactions
WHERE user_id = 1;

-- 3
-- answer
SELECT 
product_types.id AS product_type_id, 
COUNT(*) AS total_transaction
FROM transactions
INNER JOIN transaction_details 
ON transactions.id = transaction_details.transaction_id
INNER JOIN products 
ON transaction_details.product_id = products.id
INNER JOIN product_types 
ON products.product_type_id = product_types.id
WHERE products.product_type_id = 2;

-- 4
-- test
SELECT * FROM products; 
SELECT * FROM product_types; 
-- answer
SELECT products.*, 
product_types.name AS product_type_name
FROM products
LEFT JOIN product_types 
ON products.product_type_id = product_types.id;

-- 5
SELECT transactions.*, 
products.name AS product_name, 
users.name AS user_name
FROM transactions
INNER JOIN transaction_details 
ON transactions.id = transaction_details.transaction_id
INNER JOIN products 
ON transaction_details.product_id = products.id
INNER JOIN users 
ON transactions.user_id = users.id;

-- 6
DELIMITER //

CREATE TRIGGER delete_after_transaction
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transaction_details 
    WHERE transaction_id = OLD.id;
END;

//
DELIMITER ;

-- test 6
SELECT * FROM TRANSACTION_DETAILS;
SELECT * FROM TRANSACTIONS;

DELETE FROM transactions WHERE id = 1;

SELECT * FROM TRANSACTION_DETAILS;
SELECT * FROM TRANSACTIONS;

-- 7

DELIMITER //

CREATE TRIGGER update_total_qty
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
  
    DECLARE total_quantity INT(11);

    SELECT total_qty
    INTO total_quantity
    FROM transactions
    WHERE id = OLD.transaction_id;

    UPDATE transactions
    SET total_qty = total_quantity - OLD.qty
    WHERE id = OLD.transaction_id;

END;
//
DELIMITER ;

-- test 7
SELECT * FROM TRANSACTION_DETAILS;
SELECT * FROM TRANSACTIONS;

DELETE FROM transaction_details WHERE id = 1;

SELECT * FROM TRANSACTION_DETAILS;
SELECT * FROM TRANSACTIONS;

-- 8
-- test
SELECT * FROM products; 
SELECT DISTINCT product_id FROM transaction_details;
-- answer
SELECT *
FROM products
WHERE id 
NOT IN (SELECT DISTINCT product_id FROM transaction_details);




