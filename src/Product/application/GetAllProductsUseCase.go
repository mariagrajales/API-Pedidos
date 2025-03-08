package application

import (
	"api-order/src/Product/domain/entities"
	"api-order/src/Product/domain/ports"
)

type GetAllProductsUseCase struct {
	ProductRepository ports.IProduct
}

func NewGetAllProductsUseCase(productRepository ports.IProduct) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{ProductRepository: productRepository}
}

func (p *GetAllProductsUseCase) Run() ([]entities.Product, error) {
	products, err := p.ProductRepository.GetAll()

	if err != nil {
		return []entities.Product{}, err
	}

	return products, nil
}