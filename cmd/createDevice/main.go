package createDevice

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"smart-home/layer/go/src/smart-home/config"
	"smart-home/layer/go/src/smart-home/models"
	"smart-home/layer/go/src/smart-home/services"
)

var deviceService services.DeviceService

func init() {
	config.InitDeviceServiceConfig()
	deviceService = config.AppConfig.DeviceService
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var device models.Device
	if err := json.Unmarshal([]byte(request.Body), &device); err != nil {
		config.AppConfig.Log.Printf("Failed to unmarshal request body: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	if err := deviceService.CreateDevice(context.Background(), &device); err != nil {
		config.AppConfig.Log.Printf("Failed to create device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusCreated}, nil
}

func main() {
	lambda.Start(Handler)
}
