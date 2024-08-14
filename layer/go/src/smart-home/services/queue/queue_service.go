package queue

import (
	"context"
	"smart-home/layer/go/src/smart-home/models"
)

type Service interface {
	ProcessMessages(context.Context, *models.Device) error
}
