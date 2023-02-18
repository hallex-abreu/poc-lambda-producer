package lambda

import (
	"fmt"
	usecases "github/hallex-abreu/poc-lambda-producer/src/application/use-cases"
	inmemoryrepository "github/hallex-abreu/poc-lambda-producer/src/infra/database/in-memory-repository"
	"github/hallex-abreu/poc-lambda-producer/src/infra/logs"
	inmemoryqueue "github/hallex-abreu/poc-lambda-producer/src/infra/queue/in-memory-queue"
	"log"
)

func Execute() {
	productRepository := inmemoryrepository.InMemoryProductRepository{}
	queueAdapter := inmemoryqueue.InMemoryQueue{}
	logAdapter := logs.LocalLog{}

	err := usecases.ProcessProductUseCase(productRepository, queueAdapter, logAdapter)

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Print("Success process")
	}
}
