package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"poc2/back/application/user"
	"poc2/back/interface/http/handlers"
	userModel "poc2/back/interface/model"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUseCase is a mock of user.UseCase
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) RegisterUser(ctx context.Context, request userModel.RegisterUserRequest) error {
	args := m.Called(ctx, request)
	return args.Error(0)
}

func (m *MockUserUseCase) LoginUser(ctx context.Context, request userModel.LoginUserRequest) (*userModel.LoginUserResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userModel.LoginUserResponse), args.Error(1)
}

func (m *MockUserUseCase) GetUserProfile(ctx context.Context, userID int) (*userModel.UserProfileResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userModel.UserProfileResponse), args.Error(1)
}

func (m *MockUserUseCase) UpdateUserProfile(ctx context.Context, userID int, request userModel.UpdateUserProfileRequest) error {
	args := m.Called(ctx, userID, request)
	return args.Error(0)
}

func (m *MockUserUseCase) ChangePassword(ctx context.Context, userID int, request userModel.ChangePasswordRequest) error {
	args := m.Called(ctx, userID, request)
	return args.Error(0)
}

func (m *MockUserUseCase) DeleteUser(ctx context.Context, userID int, userType string) error {
	args := m.Called(ctx, userID, userType)
	return args.Error(0)
}

func (m *MockUserUseCase) ToggleFavorite(ctx context.Context, userID int, request userModel.FavoriteRequest) error {
	args := m.Called(ctx, userID, request)
	return args.Error(0)
}

func (m *MockUserUseCase) GetUserFavorites(ctx context.Context, userID int) (*userModel.FavoritesResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userModel.FavoritesResponse), args.Error(1)
}

func (m *MockUserUseCase) UpdateLastSeenFavorite(ctx context.Context, userID int, request userModel.FavoriteRequest) error {
	args := m.Called(ctx, userID, request)
	return args.Error(0)
}

func (m *MockUserUseCase) GetOrganizerCulturais(ctx context.Context, userID int) (*userModel.OrganizerCulturaisResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userModel.OrganizerCulturaisResponse), args.Error(1)
}

func (m *MockUserUseCase) GetOrganizerInfo(ctx context.Context, userID int) (*userModel.OrganizerInfoResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userModel.OrganizerInfoResponse), args.Error(1)
}

func TestHandleRegisterUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockUserUseCase)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.RegisterUserRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
			UserType: "common",
		})

		req, err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		mockUseCase.On("RegisterUser", mock.Anything, mock.AnythingOfType("userModel.RegisterUserRequest")).Return(nil)

		handler.HandleRegisterUser(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("user already exists", func(t *testing.T) {
		mockUseCase := new(MockUserUseCase)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.RegisterUserRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
			UserType: "common",
		})

		req, err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		mockUseCase.On("RegisterUser", mock.Anything, mock.AnythingOfType("userModel.RegisterUserRequest")).Return(errors.New("user already exists"))

		handler.HandleRegisterUser(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestHandleUserLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockUserUseCase)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.LoginUserRequest{
			Email:    "test@example.com",
			Password: "password",
		})

		req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		expectedResponse := &userModel.LoginUserResponse{
			UserID: 1,
			Name:   "Test User",
			Email:  "test@example.com",
		}

		mockUseCase.On("LoginUser", mock.Anything, mock.AnythingOfType("userModel.LoginUserRequest")).Return(expectedResponse, nil)

		handler.HandleUserLogin(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var actualResponse userModel.LoginUserResponse
		json.Unmarshal(rr.Body.Bytes(), &actualResponse)
		assert.Equal(t, *expectedResponse, actualResponse)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUseCase := new(MockUserUseCase)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.LoginUserRequest{
			Email:    "test@example.com",
			Password: "password",
		})

		req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		mockUseCase.On("LoginUser", mock.Anything, mock.AnythingOfType("userModel.LoginUserRequest")).Return(nil, errors.New("user not found"))

		handler.HandleUserLogin(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		mockUseCase := new(MockUserUseCase)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.LoginUserRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		})

		req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		mockUseCase.On("LoginUser", mock.Anything, mock.AnythingOfType("userModel.LoginUserRequest")).Return(nil, errors.New("invalid password"))

		handler.HandleUserLogin(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestHandleGetUserProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockUserUseCase)
		handler := handlers.NewUserHandler(mockUseCase)

		req, err := http.NewRequest("GET", "/users/1/profile/", nil)
		if err != nil {
			t.Fatal(err)
		}

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		expectedResponse := &userModel.UserProfileResponse{
			Name:  "Test User",
			Email: "test@example.com",
		}

		mockUseCase.On("GetUserProfile", mock.Anything, 1).Return(expectedResponse, nil)

		handler.HandleGetUserProfile(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var actualResponse userModel.UserProfileResponse
		json.Unmarshal(rr.Body.Bytes(), &actualResponse)
		assert.Equal(t, *expectedResponse, actualResponse)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUseCase := new(MockUserUseCase)
		handler := handlers.NewUserHandler(mockUseCase)

		req, err := http.NewRequest("GET", "/users/1/profile/", nil)
		if err != nil {
			t.Fatal(err)
		}

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.On("GetUserProfile", mock.Anything, 1).Return(nil, errors.New("user not found"))

		handler.HandleGetUserProfile(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}
