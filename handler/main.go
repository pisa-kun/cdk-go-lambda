package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (*string, error) {
	message := "Hello World!"
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
