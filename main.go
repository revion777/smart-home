package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"smart-home/config"
	"smart-home/handlers"
)

func main() {
	config.InitConfig()

	lambda.Start(handlers.CreateDeviceHandler)
	lambda.Start(handlers.GetDeviceHandler)
	lambda.Start(handlers.UpdateDeviceHandler)
	lambda.Start(handlers.DeleteDeviceHandler)
	lambda.Start(handlers.ProcessSQSMessageHandler)
}
