package queue

import (
	"context"
	"smart-home/layer/models"
)

type Service interface {
	ProcessMessages(context.Context, *models.Device) error
}
