package inmemory

import "github/hallex-abreu/poc-lambda-producer/src/domain"

type InMemoryProductRepository struct{}

func (ipr InMemoryProductRepository) GetProducts() ([]*domain.Product, error) {
	products := []*domain.Product{} //slice define with zero size

	for i := 0; i < 4; i++ {
		new_product := &domain.Product{
			Sku:         "hst717",
			Description: "Macbook",
			Price:       7200.34,
		}

		products = append(products, new_product)
	}

	return products, nil
}

func (ipr InMemoryProductRepository) ProcessProduct(product *domain.Product) error {
	return nil
}
