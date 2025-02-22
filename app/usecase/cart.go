package usecase

import (
	"errors"
	"nexmedis-technical-test/app/model/entity"
	"nexmedis-technical-test/app/repository"
)

type CartUsecase interface {
	AddToCart(userID, productID, quantity int) error
	GetCartItems(userID int) ([]entity.CartItem, error)
	Checkout(userID int) (int, error)
}

type cartUsecase struct {
	cartRepo repository.CartRepository
}

func NewCartUsecase(cartRepo repository.CartRepository) CartUsecase {
	return &cartUsecase{cartRepo: cartRepo}
}

func (u *cartUsecase) AddToCart(userID, productID, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return u.cartRepo.AddToCart(userID, productID, quantity)
}

func (u *cartUsecase) GetCartItems(userID int) ([]entity.CartItem, error) {
	return u.cartRepo.GetCartItems(userID)
}

func (u *cartUsecase) Checkout(userID int) (int, error) {
	return u.cartRepo.Checkout(userID)
}
