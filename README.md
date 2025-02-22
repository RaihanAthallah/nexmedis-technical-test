# 1. E-Commerce API

## Overview

The E-Commerce API provides endpoints for user authentication, product browsing, cart management, order checkout, and banking transactions such as deposits and withdrawals.

## Base URL

```
http://localhost:8090/api/v1
```

## Authentication

### Login

**Endpoint:**

```
POST /user/login
```

**Request Body:**

```json
{
  "username": "raihan",
  "password": "raihan"
}
```

**Response:**

```json
{
  "status": 200,
  "message": "Login success",
  "data": {
    "RefreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InJhaWhhbiIsImVtYWlsIjoicmFpaGFuQG1haWwuY29tIiwiaWQiOiJTS0FENTF6NUgyOG1FVEF5cUxVdzE1dS9xQ0V2ZERYNGpodnYwSWs9IiwiZXhwIjoxNzQwODIyMzQzLCJpYXQiOjE3NDAyMTc1NDN9.abcfRGnxtGXJktyZaZltAcPPbzVKWOT33470twDNz4A",
    "AccessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InJhaWhhbiIsImVtYWlsIjoicmFpaGFuQG1haWwuY29tIiwiaWQiOiJTS0FENTF6NUgyOG1FVEF5cUxVdzE1dS9xQ0V2ZERYNGpodnYwSWs9IiwiZXhwIjoxNzQwMjE4NDQzLCJpYXQiOjE3NDAyMTc1NDN9.JsotyR9F5IxG2Cw-T71l4WQhCq_QsllIlLtuFqO1Qi8"
  }
}
```

## Products

### View Products

**Endpoint:**

```
GET /products?name=Apple&limit=5&page=1
```

**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
  "status": 200,
  "message": "Products retrieved successfully",
  "data": [
    { "id": 1, "name": "Apple iPhone 15 Pro", "description": "Latest iPhone with A17 Pro chip and titanium body", "price": 1199.99, "stock": 50, "category": "Smartphones", "created_at": "2025-02-21T21:44:56.61461Z" },
    { "id": 11, "name": "Apple iPad Pro 12.9", "description": "M2-powered tablet with Liquid Retina XDR display", "price": 1299.99, "stock": 35, "category": "Tablets", "created_at": "2025-02-21T21:44:56.61461Z" },
    { "id": 15, "name": "Apple Watch Ultra", "description": "Premium smartwatch with rugged design and long battery life", "price": 799.99, "stock": 30, "category": "Wearables", "created_at": "2025-02-21T21:44:56.61461Z" }
  ]
}
```

## Shopping Cart

### Add Item to Cart

**Endpoint:**

```
POST /carts/add
```

**Headers:**

```
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "product_id": 2,
  "quantity": 2
}
```

**Response:**

```json
{
  "status": 200,
  "message": "Success add to cart",
  "data": ""
}
```

### View Cart

**Endpoint:**

```
GET /carts
```

**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
    "status": 200,
    "message": "success get cart",
    "data": [
        {
            "id": 4,
            "user_id": 1,
            "product_id": 2,
            "quantity": 2,
            "created_at": "2025-02-22T09:52:35.453344Z",
            "updated_at": "2025-02-22T09:52:35.453344Z"
        }
    ]
}
```

### Checkout

**Endpoint:**

```
POST /carts/checkout
```

**Headers:**

```
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "product_id": 2,
  "quantity": 2
}
```

**Response:**

```json
{
  "status": 200,
  "message": "Checkout successful",
  "data": 0
}
```

## Bank Transactions

### Deposit Funds

**Endpoint:**

```
POST /banks/deposit
```

**Headers:**

```
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "amount": 60000
}
```

**Response:**

```json
{
    "status": 200,
    "message": "Deposit successful",
    "data": ""
}
```

### Withdraw Funds

**Endpoint:**

```
POST /banks/withdraw
```

**Headers:**

```
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "amount": 60000
}
```

**Response:**

```json
{
    "status": 200,
    "message": "Withdraw successful",
    "data": ""
}
```

## Authorization

All endpoints (except login) require an Authorization header:

```
Authorization: Bearer <token>
```

## Notes

- Ensure the JWT token is valid before making requests.
- API requests should be formatted correctly to avoid errors.
- Product stock availability is checked before adding to the cart.

---

# 2. Indexing Strategy
Index unique disarankan dilakukan untuk field yang berpotensi bernilai unique, maka dari itu jika diasumsikan username dan email adalah unique, maka kedua field ini bisa dijadikan index untuk meningkatkan performa . Individual index lebih tepat daripada composite index karena query yang digunakan hanya menggunakan salah satu dari username atau email. Kemudian penggunaan type index yaitu 'hash' dikarenakan untuk pencarian hanya perlu exact match queries (=).

```
CREATE INDEX idx_users_username ON users USING hash (username);
CREATE INDEX idx_users_email ON users USING hash (email);
```
Untuk field created_at yang berisikan data bertipe timestamp maka penggunaan btree bisa dijadikan pertimbangan karena mendukung semua operasi untuk timestamp (<=, =, >=, <, >)

```
CREATE INDEX idx_users_created_at ON users USING btree (created_at);
```

# 3. Preventing Race Condition
```Go
func (r *bankRepository) Withdraw(userID int, amount float64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Check balance
	var balance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return err
	}

	if balance < amount {
		tx.Rollback()
		return errors.New("insufficient balance")
	}

	// Update balance
	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE user_id = $2", amount, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("Withdrew $%.2f from user %d\n", amount, userID)
	return nil
}
```
Beberapa hal yang bisa mencegah race condition :
- Menggunakan database transactions (tx.Begin(), tx.Commit(), tx.Rollback())
- Pengecekan balance sebelum withdrawal

# 4. Optimizing Get Top 5 Customers
```
SELECT 
    customer_id,
    COUNT(*) as total_orders,
    SUM(amount) as total_spent
FROM orders
WHERE order_date >= NOW() - INTERVAL '1 month'
GROUP BY customer_id
ORDER BY total_spent DESC
LIMIT 5;
```
Untuk optimasi pada production bisa dilakukan dengan melakukan indexing composite untuk field order_date, customer_id, dan amout. Selain itu penggunaan service untuk caching seperti redis juga bisa menjadi pertimbangan.


# 5. Refactoring Complex Monolith Program
Menurut saya ada beberapa proses yang harus dilakukan untuk memecah program yang kompleks menjadi service yang lebih kecil.
- Baca kode terlebih dahulu dan lakukan identifikasi fungsionalitas atau fitur apa yang bisa dipisahkan dari kode monolith. 
- Kemudian dari beberapa fungsionalitas tersebut buat ulang terpisah tanpa perlu mengubah kode sumber monolith atau legacy code. 
- Kemudian arahkan proses pada kode program monolith ke service baru dengan menggunakan API
- Lakukan testing dan pastikan aplikasi monolith masih berfungsi seperti seharusnya
- Apabila testing telah berhasil dilakukan maka dilanjutkan untuk fitur lain dan dilakukan secara bertahap hingga keselurah kode monolith sudah diarahkan ke services yang sudah dipecah
- Penggunaan versioning pada API, agar fitur lama tidak terganggu selama proses migrasi.
