package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event json.RawMessage) error {
	log.Println("Hello, World!!!!")
	return nil
}

func main() {
	lambda.Start(Handler)
}
