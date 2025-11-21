package user_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	applicationUser "poc2/back/application/user"
	domainUser "poc2/back/domain/user"
	"poc2/back/interface/model"
	"poc2/back/mocks"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	req := model.RegisterUserRequest{
		Name:        "Test User",
		Email:       "test@example.com",
		Document:    "12345678900",
		CompanyName: "Test Company",
		Password:    "password",
		Type:        "common",
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().RegisterUser(
			ctx,
			req.Name,
			req.Email,
			req.Document,
			req.CompanyName,
			req.Password,
			domainUser.UserTypeCommon,
		).Return(nil)

		err := useCase.RegisterUser(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("invalid user type", func(t *testing.T) {
		invalidReq := req
		invalidReq.Type = "invalid"

		err := useCase.RegisterUser(ctx, invalidReq)
		assert.Error(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		mockUserService.EXPECT().RegisterUser(
			ctx,
			req.Name,
			req.Email,
			req.Document,
			req.CompanyName,
			req.Password,
			domainUser.UserTypeCommon,
		).Return(errors.New("service error"))

		err := useCase.RegisterUser(ctx, req)
		assert.Error(t, err)
		assert.Equal(t, "service error", err.Error())
	})
}

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	req := model.LoginUserRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	expectedUser := &domainUser.User{
		ID:   1,
		Type: "common",
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().AuthenticateUser(ctx, req.Email, req.Password).Return(expectedUser, nil)

		resp, err := useCase.LoginUser(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, resp.UserID)
		assert.Equal(t, expectedUser.Type, resp.Type)
	})

	t.Run("service error", func(t *testing.T) {
		mockUserService.EXPECT().AuthenticateUser(ctx, req.Email, req.Password).Return(nil, errors.New("auth error"))

		resp, err := useCase.LoginUser(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestGetUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	userID := 1
	expectedUser := &domainUser.User{
		ID:          userID,
		Name:        "Test User",
		Email:       "test@example.com",
		CompanyName: "Test Company",
		Type:        "common",
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().GetUserByID(ctx, userID).Return(expectedUser, nil)

		resp, err := useCase.GetUserProfile(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, resp.UserID)
		assert.Equal(t, expectedUser.Name, resp.Name)
	})

	t.Run("service error", func(t *testing.T) {
		mockUserService.EXPECT().GetUserByID(ctx, userID).Return(nil, errors.New("user not found"))

		resp, err := useCase.GetUserProfile(ctx, userID)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestUpdateUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	userID := 1
	req := model.UpdateUserProfileRequest{
		Name:        "Updated Name",
		Email:       "updated@example.com",
		CompanyName: "Updated Company",
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().UpdateUserProfile(ctx, userID, req.Name, req.Email, req.CompanyName).Return(nil)

		err := useCase.UpdateUserProfile(ctx, userID, req)
		assert.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		mockUserService.EXPECT().UpdateUserProfile(ctx, userID, req.Name, req.Email, req.CompanyName).Return(errors.New("update error"))

		err := useCase.UpdateUserProfile(ctx, userID, req)
		assert.Error(t, err)
	})
}

func TestChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	userID := 1
	req := model.ChangePasswordRequest{
		CurrentPassword: "old",
		NewPassword:     "new",
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().ChangeUserPassword(ctx, userID, req.CurrentPassword, req.NewPassword).Return(nil)

		err := useCase.ChangePassword(ctx, userID, req)
		assert.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		mockUserService.EXPECT().ChangeUserPassword(ctx, userID, req.CurrentPassword, req.NewPassword).Return(errors.New("change error"))

		err := useCase.ChangePassword(ctx, userID, req)
		assert.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	userID := 1

	t.Run("common user success", func(t *testing.T) {
		mockUserService.EXPECT().DeleteUser(ctx, userID).Return(nil)

		err := useCase.DeleteUser(ctx, userID, "common")
		assert.NoError(t, err)
	})

	t.Run("organizer user success", func(t *testing.T) {
		eventIDs := []int{1, 2}
		attractionIDs := []int{3, 4}

		mockCulturalService.EXPECT().GetEventsIDsByOrganizer(ctx, userID).Return(eventIDs, nil)
		mockCulturalService.EXPECT().GetTouristAttractionsIDsByOrganizer(ctx, userID).Return(attractionIDs, nil)
		mockUserService.EXPECT().RemoveEventFromAllUsers(ctx, eventIDs).Return(nil)
		mockUserService.EXPECT().RemoveTouristAttractionFromAllUsers(ctx, attractionIDs).Return(nil)
		mockUserService.EXPECT().DeleteUser(ctx, userID).Return(nil)

		err := useCase.DeleteUser(ctx, userID, "organizer")
		assert.NoError(t, err)
	})

	t.Run("organizer user get events error", func(t *testing.T) {
		mockCulturalService.EXPECT().GetEventsIDsByOrganizer(ctx, userID).Return(nil, errors.New("error"))

		err := useCase.DeleteUser(ctx, userID, "organizer")
		assert.Error(t, err)
	})
}

func TestToggleFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	userID := 1
	req := model.FavoriteRequest{
		CulturalID:   10,
		CulturalType: "event",
		IsFavorite:   true,
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().ToggleFavorite(ctx, userID, req.CulturalType, req.CulturalID, req.IsFavorite).Return(nil)

		err := useCase.ToggleFavorite(ctx, userID, req)
		assert.NoError(t, err)
	})
}

func TestGetUserFavorites(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	userID := 1
	favorites := []domainUser.CulturalList{
		{ID: 1, Type: "event"},
		{ID: 2, Type: "attraction"},
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().GetUserFavorites(ctx, userID).Return(favorites, nil)

		resp, err := useCase.GetUserFavorites(ctx, userID)
		assert.NoError(t, err)
		assert.Len(t, resp, 2)
		assert.Equal(t, 1, resp[0].ID)
	})
}

func TestUpdateLastSeenFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	userID := 1
	req := model.FavoriteRequest{
		CulturalID:   10,
		CulturalType: "event",
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().UpdateLastSeenFavorite(ctx, userID, req.CulturalID, req.CulturalType).Return(nil)

		err := useCase.UpdateLastSeenFavorite(ctx, userID, req)
		assert.NoError(t, err)
	})
}

func TestGetOrganizerCulturais(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	organizerID := 1
	culturais := []domainUser.CulturalList{
		{ID: 1, Type: "event"},
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().GetCulturaisByOrganizerID(ctx, organizerID).Return(culturais, nil)

		resp, err := useCase.GetOrganizerCulturais(ctx, organizerID)
		assert.NoError(t, err)
		assert.Len(t, resp, 1)
	})
}

func TestGetOrganizerInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationUser.NewUseCase(mockUserService, mockCulturalService)

	ctx := context.Background()
	organizerID := 1
	createdAt := time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05") // 1 year ago
	organizer := &domainUser.User{
		ID:        organizerID,
		Name:      "Organizer",
		Email:     "org@example.com",
		CreatedAt: createdAt,
	}
	culturais := []domainUser.CulturalList{
		{ID: 1, Type: "event"},
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.EXPECT().GetUserByID(ctx, organizerID).Return(organizer, nil)
		mockUserService.EXPECT().GetCulturaisByOrganizerID(ctx, organizerID).Return(culturais, nil)

		resp, err := useCase.GetOrganizerInfo(ctx, organizerID)
		assert.NoError(t, err)
		assert.Equal(t, organizer.Name, resp.Name)
		assert.Equal(t, "1 ano", resp.OrganizerSince)
	})
}
