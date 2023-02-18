package lambda

import (
	"fmt"
	usecases "github/hallex-abreu/poc-lambda-producer/src/application/use-cases"
	"github/hallex-abreu/poc-lambda-producer/src/infra/database/mysql"
	dbconfig "github/hallex-abreu/poc-lambda-producer/src/infra/database/mysql/db-config"
	"github/hallex-abreu/poc-lambda-producer/src/infra/logs"
	"github/hallex-abreu/poc-lambda-producer/src/infra/queue/sqs"
	"log"
)

func Execute() {
	productRepository := mysql.MysqlProductRepository{}
	queueAdapter := sqs.Sqs{}
	logAdapter := logs.LocalLog{}

	dbconfig.Connection()

	err := usecases.ProcessProductUseCase(productRepository, queueAdapter, logAdapter)

	if err != nil {
		log.Fatal(err.Error())
		fmt.Print("lambda terminated with error")
	} else {
		fmt.Print("lambda ended successfully")
	}
}
