package services

import (
	"context"
	"smart-home/layer/go/src/smart-home/models"
)

type DeviceService interface {
	CreateDevice(ctx context.Context, device *models.Device) error
	GetDevice(ctx context.Context, id string) (*models.Device, error)
	UpdateDevice(ctx context.Context, device *models.Device) error
	DeleteDevice(ctx context.Context, id string) error
}
