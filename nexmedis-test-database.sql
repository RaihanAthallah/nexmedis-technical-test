-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION pg_database_owner;

COMMENT ON SCHEMA public IS 'standard public schema';

-- DROP SEQUENCE public.accounts_id_seq;

CREATE SEQUENCE public.accounts_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.cart_items_id_seq;

CREATE SEQUENCE public.cart_items_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.order_details_id_seq;

CREATE SEQUENCE public.order_details_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.orders_id_seq;

CREATE SEQUENCE public.orders_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.products_id_seq;

CREATE SEQUENCE public.products_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.users_id_seq;

CREATE SEQUENCE public.users_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;-- public.products definition

-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	id serial4 NOT NULL,
	"name" varchar(255) NOT NULL,
	description text NULL,
	price numeric(10, 2) NOT NULL,
	stock int4 NOT NULL,
	category varchar(100) NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id),
	CONSTRAINT products_price_check CHECK ((price >= (0)::numeric)),
	CONSTRAINT products_stock_check CHECK ((stock >= 0))
);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id int4 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE) NOT NULL,
	username varchar(255) NOT NULL,
	"password" text NOT NULL,
	email varchar(255) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_key UNIQUE (username)
);


-- public.accounts definition

-- Drop table

-- DROP TABLE public.accounts;

CREATE TABLE public.accounts (
	id serial4 NOT NULL,
	user_id int4 NOT NULL,
	balance numeric(15, 2) DEFAULT 0 NOT NULL,
	CONSTRAINT accounts_pkey PRIMARY KEY (id),
	CONSTRAINT accounts_user_id_key UNIQUE (user_id),
	CONSTRAINT accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.cart_items definition

-- Drop table

-- DROP TABLE public.cart_items;

CREATE TABLE public.cart_items (
	id serial4 NOT NULL,
	user_id int4 NOT NULL,
	product_id int4 NOT NULL,
	quantity int4 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT cart_items_pkey PRIMARY KEY (id),
	CONSTRAINT cart_items_quantity_check CHECK ((quantity > 0)),
	CONSTRAINT cart_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE,
	CONSTRAINT cart_items_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.orders definition

-- Drop table

-- DROP TABLE public.orders;

CREATE TABLE public.orders (
	id serial4 NOT NULL,
	user_id int4 NOT NULL,
	total_price numeric(10, 2) NOT NULL,
	status varchar(50) DEFAULT 'pending'::character varying NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id),
	CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.order_details definition

-- Drop table

-- DROP TABLE public.order_details;

CREATE TABLE public.order_details (
	id serial4 NOT NULL,
	order_id int4 NOT NULL,
	product_id int4 NOT NULL,
	quantity int4 NOT NULL,
	price numeric(10, 2) NOT NULL,
	subtotal numeric(10, 2) GENERATED ALWAYS AS ((quantity::numeric * price)) STORED NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT order_details_pkey PRIMARY KEY (id),
	CONSTRAINT order_details_quantity_check CHECK ((quantity > 0)),
	CONSTRAINT order_details_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE,
	CONSTRAINT order_details_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE
);

--SELECT quantity FROM cart_items WHERE user_id = 1 AND product_id = 2;
--INSERT INTO accounts (user_id, balance) VALUES
--(1, 5000.00);


INSERT INTO orders (user_id, total_price, status, created_at, updated_at) VALUES
(1, 150.75, 'completed', '2024-06-10 10:15:30', '2024-06-10 10:15:30'),
(2, 89.99, 'pending', '2024-06-11 12:45:20', '2024-06-11 12:45:20'),
(3, 210.50, 'completed', '2024-06-12 14:20:10', '2024-06-12 14:20:10'),
(4, 45.00, 'cancelled', '2024-06-13 16:05:40', '2024-06-13 16:05:40'),
(5, 399.99, 'shipped', '2024-06-14 18:30:50', '2024-06-14 18:30:50'),
(6, 120.25, 'completed', '2024-06-15 09:10:05', '2024-06-15 09:10:05'),
(1, 75.60, 'completed', '2024-06-16 11:50:25', '2024-06-16 11:50:25'),
(2, 99.99, 'pending', '2024-06-17 13:35:55', '2024-06-17 13:35:55'),
(3, 250.00, 'shipped', '2024-06-18 15:00:45', '2024-06-18 15:00:45'),
(4, 30.99, 'completed', '2024-06-19 17:20:30', '2024-06-19 17:20:30'),
(5, 189.75, 'pending', '2024-06-20 19:45:15', '2024-06-20 19:45:15'),
(6, 300.50, 'completed', '2024-06-21 21:10:50', '2024-06-21 21:10:50');


INSERT INTO users (username, password, email) VALUES
('alice123', 'hashed_password1', 'alice@example.com'),
('bob456', 'hashed_password2', 'bob@example.com'),
('charlie789', 'hashed_password3', 'charlie@example.com'),
('david999', 'hashed_password4', 'david@example.com'),
('eve777', 'hashed_password5', 'eve@example.com')
RETURNING id;

INSERT INTO users (username, password, email) VALUES
('alice_j', 'hashed_password_1', 'alice.johnson@example.com'),
('bob_s', 'hashed_password_2', 'bob.smith@example.com'),
('charlie_b', 'hashed_password_3', 'charlie.brown@example.com'),
('david_w', 'hashed_password_4', 'david.wilson@example.com'),
('emma_d', 'hashed_password_5', 'emma.davis@example.com'),
('frank_m', 'hashed_password_6', 'frank.miller@example.com'),
('grace_l', 'hashed_password_7', 'grace.lee@example.com'),
('henry_c', 'hashed_password_8', 'henry.clark@example.com'),
('isabella_w', 'hashed_password_9', 'isabella.white@example.com'),
('jack_t', 'hashed_password_10', 'jack.thompson@example.com');


INSERT INTO products (name, description, price, stock, category, created_at) VALUES
('Apple iPhone 15 Pro', 'Latest iPhone with A17 Pro chip and titanium body', 1199.99, 50, 'Smartphones', NOW()),
('Samsung Galaxy S23 Ultra', 'Flagship Samsung phone with S Pen and 200MP camera', 1099.99, 30, 'Smartphones', NOW()),
('Dell XPS 15', 'High-performance laptop with Intel Core i9 and 32GB RAM', 1899.99, 20, 'Laptops', NOW()),
('MacBook Air M2', 'Ultra-thin laptop with Apple M2 chip and 18-hour battery', 1299.99, 40, 'Laptops', NOW()),
('Sony WH-1000XM5', 'Noise-canceling headphones with superior sound quality', 399.99, 60, 'Accessories', NOW()),
('Logitech MX Master 3', 'Ergonomic wireless mouse with precision tracking', 99.99, 100, 'Accessories', NOW()),
('PlayStation 5', 'Next-gen gaming console with ultra-fast SSD', 499.99, 25, 'Gaming', NOW()),
('Xbox Series X', '4K gaming console with Game Pass Ultimate support', 499.99, 30, 'Gaming', NOW()),
('NVIDIA RTX 4090', 'Powerful graphics card for high-end gaming and AI', 1599.99, 10, 'PC Components', NOW()),
('Samsung 49" Odyssey G9', 'Ultra-wide curved gaming monitor with 240Hz refresh rate', 1499.99, 15, 'Monitors', NOW()),
('Apple iPad Pro 12.9', 'M2-powered tablet with Liquid Retina XDR display', 1299.99, 35, 'Tablets', NOW()),
('Google Pixel 8 Pro', 'Google’s AI-powered smartphone with best-in-class camera', 999.99, 40, 'Smartphones', NOW()),
('Razer BlackWidow V4', 'RGB mechanical keyboard for gaming', 179.99, 50, 'Accessories', NOW()),
('AMD Ryzen 9 7950X', '16-core processor for high-performance computing', 699.99, 20, 'PC Components', NOW()),
('Apple Watch Ultra', 'Premium smartwatch with rugged design and long battery life', 799.99, 30, 'Wearables', NOW());


INSERT INTO orders (user_id, total_price, status) VALUES
(8, 150.00, 'completed'),
(8, 200.50, 'pending'),
(9, 99.99, 'completed'),
(9, 75.25, 'cancelled'),
(10, 180.75, 'completed'),
(10, 250.00, 'pending'),
(11, 300.00, 'completed'),
(12, 450.25, 'completed'),
(12, 500.00, 'pending'),
(13, 220.00, 'cancelled'),
(14, 130.50, 'completed'),
(14, 310.75, 'pending'),
(15, 275.00, 'completed'),
(16, 400.00, 'completed'),
(17, 125.00, 'pending');



--SELECT 
--    o.user_id, 
--    SUM(o.total_price) AS total_spent
--FROM orders o
--WHERE o.created_at <= NOW() - INTERVAL '1 month'
--GROUP BY o.user_id
--ORDER BY total_spent DESC
--LIMIT 5;



