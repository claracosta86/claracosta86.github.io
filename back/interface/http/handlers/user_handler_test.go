package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"poc2/back/interface/http/handlers"
	userModel "poc2/back/interface/model"
	mock "poc2/back/mocks"
)

func TestHandleRegisterUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.RegisterUserRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
			Type:     "common",
		})

		req, err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(nil)
		handler.HandleRegisterUser(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("user already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.RegisterUserRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password",
			Type:     "common",
		})

		req, err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(errors.New("user already exists"))

		handler.HandleRegisterUser(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)

	})
}

func TestHandleUserLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
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
			Type:   "common",
		}

		mockUseCase.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(expectedResponse, nil)

		handler.HandleUserLogin(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var actualResponse userModel.LoginUserResponse
		json.Unmarshal(rr.Body.Bytes(), &actualResponse)
		assert.Equal(t, *expectedResponse, actualResponse)

	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
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

		mockUseCase.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("user not found"))

		handler.HandleUserLogin(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

	})

	t.Run("invalid password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
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

		mockUseCase.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("invalid password"))

		handler.HandleUserLogin(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)

	})
}

func TestHandleGetUserProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, err := http.NewRequest("GET", "/users/1/profile/", nil)
		if err != nil {
			t.Fatal(err)
		}

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		expectedResponse := &userModel.GetUserProfileResponse{
			Name:  "Test User",
			Email: "test@example.com",
		}

		mockUseCase.EXPECT().GetUserProfile(gomock.Any(), 1).Return(expectedResponse, nil)

		handler.HandleGetUserProfile(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var actualResponse userModel.GetUserProfileResponse
		json.Unmarshal(rr.Body.Bytes(), &actualResponse)
		assert.Equal(t, *expectedResponse, actualResponse)

	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, err := http.NewRequest("GET", "/users/1/profile/", nil)
		if err != nil {
			t.Fatal(err)
		}

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().GetUserProfile(gomock.Any(), 1).Return(nil, errors.New("user not found"))

		handler.HandleGetUserProfile(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

	})
}

func TestHandleEditUserProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.UpdateUserProfileRequest{
			Name: "Updated Name",
		})
		req, _ := http.NewRequest("PATCH", "/users/1/profile/edit", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateUserProfile(gomock.Any(), 1, gomock.Any()).Return(nil)

		handler.HandleEditUserProfile(rr, req)
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.UpdateUserProfileRequest{
			Name: "Updated Name",
		})
		req, _ := http.NewRequest("PATCH", "/users/1/profile/edit", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateUserProfile(gomock.Any(), 1, gomock.Any()).Return(errors.New("user not found"))

		handler.HandleEditUserProfile(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestHandleChangeUserPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.ChangePasswordRequest{
			CurrentPassword: "old",
			NewPassword:     "new",
		})
		req, _ := http.NewRequest("PATCH", "/users/1/profile/change-password", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().ChangePassword(gomock.Any(), 1, gomock.Any()).Return(nil)

		handler.HandleChangeUserPassword(rr, req)
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("incorrect password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.ChangePasswordRequest{
			CurrentPassword: "old",
			NewPassword:     "new",
		})
		req, _ := http.NewRequest("PATCH", "/users/1/profile/change-password", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().ChangePassword(gomock.Any(), 1, gomock.Any()).Return(errors.New("current password is incorrect"))

		handler.HandleChangeUserPassword(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}

func TestHandleDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("DELETE", "/users/1/profile/delete", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		chiCtx.URLParams.Add("userType", "common")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().DeleteUser(gomock.Any(), 1, "common").Return(nil)

		handler.HandleDeleteUser(rr, req)
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("DELETE", "/users/1/profile/delete", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		chiCtx.URLParams.Add("userType", "common")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().DeleteUser(gomock.Any(), 1, "common").Return(errors.New("user not found"))

		handler.HandleDeleteUser(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestHandleUpdateFavorites(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.FavoriteRequest{
			CulturalID: 1,
		})
		req, _ := http.NewRequest("PATCH", "/users/1/profile/favorites", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateFavorites(gomock.Any(), 1, gomock.Any()).Return(nil)

		handler.HandleUpdateFavorites(rr, req)
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.FavoriteRequest{
			CulturalID: 1,
		})
		req, _ := http.NewRequest("PATCH", "/users/1/profile/favorites", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateFavorites(gomock.Any(), 1, gomock.Any()).Return(errors.New("user not found"))

		handler.HandleUpdateFavorites(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestHandleGetUserFavorites(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/users/1/favorites", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		expectedResponse := []userModel.CulturalList{{ID: 1}}
		mockUseCase.EXPECT().GetUserFavorites(gomock.Any(), 1).Return(expectedResponse, nil)

		handler.HandleGetUserFavorites(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/users/1/favorites", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().GetUserFavorites(gomock.Any(), 1).Return(nil, errors.New("user not found"))

		handler.HandleGetUserFavorites(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestHandleLastSeenFavorite(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.FavoriteRequest{
			CulturalID: 1,
		})
		req, _ := http.NewRequest("PATCH", "/users/favorites/last-seen", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateLastSeenFavorite(gomock.Any(), 1, gomock.Any()).Return(nil)

		handler.HandleLastSeenFavorite(rr, req)
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		requestBody, _ := json.Marshal(userModel.FavoriteRequest{
			CulturalID: 1,
		})
		req, _ := http.NewRequest("PATCH", "/users/favorites/last-seen", bytes.NewBuffer(requestBody))
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateLastSeenFavorite(gomock.Any(), 1, gomock.Any()).Return(errors.New("user not found"))

		handler.HandleLastSeenFavorite(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestHandleGetOrganizerCulturais(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/users/culturais/", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		expectedResponse := []userModel.CulturalList{{ID: 1}}
		mockUseCase.EXPECT().GetOrganizerCulturais(gomock.Any(), 1).Return(expectedResponse, nil)

		handler.HandleGetOrganizerCulturais(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/users/culturais/", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().GetOrganizerCulturais(gomock.Any(), 1).Return(nil, errors.New("user not found"))

		handler.HandleGetOrganizerCulturais(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestHandleGetOrganizerInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/users/1/info", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		expectedResponse := &userModel.GetOrganizerInfoResponse{Name: "Organizer"}
		mockUseCase.EXPECT().GetOrganizerInfo(gomock.Any(), 1).Return(expectedResponse, nil)

		handler.HandleGetOrganizerInfo(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockUserUseCase(ctrl)
		handler := handlers.NewUserHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/users/1/info", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().GetOrganizerInfo(gomock.Any(), 1).Return(nil, errors.New("user not found"))

		handler.HandleGetOrganizerInfo(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
