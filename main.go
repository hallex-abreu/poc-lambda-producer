package main

import (
	"github/hallex-abreu/poc-lambda-producer/src/infra/lambda"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lambda.Execute()
}
