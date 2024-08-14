package deleteDevice

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"smart-home/layer/go/src/smart-home/config"
	"smart-home/layer/go/src/smart-home/services"
)

var deviceService services.DeviceService

func init() {
	config.InitDeviceServiceConfig()
	deviceService = config.AppConfig.DeviceService
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	if err := deviceService.DeleteDevice(context.Background(), id); err != nil {
		config.AppConfig.Log.Printf("Failed to delete device: %v", err)
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusNoContent}, nil
}

func main() {
	lambda.Start(Handler)
}
