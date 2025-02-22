package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"nexmedis-technical-test/app/model/entity"
)

type CartRepository interface {
	AddToCart(userID, productID, quantity int) error
	GetCartItems(userID int) ([]entity.CartItem, error)
	Checkout(userID int) (int, error)
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) AddToCart(userID, productID, quantity int) error {
	// Check if item already exists in cart
	var existingQuantity int

	err := r.db.QueryRow("SELECT quantity FROM cart_items WHERE user_id = $1 AND product_id = $2", userID, productID).Scan(&existingQuantity)

	if err == nil {
		// Update quantity if product is already in the cart
		_, err = r.db.Exec("UPDATE cart_items SET quantity = $1 WHERE user_id = $2 AND product_id = $3", existingQuantity+quantity, userID, productID)
		fmt.Println("err2", err)
		return err
	} else if err != sql.ErrNoRows {
		return err
	}

	// Insert new cart item if not found
	_, err = r.db.Exec("INSERT INTO cart_items (user_id, product_id, quantity) VALUES ($1, $2, $3)", userID, productID, quantity)
	return err
}

func (r *cartRepository) GetCartItems(userID int) ([]entity.CartItem, error) {
	rows, err := r.db.Query("SELECT id, user_id, product_id, quantity, created_at, updated_at FROM cart_items WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []entity.CartItem
	for rows.Next() {
		var cart entity.CartItem
		if err := rows.Scan(&cart.ID, &cart.UserID, &cart.ProductID, &cart.Quantity, &cart.CreatedAt, &cart.UpdatedAt); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cart)
	}

	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	fmt.Println("cartItems", cartItems)
	return cartItems, nil
}

func (r *cartRepository) Checkout(userID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Step 1: Fetch cart items
	rows, err := tx.Query(`
		SELECT c.product_id, c.quantity, p.price 
		FROM cart_items c
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = $1`, userID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var cartItems []entity.OrderDetail
	var totalPrice float64

	for rows.Next() {
		var item entity.OrderDetail
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price); err != nil {
			return 0, err
		}
		item.Subtotal = float64(item.Quantity) * item.Price
		totalPrice += item.Subtotal
		cartItems = append(cartItems, item)
	}

	if len(cartItems) == 0 {
		return 0, err
	}

	// Step 2: Insert order into orders table
	var orderID int
	err = tx.QueryRow("INSERT INTO orders (user_id, total_price) VALUES ($1, $2) RETURNING id", userID, totalPrice).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	// Step 3: Insert into order_details table
	stmt, err := tx.Prepare("INSERT INTO order_details (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	for _, item := range cartItems {
		_, err := stmt.Exec(orderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return 0, err
		}
	}

	// Step 4: Clear cart after checkout
	_, err = tx.Exec("DELETE FROM cart_items WHERE user_id = $1", userID)
	if err != nil {
		return 0, err
	}

	// Step 5: Commit transaction
	tx.Commit()
	return 0, nil
}
