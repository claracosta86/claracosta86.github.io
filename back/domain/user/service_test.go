package user_test

import (
	"context"
	"errors"
	"testing"

	"poc2/back/domain/user"
	"poc2/back/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository(ctrl)
		service := user.NewService(mockRepo)

		mockRepo.EXPECT().FindByEmail(ctx, "test@example.com").Return(nil, nil) // User not found
		mockRepo.EXPECT().Save(ctx, gomock.Any()).Return(nil)

		err := service.RegisterUser(ctx, "Test User", "test@example.com", "12345678901", "Test Company", "password123", user.UserTypeCommon)
		assert.NoError(t, err)
	})

	t.Run("user already exists", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository(ctrl)
		service := user.NewService(mockRepo)

		existingUser := &user.User{ID: 1, Email: "test@example.com"}
		mockRepo.EXPECT().FindByEmail(ctx, "test@example.com").Return(existingUser, nil)

		err := service.RegisterUser(ctx, "Test User", "test@example.com", "12345678901", "Test Company", "password123", user.UserTypeCommon)
		assert.Error(t, err)
		assert.Equal(t, "user already exists with this email", err.Error())
	})
}

func TestService_AuthenticateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository(ctrl)
		service := user.NewService(mockRepo)

		u := &user.User{ID: 1, Email: "test@example.com", Password: "hashed_password"}
		mockRepo.EXPECT().FindByEmail(ctx, "test@example.com").Return(u, nil)
		mockRepo.EXPECT().CheckPassword(ctx, 1, "password123").Return(true, nil)

		result, err := service.AuthenticateUser(ctx, "test@example.com", "password123")
		assert.NoError(t, err)
		assert.Equal(t, u, result)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository(ctrl)
		service := user.NewService(mockRepo)

		mockRepo.EXPECT().FindByEmail(ctx, "wrong@example.com").Return(nil, errors.New("not found"))

		result, err := service.AuthenticateUser(ctx, "wrong@example.com", "password123")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("invalid password", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository(ctrl)
		service := user.NewService(mockRepo)

		u := &user.User{ID: 1, Email: "test@example.com", Password: "hashed_password"}
		mockRepo.EXPECT().FindByEmail(ctx, "test@example.com").Return(u, nil)
		mockRepo.EXPECT().CheckPassword(ctx, 1, "wrongpass").Return(false, nil)

		result, err := service.AuthenticateUser(ctx, "test@example.com", "wrongpass")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid password", err.Error())
	})
}

func TestService_GetCulturaisByOrganizerID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository(ctrl)
		service := user.NewService(mockRepo)

		organizer := &user.User{ID: 1, Type: string(user.UserTypeOrganizer)}
		culturais := []user.CulturalList{{ID: 1, Title: "Event 1"}}

		mockRepo.EXPECT().FindByID(ctx, 1).Return(organizer, nil)
		mockRepo.EXPECT().GetCulturaisByOrganizerID(ctx, 1).Return(culturais, nil)

		result, err := service.GetCulturaisByOrganizerID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, culturais, result)
	})

	t.Run("not an organizer", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository(ctrl)
		service := user.NewService(mockRepo)

		commonUser := &user.User{ID: 2, Type: string(user.UserTypeCommon)}

		mockRepo.EXPECT().FindByID(ctx, 2).Return(commonUser, nil)

		result, err := service.GetCulturaisByOrganizerID(ctx, 2)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "is not an organizer")
	})
}
