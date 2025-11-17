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
	"github.com/stretchr/testify/mock"

	"poc2/back/interface/http/handlers"
	notificationModel "poc2/back/interface/model"
)

// MockNotificationUseCase is a mock of notification.UseCase
type MockNotificationUseCase struct {
	mock.Mock
}

func (m *MockNotificationUseCase) GetNotifications(ctx context.Context, userID int) (*notificationModel.GetNotificationsResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*notificationModel.GetNotificationsResponse), args.Error(1)
}

func (m *MockNotificationUseCase) MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	args := m.Called(ctx, userID, notificationIDs)
	return args.Error(0)
}

func TestHandleGetUserNotifications(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockNotificationUseCase)
		handler := handlers.NewNotificationHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/notifications/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		expectedResponse := notificationModel.GetNotificationsResponse{
			Cultural: []notificationModel.NotificationCulturalList{
				{ID: 1, NotificationType: "update"},
			},
		}
		mockUseCase.On("GetNotifications", mock.Anything, 1).Return(expectedResponse, nil)

		handler.HandleGetUserNotifications(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp notificationModel.GetNotificationsResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, expectedResponse, resp)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockUseCase := new(MockNotificationUseCase)
		handler := handlers.NewNotificationHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/notifications/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.On("GetNotifications", mock.Anything, 1).Return(nil, errors.New("some error"))

		handler.HandleGetUserNotifications(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestHandleMarkNotificationsAsSeen(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockNotificationUseCase)
		handler := handlers.NewNotificationHandler(mockUseCase)

		requestBody, _ := json.Marshal(notificationModel.SeenNotificationsRequest{
			NotificationIDs: []int{1, 2},
		})

		req, _ := http.NewRequest("PATCH", "/notifications/1/seen", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.On("MarkNotificationsAsSeen", mock.Anything, 1, []int{1, 2}).Return(nil)

		handler.HandleMarkNotificationsAsSeen(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockUseCase := new(MockNotificationUseCase)
		handler := handlers.NewNotificationHandler(mockUseCase)

		requestBody, _ := json.Marshal(notificationModel.SeenNotificationsRequest{
			NotificationIDs: []int{1, 2},
		})

		req, _ := http.NewRequest("PATCH", "/notifications/1/seen", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.On("MarkNotificationsAsSeen", mock.Anything, 1, []int{1, 2}).Return(errors.New("some error"))

		handler.HandleMarkNotificationsAsSeen(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}
