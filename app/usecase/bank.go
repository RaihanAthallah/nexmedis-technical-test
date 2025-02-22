package usecase

import (
	"errors"
	"nexmedis-technical-test/app/repository"
)

type BankUsecase interface {
	// Deposit(userID int, amount float64)
	// Withdraw(userID int, amount float64)
	Deposit(userID int, amount float64) error
	Withdraw(userID int, amount float64) error
	GetDetails(userID int) (float64, error)
}

type bankUsecase struct {
	bankRepo repository.BankRepository
}

func NewBankUsecase(bankRepo repository.BankRepository) BankUsecase {
	return &bankUsecase{bankRepo: bankRepo}
}

// Deposit: Handle concurrency here, not in the repository
// func (u *bankUsecase) Deposit(userID int, amount float64) {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()
// 		err := u.bankRepo.Deposit(userID, amount)
// 		if err != nil {
// 			log.Println("Deposit failed:", err)
// 		}
// 	}()

// 	wg.Wait() // Wait for goroutine to finish
// }

// Withdraw: Handle concurrency at the usecase level
// func (u *bankUsecase) Withdraw(userID int, amount float64) {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()
// 		err := u.bankRepo.Withdraw(userID, amount)
// 		if err != nil {
// 			log.Println("Withdrawal failed:", err)
// 		}
// 	}()

// 	wg.Wait()
// }

// Deposit without WaitGroup (simpler and more efficient)
func (u *bankUsecase) Deposit(userID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	err := u.bankRepo.Deposit(userID, amount)
	if err != nil {
		return err
	}
	return nil
}

// Withdraw without WaitGroup (simpler and more efficient)
func (u *bankUsecase) Withdraw(userID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	err := u.bankRepo.Withdraw(userID, amount)
	if err != nil {
		return err
	}
	return nil
}

func (u *bankUsecase) GetDetails(userID int) (float64, error) {
	return u.bankRepo.GetDetails(userID)
}
