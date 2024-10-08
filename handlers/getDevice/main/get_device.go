package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"smart-home/config"
	"smart-home/services"
)

var deviceService services.DeviceService

func init() {
	config.InitDeviceServiceConfig()
	deviceService = config.AppConfig.DeviceService
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	device, err := deviceService.GetDevice(context.Background(), id)
	if err != nil {
		log.Printf("Failed to get device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}
	if device == nil {
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound}, nil
	}

	body, err := json.Marshal(device)
	if err != nil {
		log.Printf("Failed to marshal device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(body)}, nil
}

func main() {
	lambda.Start(Handler)
}
