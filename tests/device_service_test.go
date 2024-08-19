package tests

import (
	"context"
	"errors"
	"smart-home/models"
	"smart-home/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeviceRepository struct {
	mock.Mock
}

func (m *MockDeviceRepository) CreateDevice(ctx context.Context, device *models.Device) error {
	args := m.Called(ctx, device)
	return args.Error(0)
}

func (m *MockDeviceRepository) GetDevice(ctx context.Context, id string) (*models.Device, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Device), args.Error(1)
}

func (m *MockDeviceRepository) UpdateDevice(ctx context.Context, device *models.Device) error {
	args := m.Called(ctx, device)
	return args.Error(0)
}

func (m *MockDeviceRepository) DeleteDevice(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestDeviceServiceImpl_CreateDevice(t *testing.T) {
	mockRepo := new(MockDeviceRepository)
	service := services.NewDeviceService(mockRepo)

	device := &models.Device{ID: "123", Name: "Test Device"}

	mockRepo.On("CreateDevice", mock.Anything, device).Return(nil)

	err := service.CreateDevice(context.Background(), device)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeviceServiceImpl_GetDevice(t *testing.T) {
	mockRepo := new(MockDeviceRepository)
	service := services.NewDeviceService(mockRepo)

	device := &models.Device{ID: "123", Name: "Test Device"}

	mockRepo.On("GetDevice", mock.Anything, "123").Return(device, nil)

	result, err := service.GetDevice(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, device, result)

	mockRepo.AssertExpectations(t)
}

func TestDeviceServiceImpl_GetDevice_NotFound(t *testing.T) {
	mockRepo := new(MockDeviceRepository)
	service := services.NewDeviceService(mockRepo)

	mockRepo.On("GetDevice", mock.Anything, "123").Return(nil, errors.New("not found"))

	result, err := service.GetDevice(context.Background(), "123")
	assert.Error(t, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}

func TestDeviceServiceImpl_UpdateDevice(t *testing.T) {
	mockRepo := new(MockDeviceRepository)
	service := services.NewDeviceService(mockRepo)

	device := &models.Device{ID: "123", Name: "Updated Device"}

	mockRepo.On("UpdateDevice", mock.Anything, device).Return(nil)

	err := service.UpdateDevice(context.Background(), device)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeviceServiceImpl_DeleteDevice(t *testing.T) {
	mockRepo := new(MockDeviceRepository)
	service := services.NewDeviceService(mockRepo)

	mockRepo.On("DeleteDevice", mock.Anything, "123").Return(nil)

	err := service.DeleteDevice(context.Background(), "123")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
