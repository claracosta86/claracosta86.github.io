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

// MockCommentaryUseCase is a mock of commentary.UseCase
type MockCommentaryUseCase struct {
	mock.Mock
}

func (m *MockCommentaryUseCase) CreateCommentary(ctx context.Context, req commentaryModel.CreateCommentaryRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockCommentaryUseCase) GetCommentaries(ctx context.Context, culturalID int, culturalType string) ([]commentaryModel.CommentaryResponse, error) {
	args := m.Called(ctx, culturalID, culturalType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]commentaryModel.CommentaryResponse), args.Error(1)
}

func TestHandleCreateCommentary(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockCommentaryUseCase)
		handler := handlers.NewCommentaryHandler(mockUseCase)

		requestBody, _ := json.Marshal(commentaryModel.CreateCommentaryRequest{
			UserID:       1,
			CulturalID:   1,
			CulturalType: "event",
			Commentary:   "Great event!",
		})

		req, _ := http.NewRequest("POST", "/commentarys/", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		mockUseCase.On("CreateCommentary", mock.Anything, mock.AnythingOfType("commentaryModel.CreateCommentaryRequest")).Return(nil)

		handler.HandleCreateCommentary(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("cultural not found", func(t *testing.T) {
		mockUseCase := new(MockCommentaryUseCase)
		handler := handlers.NewCommentaryHandler(mockUseCase)

		requestBody, _ := json.Marshal(commentaryModel.CreateCommentaryRequest{
			UserID:       1,
			CulturalID:   99,
			CulturalType: "event",
			Commentary:   "Great event!",
		})

		req, _ := http.NewRequest("POST", "/commentarys/", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		mockUseCase.On("CreateCommentary", mock.Anything, mock.AnythingOfType("commentaryModel.CreateCommentaryRequest")).Return(errors.New("cultural not found"))

		handler.HandleCreateCommentary(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestHandleGetCommentary(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(MockCommentaryUseCase)
		handler := handlers.NewCommentaryHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/commentarys/event/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("culturalType", "event")
		chiCtx.URLParams.Add("culturalID", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		expectedResponse := []commentaryModel.CommentaryResponse{
			{UserName: "User1", Comment: "Comment 1"},
		}
		mockUseCase.On("GetCommentaries", mock.Anything, 1, "event").Return(expectedResponse, nil)

		handler.HandleGetCommentary(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp []commentaryModel.CommentaryResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, expectedResponse, resp)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUseCase := new(MockCommentaryUseCase)
		handler := handlers.NewCommentaryHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/commentarys/event/99", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("culturalType", "event")
		chiCtx.URLParams.Add("culturalID", "99")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.On("GetCommentaries", mock.Anything, 99, "event").Return(nil, errors.New("commentaries not found"))

		handler.HandleGetCommentary(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}
