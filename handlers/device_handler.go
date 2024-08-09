package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"smart-home/config"
	"smart-home/models"
	"smart-home/repositories"
	"smart-home/services"
)

var deviceService services.DeviceService

func init() {
	deviceService = services.DeviceService(repositories.NewDynamoDeviceRepository(config.AppConfig.DynamoClient))
}

func CreateDeviceHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
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

func GetDeviceHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	device, err := deviceService.GetDevice(context.Background(), id)
	if err != nil {
		config.AppConfig.Log.Printf("Failed to get device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}
	if device == nil {
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound}, nil
	}

	body, err := json.Marshal(device)
	if err != nil {
		config.AppConfig.Log.Printf("Failed to marshal device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(body)}, nil
}

func UpdateDeviceHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
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

func DeleteDeviceHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	if err := deviceService.DeleteDevice(context.Background(), id); err != nil {
		config.AppConfig.Log.Printf("Failed to delete device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusNoContent}, nil
}
