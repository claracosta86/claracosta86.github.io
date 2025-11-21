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
	commentModel "poc2/back/interface/model"
	mock "poc2/back/mocks"
)

func TestHandleCreateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCommentUseCase(ctrl)
		handler := handlers.NewCommentHandler(mockUseCase)

		requestBody, _ := json.Marshal(commentModel.CreateCommentRequest{
			UserID:       1,
			CulturalID:   1,
			CulturalType: "event",
			Comment:      "Great event!",
		})

		req, _ := http.NewRequest("POST", "/comments/", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().CreateComment(gomock.Any(), gomock.AssignableToTypeOf(commentModel.CreateCommentRequest{})).Return(nil)

		handler.HandleCreateComment(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("cultural not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCommentUseCase(ctrl)
		handler := handlers.NewCommentHandler(mockUseCase)

		requestBody, _ := json.Marshal(commentModel.CreateCommentRequest{
			UserID:       1,
			CulturalID:   99,
			CulturalType: "event",
			Comment:      "Great event!",
		})

		req, _ := http.NewRequest("POST", "/comments/", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().CreateComment(gomock.Any(), gomock.AssignableToTypeOf(commentModel.CreateCommentRequest{})).Return(errors.New("cultural not found"))

		handler.HandleCreateComment(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestHandleGetComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCommentUseCase(ctrl)
		handler := handlers.NewCommentHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/comments/event/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("culturalType", "event")
		chiCtx.URLParams.Add("culturalID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		expectedResponse := commentModel.GetCommentsResponse{
			Comments: []commentModel.Comment{
				{UserName: "User1", Comment: "Comment 1"},
			},
		}
		mockUseCase.EXPECT().GetComments(gomock.Any(), 1, "event").Return(expectedResponse, nil)

		handler.HandleGetComment(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp commentModel.GetCommentsResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, expectedResponse, resp)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCommentUseCase(ctrl)
		handler := handlers.NewCommentHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/comments/event/99", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("culturalType", "event")
		chiCtx.URLParams.Add("culturalID", "99")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.EXPECT().GetComments(gomock.Any(), 99, "event").Return(commentModel.GetCommentsResponse{}, errors.New("comments not found"))

		handler.HandleGetComment(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
