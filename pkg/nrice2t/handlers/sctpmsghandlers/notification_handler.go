package sctpmsghandlers

import (
	"nRIC/pkg/nrice2t/models"
)

type NotificationHandler interface {
	Handle(*models.NotificationRequest)
}
