package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() string {
	foo := os.Getenv("FOO")
	return foo
}

func main() {
	lambda.Start(HandleRequest)
}
