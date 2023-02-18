package inmemoryqueue

import "github/hallex-abreu/poc-lambda-producer/src/domain"

type InMemoryQueue struct{}

func (q InMemoryQueue) SendingInSqs(product *domain.Product) error {
	return nil
}
