package queue

import (
	"context"
	"smart-home/models"
)

type Service interface {
	ProcessMessages(context.Context, *models.Device) error
}
