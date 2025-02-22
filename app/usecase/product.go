package usecase

import (
	"nexmedis-technical-test/app/model/entity"
	"nexmedis-technical-test/app/repository"
)

type ProductUsecase interface {
	GetProducts(name string, limit, offset int) ([]entity.Product, error)
	GetProductByID(id int) (*entity.Product, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo}
}

func (u *productUsecase) GetProducts(name string, limit, offset int) ([]entity.Product, error) {
	return u.productRepo.GetProducts(name, limit, offset)
}

func (u *productUsecase) GetProductByID(id int) (*entity.Product, error) {
	return u.productRepo.GetProductByID(id)
}
