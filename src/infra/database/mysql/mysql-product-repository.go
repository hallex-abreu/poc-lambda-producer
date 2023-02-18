package mysql

import (
	"github/hallex-abreu/poc-lambda-producer/src/domain"
	dbconfig "github/hallex-abreu/poc-lambda-producer/src/infra/database/mysql/db-config"
	"time"
)

type MysqlProductRepository struct{}

func (mpr MysqlProductRepository) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	result := dbconfig.Mysql.Where("status_process = ?", "P").Limit(1000).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (mpr MysqlProductRepository) ProcessIntegrationProduct(product *domain.Product) error {
	status_integrated := "I"
	date_integrated := time.Now()

	product.StatusProcess = &status_integrated
	product.DateProcess = &date_integrated

	result := dbconfig.Mysql.Where("sku = ?", product.Sku).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
