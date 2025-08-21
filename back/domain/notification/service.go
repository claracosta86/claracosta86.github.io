package notification

import (
	"context"
)

// Service defines the business logic for user operations
type Service interface {
	// GetNotifications retrieves notifications for a user
	GetNotifications(ctx context.Context, userID int) ([]NotificationCulturalList, error)
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

func (s *service) GetNotifications(ctx context.Context, userID int) ([]NotificationCulturalList, error) {
	return  s.repository.FindByUserID(ctx, userID)
}
