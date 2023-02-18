package repositories

import "github/hallex-abreu/poc-lambda-producer/src/domain"

type ProductRepository interface {
	GetProducts() ([]*domain.Product, error)
	ProcessIntegrationProduct(product *domain.Product) error
}
