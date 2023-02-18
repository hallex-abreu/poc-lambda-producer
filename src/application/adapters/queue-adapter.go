package adapters

import "github/hallex-abreu/poc-lambda-producer/src/domain"

type QueueAdapter interface {
	SendingInSqs(product *domain.Product) error
}
