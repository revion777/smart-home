package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"smart-home/config"
	"smart-home/queue"
	"smart-home/repositories"
	"smart-home/services"
)

var (
	queueService services.QueueService
)

func init() {
	deviceRepo := repositories.NewDynamoDeviceRepository(config.AppConfig.DynamoClient)
	queueService = services.NewQueueService(queue.NewSQSQueue(config.AppConfig.SQSClient,
		config.AppConfig.SQSQueueURL),
		deviceRepo)
}

func ProcessSQSMessageHandler(ctx context.Context, event events.SQSEvent) {
	err := queueService.ProcessMessages(ctx)
	if err != nil {
		config.AppConfig.Log.Printf("Failed to process SQS messages: %v", err)
	}
}
