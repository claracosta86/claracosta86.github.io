package notification_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"poc2/back/domain/notification"
	"poc2/back/mocks"
)

func TestGetNotifications(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockNotificationRepository(ctrl)
	service := notification.NewService(mockRepo)

	ctx := context.Background()
	userID := 1
	expectedNotifications := []notification.NotificationCulturalList{
		{
			CulturalID:   1,
			CulturalType: "event",
			Title:        "New event",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindByUserID(ctx, userID).Return(expectedNotifications, nil)

		notifications, err := service.GetNotifications(ctx, userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedNotifications, notifications)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindByUserID(ctx, userID).Return(nil, errors.New("repository error"))

		notifications, err := service.GetNotifications(ctx, userID)

		assert.Error(t, err)
		assert.Nil(t, notifications)
	})
}

func TestMarkNotificationsAsSeen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockNotificationRepository(ctrl)
	service := notification.NewService(mockRepo)

	ctx := context.Background()
	userID := 1
	notificationIDs := []int{1, 2, 3}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().MarkAsSeen(ctx, userID, notificationIDs).Return(nil)

		err := service.MarkNotificationsAsSeen(ctx, userID, notificationIDs)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().MarkAsSeen(ctx, userID, notificationIDs).Return(errors.New("repository error"))

		err := service.MarkNotificationsAsSeen(ctx, userID, notificationIDs)

		assert.Error(t, err)
	})
}
