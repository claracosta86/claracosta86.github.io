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
	commentaryModel "poc2/back/interface/model"
)

// MockCommentUseCase is a mock of comment.UseCase
type MockCommentUseCase struct {
	mock.Mock
}

func (m *MockCommentUseCase) CreateComment(ctx context.Context, req commentaryModel.CreateCommentRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockCommentUseCase) GetComments(ctx context.Context, culturalID int, culturalType string) (commentaryModel.GetCommentsResponse, error) {
	args := m.Called(ctx, culturalID, culturalType)
	if args.Get(0) == nil {
		return commentaryModel.GetCommentsResponse{}, args.Error(1)
	}
	return args.Get(0).(commentaryModel.GetCommentsResponse), args.Error(1)
}

func TestHandleCreateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockCommentUseCase)
		handler := handlers.NewCommentHandler(mockUseCase)

		requestBody, _ := json.Marshal(commentaryModel.CreateCommentRequest{
			UserID:       1,
			CulturalID:   1,
			CulturalType: "event",
			Comment:      "Great event!",
		})

		req, _ := http.NewRequest("POST", "/comments/", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		mockUseCase.On("CreateComment", mock.Anything, mock.AnythingOfType("commentaryModel.CreateCommentRequest")).Return(nil)

		handler.HandleCreateComment(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("cultural not found", func(t *testing.T) {
		mockUseCase := new(MockCommentUseCase)
		handler := handlers.NewCommentHandler(mockUseCase)

		requestBody, _ := json.Marshal(commentaryModel.CreateCommentRequest{
			UserID:       1,
			CulturalID:   99,
			CulturalType: "event",
			Comment:      "Great event!",
		})

		req, _ := http.NewRequest("POST", "/comments/", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		mockUseCase.On("CreateComment", mock.Anything, mock.AnythingOfType("commentaryModel.CreateCommentRequest")).Return(errors.New("cultural not found"))

		handler.HandleCreateComment(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestHandleGetComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockCommentUseCase)
		handler := handlers.NewCommentHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/comments/event/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("culturalType", "event")
		chiCtx.URLParams.Add("culturalID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		expectedResponse := commentaryModel.GetCommentsResponse{
			Comments: []commentaryModel.Comment{
				{UserName: "User1", Comment: "Comment 1"},
			},
		}
		mockUseCase.On("GetComments", mock.Anything, 1, "event").Return(expectedResponse, nil)

		handler.HandleGetComment(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp commentaryModel.GetCommentsResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, expectedResponse, resp)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUseCase := new(MockCommentUseCase)
		handler := handlers.NewCommentHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/comments/event/99", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("culturalType", "event")
		chiCtx.URLParams.Add("culturalID", "99")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.On("GetComments", mock.Anything, 99, "event").Return(nil, errors.New("commentaries not found"))

		handler.HandleGetComment(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}
