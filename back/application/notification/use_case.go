package notification

import (
	"context"

	"poc2/back/domain/notification"
	"poc2/back/interface/model"
)

// UseCase defines the application use cases for user operations
type UseCase interface {
	// RegisterUser handles user registration
	GetNotifications(ctx context.Context, userID int) (*model.GetNotificationsResponse, error)
}

type useCase struct {
	notificationService notification.Service
}

// NewUseCase creates a new user use case
func NewUseCase(notificationService notification.Service) UseCase {
	return &useCase{
		notificationService: notificationService,
	}
}

// GetNotifications retrieves user notifications
func (uc *useCase) GetNotifications(ctx context.Context, userID int) (*model.GetNotificationsResponse, error) {
	notificationCulturals, err := uc.notificationService.GetNotifications(ctx, userID)
	if err != nil {
		return nil, err
	}

	culturals := make([]model.CulturalList, len(notificationCulturals))
	for _, cultural := range notificationCulturals {
		culturals = append(culturals, model.CulturalList{
			ID:    cultural.ID,
			Title: cultural.Title,
		})
	}

	return &model.GetNotificationsResponse{
		UserID:   userID,
		Cultural: culturals,
	}, nil
}
