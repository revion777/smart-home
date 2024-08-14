package processSqs

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"smart-home/layer/go/src/smart-home/config"
	"smart-home/layer/go/src/smart-home/models"
	"smart-home/layer/go/src/smart-home/services/queue"
)

var (
	queueService queue.Service
)

func init() {
	config.InitQueueServiceConfig()
	queueService = config.AppConfig.QueueService
}

func Handler(sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		var device models.Device

		if err := json.Unmarshal([]byte(message.Body), &device); err != nil {
			config.AppConfig.Log.Printf(
				"Failed to unmarshal SQS message: %v, messageId: %v, body: %v",
				err, message.MessageId, message.Body)
			continue
		}

		if err := queueService.ProcessMessages(context.Background(), &device); err != nil {
			config.AppConfig.Log.Printf(
				"Failed to process SQS message: %v, messageId: %v", err, message.MessageId)
			return err
		}

		config.AppConfig.Log.Printf("Successfully processed SQS message: %s", message.MessageId)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
