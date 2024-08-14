package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"os"
	"smart-home/layer/go/src/smart-home/repositories"
	services2 "smart-home/layer/go/src/smart-home/services"
	queue2 "smart-home/layer/go/src/smart-home/services/queue"
)

type Config struct {
	DeviceService services2.DeviceService
	QueueService  queue2.Service
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
	AppConfig.DeviceService = services2.NewDeviceService(deviceRepository)
	AppConfig.Log = log.New(os.Stdout, "", log.LstdFlags)
}

func InitQueueServiceConfig() {
	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(awsCfg)
	deviceRepository := repositories.NewDeviceRepository(dynamoClient)
	deviceService := services2.NewDeviceService(deviceRepository)
	AppConfig.QueueService = queue2.NewQueueService(deviceService)
	AppConfig.Log = log.New(os.Stdout, "", log.LstdFlags)
}
