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
	notificationModel "poc2/back/interface/model"
	mock "poc2/back/mocks"
)

func TestHandleGetUserNotifications(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockNotificationUseCase(ctrl)
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

		mockUseCase.EXPECT().GetNotifications(gomock.Any(), 1).Return(&expectedResponse, nil)

		handler.HandleGetUserNotifications(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp notificationModel.GetNotificationsResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, expectedResponse, resp)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockNotificationUseCase(ctrl)
		handler := handlers.NewNotificationHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/notifications/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.EXPECT().GetNotifications(gomock.Any(), 1).Return(nil, errors.New("some error"))

		handler.HandleGetUserNotifications(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandleMarkNotificationsAsSeen(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockNotificationUseCase(ctrl)
		handler := handlers.NewNotificationHandler(mockUseCase)

		requestBody, _ := json.Marshal(notificationModel.SeenNotificationPost{
			NotificationIDs: []int{1, 2},
		})

		req, _ := http.NewRequest("PATCH", "/notifications/1/seen", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.EXPECT().MarkNotificationsAsSeen(gomock.Any(), 1, []int{1, 2}).Return(nil)

		handler.HandleMarkNotificationsAsSeen(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockNotificationUseCase(ctrl)
		handler := handlers.NewNotificationHandler(mockUseCase)

		requestBody, _ := json.Marshal(notificationModel.SeenNotificationPost{
			NotificationIDs: []int{1, 2},
		})

		req, _ := http.NewRequest("PATCH", "/notifications/1/seen", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("userID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.EXPECT().MarkNotificationsAsSeen(gomock.Any(), 1, []int{1, 2}).Return(errors.New("some error"))

		handler.HandleMarkNotificationsAsSeen(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
