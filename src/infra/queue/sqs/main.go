package sqs

import (
	"encoding/json"
	"fmt"
	"github/hallex-abreu/poc-lambda-producer/src/domain"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Sqs struct{}

func (s Sqs) SendingInSqs(product *domain.Product) error {
	// Credentials configuration
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

	// AWS SDK Session Configuration
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: creds,
	})

	if err != nil {
		fmt.Println("Erro ao criar sessão:", err)
		return err
	}

	// Creating the SQS Client
	sqsClient := sqs.New(sess)

	// message to be sent
	messageBody, err := json.Marshal(product)

	// Creating the message entry
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    aws.String(os.Getenv("AWS_QUEUE_PRODUCT_URL")),
	}

	// sending the message
	result, err := sqsClient.SendMessage(sendMessageInput)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
		return err
	}

	fmt.Println("Mensagem enviada com sucesso, ID da mensagem:", *result.MessageId)
	return nil
}
