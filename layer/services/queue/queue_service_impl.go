package queue

import (
	"context"
	"smart-home/layer/models"
	"smart-home/layer/services"
	"time"
)

type ServiceImpl struct {
	deviceService services.DeviceService
}

func NewQueueService(service services.DeviceService) Service {
	return &ServiceImpl{deviceService: service}
}

func (service *ServiceImpl) ProcessMessages(ctx context.Context, device *models.Device) error {
	device, err := service.deviceService.GetDevice(ctx, device.ID)
	if err != nil {
		return err
	}

	device.ModifiedAt = time.Now().UnixMilli()
	if err := service.deviceService.UpdateDevice(ctx, device); err != nil {
		return err
	}
	return nil
}
