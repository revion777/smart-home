package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"smart-home/config"
	"smart-home/models"
	"smart-home/services"
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
		log.Printf("Failed to unmarshal request body: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	device.ID = id
	if err := deviceService.UpdateDevice(context.Background(), &device); err != nil {
		log.Printf("Failed to update device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}
