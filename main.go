package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"smart-home/handlers"
)

func main() {
	lambda.Start(handlers.CreateDevice)
	lambda.Start(handlers.GetDevice)
	lambda.Start(handlers.UpdateDevice)
	lambda.Start(handlers.DeleteDevice)
	lambda.Start(handlers.ProcessSQSMessage)
}
