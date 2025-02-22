package entity

import "time"

type Order struct {
	ID           int           `json:"id"`
	UserID       int           `json:"user_id"`
	TotalPrice   float64       `json:"total_price"`
	Status       string        `json:"status"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	OrderDetails []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID"`
}
