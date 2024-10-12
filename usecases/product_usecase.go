package usecases

import (
	"github.com/DanielDxD/goprodu/model"
	"github.com/DanielDxD/goprodu/repositories"
)

type ProductUseCase struct {
	repository repositories.ProductRepository
}

func NewProductUseCase(repository repositories.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}
func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}
func (pu *ProductUseCase) GetProduct(id int) (*model.Product, error) {
	return pu.repository.GetProduct(id)
}
