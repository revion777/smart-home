package services

import (
	"context"
	"smart-home/models"
	"smart-home/repositories"
)

type DeviceServiceImpl struct {
	repo repositories.DeviceRepository
}

func NewDeviceService(repo repositories.DeviceRepository) DeviceService {
	return &DeviceServiceImpl{repo: repo}
}

func (service *DeviceServiceImpl) CreateDevice(ctx context.Context, device *models.Device) error {
	return service.repo.CreateDevice(ctx, device)
}

func (service *DeviceServiceImpl) GetDevice(ctx context.Context, id string) (*models.Device, error) {
	return service.repo.GetDevice(ctx, id)
}

func (service *DeviceServiceImpl) UpdateDevice(ctx context.Context, device *models.Device) error {
	return service.repo.UpdateDevice(ctx, device)
}

func (service *DeviceServiceImpl) DeleteDevice(ctx context.Context, id string) error {
	return service.repo.DeleteDevice(ctx, id)
}
