package dto

type TransactionRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
