package notification

import (
	"context"

	"poc2/back/domain/notification"
	"poc2/back/domain/user"
	"poc2/back/interface/model"
)

// UseCase defines the application use cases for user operations
type UseCase interface {
	// RegisterUser handles user registration
	GetNotifications(ctx context.Context, userID int) (*model.GetNotificationsResponse, error)

	// MarkNotificationsAsSeen marks user notifications as seen
	MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error
}

type useCase struct {
	notificationService notification.Service
	userService         user.Service
}

// NewUseCase creates a new user use case
func NewUseCase(notificationService notification.Service, userService user.Service) UseCase {
	return &useCase{
		notificationService: notificationService,
		userService:         userService,
	}
}

// GetNotifications retrieves user notifications
func (uc *useCase) GetNotifications(ctx context.Context, userID int) (*model.GetNotificationsResponse, error) {
	notificationCulturals, err := uc.notificationService.GetNotifications(ctx, userID)
	if err != nil {
		return nil, err
	}

	culturals := make([]model.NotificationCulturalList, 0)
	for _, cultural := range notificationCulturals {
		notificationType, err := notification.NewNotificationType(cultural.Type.String())
		if err != nil {
			return nil, err
		}

		culturals = append(culturals, model.NotificationCulturalList{
			ID:               cultural.CulturalID,
			Title:            cultural.Title,
			CulturalType:     cultural.CulturalType,
			NotificationType: notificationType.String(),
			NotificationID:   cultural.ID,
		})
	}

	return &model.GetNotificationsResponse{
		UserID:   userID,
		Cultural: culturals,
	}, nil
}

func (uc *useCase) MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	return uc.notificationService.MarkNotificationsAsSeen(ctx, userID, notificationIDs)
}
