package usecases

import (
	"errors"
	"github/hallex-abreu/poc-lambda-producer/src/application/adapters"
	"github/hallex-abreu/poc-lambda-producer/src/application/repositories"
	"github/hallex-abreu/poc-lambda-producer/src/domain"
)

func ProcessProductUseCase(productRepository repositories.ProductRepository, queueAdapter adapters.QueueAdapter, logAdapter adapters.LogAdapter) error {
	logAdapter.Log("info", "Iniciando busca de produtos")
	products, err := productRepository.GetProducts()

	if err != nil {
		logAdapter.Log("error", err.Error())
		return errors.New(err.Error())
	}

	logAdapter.Log("info", "Finalizado busca de produtos")

	for _, product := range products {
		product_format := &domain.Product{
			Sku:         product.Sku,
			Description: product.Description,
			Price:       product.Price,
		}

		logAdapter.Log("info", "Iniciando envio de produto para fila")
		queueAdapter.SendingInSqs(product_format)
		logAdapter.Log("info", "finalizado envio de produto para fila")
	}

	logAdapter.Log("info", "finalizado processamento com sucesso")
	return nil
}
