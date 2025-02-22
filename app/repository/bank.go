package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

type BankRepository interface {
	Deposit(userID int, amount float64) error
	Withdraw(userID int, amount float64) error
	GetDetails(userID int) (float64, error)
}

type bankRepository struct {
	db *sql.DB
}

func NewBankRepository(db *sql.DB) BankRepository {
	return &bankRepository{db: db}
}

// Deposit Money (Thread-Safe with Transactions)
func (r *bankRepository) Deposit(userID int, amount float64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE user_id = $2", amount, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("Deposited $%.2f to user %d\n", amount, userID)
	return nil
}

// Withdraw Money (Thread-Safe with Transactions)
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

// Get Account Details
func (r *bankRepository) GetDetails(userID int) (float64, error) {
	var balance float64
	err := r.db.QueryRow("SELECT balance FROM accounts WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
