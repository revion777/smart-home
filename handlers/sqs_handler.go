package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"smart-home/models"
	"smart-home/services"
	"time"
)

func ProcessSQSMessage(sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		err := handleSQSMessage(message.Body)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleSQSMessage(messageBody string) error {
	var association models.DeviceHomeAssociation
	err := json.Unmarshal([]byte(messageBody), &association)
	if err != nil {
		return err
	}

	device, err := services.GetDevice(association.DeviceID)
	if err != nil {
		return err
	}
	if device.ID == "" {
		return nil // Device not found, skip update
	}

	device.HomeID = association.HomeID
	device.ModifiedAt = time.Now().UnixMilli()

	_, err = services.UpdateDevice(association.DeviceID, device)
	return err
}
