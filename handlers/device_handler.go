package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"smart-home/models"
	"smart-home/services"
	"time"
)

func CreateDevice(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var device models.Device
	if err := json.Unmarshal([]byte(request.Body), &device); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	device.CreatedAt = time.Now().UnixMilli()
	device.ModifiedAt = device.CreatedAt

	createdDevice, err := services.CreateDevice(device)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	body, _ := json.Marshal(createdDevice)
	return events.APIGatewayProxyResponse{StatusCode: http.StatusCreated, Body: string(body)}, nil
}

func GetDevice(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]

	device, err := services.GetDevice(id)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}
	if device.ID == "" {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound}, nil
	}

	body, _ := json.Marshal(device)
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(body)}, nil
}

func UpdateDevice(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]

	var device models.Device
	if err := json.Unmarshal([]byte(request.Body), &device); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}
	device.ModifiedAt = time.Now().UnixMilli()

	updatedDevice, err := services.UpdateDevice(id, device)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	body, _ := json.Marshal(updatedDevice)
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(body)}, nil
}

func DeleteDevice(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]

	if err := services.DeleteDevice(id); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusNoContent}, nil
}
