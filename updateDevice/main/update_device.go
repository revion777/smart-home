package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"smart-home/layer/config"
	"smart-home/layer/models"
	"smart-home/layer/services"
)

var deviceService services.DeviceService

func init() {
	config.InitDeviceServiceConfig()
	deviceService = config.AppConfig.DeviceService
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	var device models.Device
	if err := json.Unmarshal([]byte(request.Body), &device); err != nil {
		config.AppConfig.Log.Printf("Failed to unmarshal request body: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	device.ID = id
	if err := deviceService.UpdateDevice(context.Background(), &device); err != nil {
		config.AppConfig.Log.Printf("Failed to update device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
