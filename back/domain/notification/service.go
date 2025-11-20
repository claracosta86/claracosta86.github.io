package notification

import (
	"context"
)

// Service defines the business logic for user operations
type Service interface {
	// GetNotifications retrieves notifications for a specific user
	GetNotifications(ctx context.Context, userID int) ([]NotificationCulturalList, error)

	// MarkNotificationsAsSeen marks user notifications as seen
	MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error
}

type service struct {
	repository Repository
}

// NewService creates a new notification service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetNotifications(ctx context.Context, userID int) ([]NotificationCulturalList, error) {
	return s.repository.FindByUserID(ctx, userID)
}

func (s *service) MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	return s.repository.MarkAsSeen(ctx, userID, notificationIDs)
}
