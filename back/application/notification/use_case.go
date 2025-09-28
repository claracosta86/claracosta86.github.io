package notification

import (
	"context"

	"poc2/back/domain/notification"
	"poc2/back/interface/model"
	"poc2/back/domain/user"
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
	userService    user.Service
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
	favorites, err := uc.userService.GetUserFavorites(ctx, userID)
	if err != nil {
		return nil, err
	}

	favoritesMap := make(map[int]string)
	for _, favorite := range favorites {
		favoritesMap[favorite.ID] = favorite.Type
	}

	notificationCulturals, err := uc.notificationService.GetNotifications(ctx, userID, favoritesMap)
	if err != nil {
		return nil, err
	}

	culturals := make([]model.CulturalList, len(notificationCulturals))
	for _, cultural := range notificationCulturals {
		culturals = append(culturals, model.CulturalList{
			ID:    cultural.CulturalID,
			Title: cultural.Title,
			Type:  cultural.CulturalType,
			NotificationType: cultural.Type,
			NotificationID: cultural.ID,
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