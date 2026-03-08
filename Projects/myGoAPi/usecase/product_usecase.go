package usecase

import (
	"myGoApi/model"
	"myGoApi/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepositoy
}

func NewProductUseCase(repo repository.ProductRepositoy) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
