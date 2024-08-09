package config

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Config struct {
	DynamoClient *dynamodb.Client
	SQSClient    *sqs.Client
	SQSQueueURL  string
	Log          *log.Logger
}

var (
	AppConfig Config
)

func InitConfig() {
	awsCfg, err := loadAWSConfig()
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	AppConfig.DynamoClient = dynamodb.NewFromConfig(awsCfg)
	AppConfig.SQSClient = sqs.NewFromConfig(awsCfg)
	AppConfig.SQSQueueURL = os.Getenv("SQS_QUEUE_URL")
	if AppConfig.SQSQueueURL == "" {
		log.Fatal("SQS_QUEUE_URL is not set")
	}

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
