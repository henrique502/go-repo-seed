package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/henrique502/go-repo-seed/application/sync"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context) error {
	return sync.New().Integrations()
}
