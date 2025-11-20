package notification

import (
	"context"
)

// Service defines the business logic for user operations
type Service interface {
	// GetNotifications retrieves notifications for a user
	GetNotifications(ctx context.Context, userID int, favoritesMap map[int]string) ([]NotificationCulturalList, error)

	// MarkNotificationsAsSeen marks specified notifications as seen for a user
	MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error
}

type service struct {
	repository Repository
}

// NewService creates a new user service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetNotifications(ctx context.Context, userID int, favoritesMap map[int]string) ([]NotificationCulturalList, error) {
	return s.repository.FindByUserIDAndFavorites(ctx, userID, favoritesMap)
}

func (s *service) MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	return s.repository.MarkAsSeen(ctx, userID, notificationIDs)
}
