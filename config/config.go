package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"smart-home/repositories"
	"smart-home/services"
	"smart-home/services/queue"
)

type Config struct {
	DeviceService services.DeviceService
	QueueService  queue.Service
}

var (
	AppConfig Config
)

func InitDeviceServiceConfig() {
	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(awsCfg)
	deviceRepository := repositories.NewDeviceRepository(dynamoClient)
	AppConfig.DeviceService = services.NewDeviceService(deviceRepository)
}

func InitQueueServiceConfig() {
	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(awsCfg)
	deviceRepository := repositories.NewDeviceRepository(dynamoClient)
	deviceService := services.NewDeviceService(deviceRepository)
	AppConfig.QueueService = queue.NewQueueService(deviceService)
}
