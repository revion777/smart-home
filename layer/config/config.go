package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"os"
	"smart-home/layer/repositories"
	"smart-home/layer/services"
	"smart-home/layer/services/queue"
)

type Config struct {
	DeviceService services.DeviceService
	QueueService  queue.Service
	Log           *log.Logger
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
	AppConfig.Log = log.New(os.Stdout, "", log.LstdFlags)
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
	AppConfig.Log = log.New(os.Stdout, "", log.LstdFlags)
}
