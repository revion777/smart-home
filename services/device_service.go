package services

import (
	"smart-home/models"
	"smart-home/repositories"
)

func CreateDevice(device models.Device) (models.Device, error) {
	return repositories.CreateDevice(device)
}

func GetDevice(id string) (models.Device, error) {
	return repositories.GetDevice(id)
}

func UpdateDevice(id string, device models.Device) (models.Device, error) {
	return repositories.UpdateDevice(id, device)
}

func DeleteDevice(id string) error {
	return repositories.DeleteDevice(id)
}
