package tests

import (
	"context"
	"errors"
	"smart-home/models"
	"smart-home/services/queue"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeviceService struct {
	mock.Mock
}

func (m *MockDeviceService) CreateDevice(ctx context.Context, device *models.Device) error {
	args := m.Called(ctx, device)
	return args.Error(0)
}

func (m *MockDeviceService) GetDevice(ctx context.Context, id string) (*models.Device, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Device), args.Error(1)
}

func (m *MockDeviceService) UpdateDevice(ctx context.Context, device *models.Device) error {
	args := m.Called(ctx, device)
	return args.Error(0)
}

func (m *MockDeviceService) DeleteDevice(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestServiceImpl_ProcessMessages(t *testing.T) {
	ctx := context.TODO()
	mockService := new(MockDeviceService)

	device := &models.Device{
		ID:         "device123",
		Name:       "Smart Light",
		Type:       "Light",
		HomeID:     "home456",
		MAC:        "00:11:22:33:44:55",
		ModifiedAt: time.Now().Add(-time.Hour).UnixMilli(),
	}

	mockService.On("GetDevice", ctx, device.ID).Return(device, nil)

	mockService.On("UpdateDevice", ctx, mock.AnythingOfType("*models.Device")).Return(nil)

	queueService := queue.NewQueueService(mockService)

	err := queueService.ProcessMessages(ctx, device)

	assert.NoError(t, err)
	mockService.AssertCalled(t, "GetDevice", ctx, device.ID)
	mockService.AssertCalled(t, "UpdateDevice", ctx, mock.AnythingOfType("*models.Device"))

	assert.GreaterOrEqual(t, device.ModifiedAt, time.Now().Add(-time.Minute).UnixMilli())

	mockService.AssertExpectations(t)
}

func TestServiceImpl_ProcessMessages_GetDeviceError(t *testing.T) {
	ctx := context.TODO()
	mockService := new(MockDeviceService)
	device := &models.Device{
		ID: "device123",
	}

	mockService.On("GetDevice", ctx, device.ID).Return(nil, errors.New("database error"))

	queueService := queue.NewQueueService(mockService)

	err := queueService.ProcessMessages(ctx, device)

	assert.EqualError(t, err, "database error")
	mockService.AssertCalled(t, "GetDevice", ctx, device.ID)
	mockService.AssertNotCalled(t, "UpdateDevice", ctx, mock.AnythingOfType("*models.Device"))

	mockService.AssertExpectations(t)
}

func TestServiceImpl_ProcessMessages_UpdateDeviceError(t *testing.T) {
	ctx := context.TODO()
	mockService := new(MockDeviceService)

	device := &models.Device{
		ID:         "device123",
		Name:       "Smart Light",
		Type:       "Light",
		HomeID:     "home456",
		MAC:        "00:11:22:33:44:55",
		ModifiedAt: time.Now().Add(-time.Hour).UnixMilli(),
	}

	mockService.On("GetDevice", ctx, device.ID).Return(device, nil)
	mockService.On("UpdateDevice", ctx, mock.AnythingOfType("*models.Device")).Return(errors.New("update error"))

	queueService := queue.NewQueueService(mockService)

	err := queueService.ProcessMessages(ctx, device)

	assert.EqualError(t, err, "update error")
	mockService.AssertCalled(t, "GetDevice", ctx, device.ID)
	mockService.AssertCalled(t, "UpdateDevice", ctx, mock.AnythingOfType("*models.Device"))

	mockService.AssertExpectations(t)
}
