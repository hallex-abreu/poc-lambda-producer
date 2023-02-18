package inmemoryrepository

import "github/hallex-abreu/poc-lambda-producer/src/domain"

type InMemoryProductRepository struct{}

func (p InMemoryProductRepository) GetProducts() ([]*domain.Product, error) {
	products := []*domain.Product{} //slice define with zero size

	for i := 0; i < 4; i++ {
		new_product := &domain.Product{
			Sku:         "tetetete",
			Description: "teste",
			Price:       1234,
		}

		products = append(products, new_product)
	}

	return products, nil
}

func (p InMemoryProductRepository) ProcessProduct(product *domain.Product) error {
	return nil
}
