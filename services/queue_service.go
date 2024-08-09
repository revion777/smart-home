package services

import (
	"context"
	"encoding/json"
	"smart-home/models"
	"smart-home/queue"
	"smart-home/repositories"
	"time"
)

type QueueService interface {
	ProcessMessages(ctx context.Context) error
}

type queueService struct {
	queue      queue.Queue
	deviceRepo repositories.DeviceRepository
}

func NewQueueService(q queue.Queue, repo repositories.DeviceRepository) QueueService {
	return &queueService{queue: q, deviceRepo: repo}
}

func (s *queueService) ProcessMessages(ctx context.Context) error {
	messages, err := s.queue.ReceiveMessages(ctx)
	if err != nil {
		return err
	}

	for _, msg := range messages {
		var message models.DeviceMessage
		if err := json.Unmarshal([]byte(*msg.Body), &message); err != nil {
			return err
		}

		// Логика обработки сообщения
		device, err := s.deviceRepo.GetDevice(ctx, message.DeviceID)
		if err != nil {
			return err
		}
		if device == nil {
			continue
		}

		device.HomeID = message.HomeID
		device.ModifiedAt = time.Now().UnixMilli()
		if err := s.deviceRepo.UpdateDevice(ctx, device); err != nil {
			return err
		}
	}
	return nil
}
