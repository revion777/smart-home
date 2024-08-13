package config

import (
	"errors"
	"log"
	"os"
	"smart-home/repositories"
	"smart-home/services"
	"smart-home/services/queue"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Config struct {
	DeviceService services.DeviceService
	QueueService  queue.Service
	Log           *log.Logger
}

var (
	AppConfig Config
)

func InitConfig() {
	awsCfg, err := loadAWSConfig()
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(awsCfg)
	deviceRepository := repositories.NewDeviceRepository(dynamoClient)
	AppConfig.DeviceService = services.NewDeviceService(deviceRepository)
	AppConfig.QueueService = queue.NewQueueService(AppConfig.DeviceService)
	AppConfig.Log = log.New(os.Stdout, "", log.LstdFlags)
}

func loadAWSConfig() (aws.Config, error) {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	if accessKeyID == "" || secretAccessKey == "" || region == "" {
		return aws.Config{}, errors.New("AWS credentials or region not set in environment variables")
	}

	return aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""),
		Region:      region,
	}, nil
}
