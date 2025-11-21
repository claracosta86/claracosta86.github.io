package notification_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	applicationNotification "poc2/back/application/notification"
	domainNotification "poc2/back/domain/notification"
	"poc2/back/mocks"
)

func TestGetNotifications(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotificationService := mocks.NewMockNotificationService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationNotification.NewUseCase(mockNotificationService, mockUserService)

	ctx := context.Background()
	userID := 1

	t.Run("success", func(t *testing.T) {
		notifications := []domainNotification.NotificationCulturalList{
			{
				ID:           1,
				CulturalID:   10,
				Title:        "Event Title",
				CulturalType: "event",
				Type:         domainNotification.NotificationTypeUpdated,
			},
		}

		mockNotificationService.EXPECT().GetNotifications(ctx, userID).Return(notifications, nil)

		resp, err := useCase.GetNotifications(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, userID, resp.UserID)
		assert.Len(t, resp.Cultural, 1)
		assert.Equal(t, notifications[0].Title, resp.Cultural[0].Title)
	})

	t.Run("service error", func(t *testing.T) {
		mockNotificationService.EXPECT().GetNotifications(ctx, userID).Return(nil, errors.New("service error"))

		resp, err := useCase.GetNotifications(ctx, userID)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestMarkNotificationsAsSeen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotificationService := mocks.NewMockNotificationService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationNotification.NewUseCase(mockNotificationService, mockUserService)

	ctx := context.Background()
	userID := 1
	notificationIDs := []int{1, 2, 3}

	t.Run("success", func(t *testing.T) {
		mockNotificationService.EXPECT().MarkNotificationsAsSeen(ctx, userID, notificationIDs).Return(nil)

		err := useCase.MarkNotificationsAsSeen(ctx, userID, notificationIDs)
		assert.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		mockNotificationService.EXPECT().MarkNotificationsAsSeen(ctx, userID, notificationIDs).Return(errors.New("service error"))

		err := useCase.MarkNotificationsAsSeen(ctx, userID, notificationIDs)
		assert.Error(t, err)
	})
}
