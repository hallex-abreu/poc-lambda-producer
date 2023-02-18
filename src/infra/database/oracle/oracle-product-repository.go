package oracle

import (
	"github/hallex-abreu/poc-lambda-producer/src/domain"
	dbconfig "github/hallex-abreu/poc-lambda-producer/src/infra/database/oracle/db-config"
	"time"
)

type OracleProductRepository struct{}

func (opr OracleProductRepository) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	result := dbconfig.Oracle.Where("status_process = ?", "P").Limit(1000).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (opr OracleProductRepository) ProcessIntegrationProduct(product *domain.Product) error {
	status_integrated := "I"
	date_integrated := time.Now()

	product.StatusProcess = &status_integrated
	product.DateProcess = &date_integrated

	result := dbconfig.Oracle.Where("sku = ?", product.Sku).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
