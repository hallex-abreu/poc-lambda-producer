package usecases

import (
	"encoding/json"
	"errors"
	"github/hallex-abreu/poc-lambda-producer/src/application/adapters"
	"github/hallex-abreu/poc-lambda-producer/src/application/repositories"
	"github/hallex-abreu/poc-lambda-producer/src/domain"
	"strconv"
)

func ProcessProductUseCase(productRepository repositories.ProductRepository, queueAdapter adapters.QueueAdapter, logAdapter adapters.LogAdapter) error {
	logAdapter.Log("info", "Iniciando busca de produtos")
	products, err := productRepository.GetProducts()

	if err != nil {
		logAdapter.Log("error", err.Error())
		return errors.New(err.Error())
	}

	if len(products) == 0 {
		logAdapter.Log("info", "Sem novos produtos para integrar")
		return nil
	}

	logAdapter.Log("info", "Finalizado busca de produtos com total de: "+strconv.Itoa(len(products)))

	for _, product := range products {
		product_format := &domain.Product{
			Sku:         product.Sku,
			Description: product.Description,
			Price:       product.Price,
		}

		logAdapter.Log("info", "Iniciando envio de produto para fila")
		queueAdapter.SendingInSqs(product_format)
		product_log, _ := json.Marshal(product)
		logAdapter.Log("info", "Finalizado envio de produto para fila: "+string(product_log))
		logAdapter.Log("info", "Iniciando processo de integração")
		err := productRepository.ProcessIntegrationProduct(product)

		if err != nil {
			logAdapter.Log("error", err.Error())
			return errors.New(err.Error())
		}

		logAdapter.Log("info", "Finalizando processo de integração")
	}

	return nil
}
