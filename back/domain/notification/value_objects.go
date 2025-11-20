package notification

import "errors"

type NotificationType string

const (
	NotificationTypeCommented NotificationType = "commented"
	NotificationTypeCanceled  NotificationType = "canceled"
	NotificationTypeUpdated   NotificationType = "updated"
)

func NewNotificationType(value string) (NotificationType, error) {
	validTypes := map[string]bool{
		"commented": true,
		"canceled":  true,
		"updated":   true,
	}
	if !validTypes[value] && value != "" {
		return NotificationType(""), errors.New("invalid notification type")
	}
	return NotificationType(value), nil
}

func (n NotificationType) String() string {
	return string(n)
}
