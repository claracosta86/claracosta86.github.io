package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"poc2/back/interface/http/handlers"
	culturalModel "poc2/back/interface/model"
)

// MockCulturalUseCase is a mock of cultural.UseCase
type MockCulturalUseCase struct {
	mock.Mock
}

func (m *MockCulturalUseCase) CreateCultural(ctx context.Context, req culturalModel.CreateCulturalRequest) (culturalModel.CreateCulturalResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return culturalModel.CreateCulturalResponse{}, args.Error(1)
	}
	return args.Get(0).(culturalModel.CreateCulturalResponse), args.Error(1)
}

func (m *MockCulturalUseCase) GetCultural(ctx context.Context, id int, culturalType string) (culturalModel.GetCulturalResponse, error) {
	args := m.Called(ctx, id, culturalType)
	if args.Get(0) == nil {
		return culturalModel.GetCulturalResponse{}, args.Error(1)
	}
	return args.Get(0).(culturalModel.GetCulturalResponse), args.Error(1)
}

func (m *MockCulturalUseCase) UpdateCultural(ctx context.Context, req culturalModel.UpdateCulturalRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockCulturalUseCase) DeleteCultural(ctx context.Context, id int, culturalType string) error {
	args := m.Called(ctx, id, culturalType)
	return args.Error(0)
}

func (m *MockCulturalUseCase) GetAllCulturais(ctx context.Context) (culturalModel.GetAllCulturaisResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return culturalModel.GetAllCulturaisResponse{}, args.Error(1)
	}
	return args.Get(0).(culturalModel.GetAllCulturaisResponse), args.Error(1)
}

func (m *MockCulturalUseCase) GetHomeCulturais(ctx context.Context) (culturalModel.GetAllCulturaisResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return culturalModel.GetAllCulturaisResponse{}, args.Error(1)
	}
	return args.Get(0).(culturalModel.GetAllCulturaisResponse), args.Error(1)
}

func TestHandleCreateCultural(t *testing.T) {
	// Create a temporary directory for uploads
	uploadDir := t.TempDir()
	staticDir := filepath.Join(uploadDir, "static")
	thumbDir := filepath.Join(staticDir, "culturalthumbs")
	os.MkdirAll(thumbDir, 0755)

	// Change working directory to the temp dir to handle relative paths
	originalWd, _ := os.Getwd()
	os.Chdir(uploadDir)
	defer os.Chdir(originalWd)

	t.Run("success with image", func(t *testing.T) {
		mockUseCase := new(MockCulturalUseCase)
		handler := handlers.NewCulturalHandler(mockUseCase)

		// Create multipart form
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Add JSON data
		jsonData, _ := json.Marshal(culturalModel.CreateCulturalRequest{
			Title: "Test Event",
			Type:  "event",
		})
		_ = writer.WriteField("data", string(jsonData))

		// Add image file
		part, _ := writer.CreateFormFile("image", "test.jpg")
		_, _ = io.WriteString(part, "fake image data")
		writer.Close()

		req, _ := http.NewRequest("POST", "/cultural", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		rr := httptest.NewRecorder()

		expectedResponse := culturalModel.CreateCulturalResponse{ID: 1, Type: "event"}
		mockUseCase.On("CreateCultural", mock.Anything, mock.AnythingOfType("culturalModel.CreateCulturalRequest")).Return(expectedResponse, nil)

		handler.HandleCreateCultural(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		var resp map[string]string
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, "Cultural created successfully", resp["status"])
		assert.Equal(t, "1", resp["id"])
		assert.Equal(t, "event", resp["type"])

		mockUseCase.AssertExpectations(t)
	})

	t.Run("success without image", func(t *testing.T) {
		mockUseCase := new(MockCulturalUseCase)
		handler := handlers.NewCulturalHandler(mockUseCase)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		jsonData, _ := json.Marshal(culturalModel.CreateCulturalRequest{
			Title: "Test Attraction",
			Type:  "tourist_attraction",
		})
		_ = writer.WriteField("data", string(jsonData))
		writer.Close()

		req, _ := http.NewRequest("POST", "/cultural", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		rr := httptest.NewRecorder()

		expectedResponse := &culturalModel.CreateCulturalResponse{ID: 2, Type: "tourist_attraction"}
		mockUseCase.On("CreateCultural", mock.Anything, mock.AnythingOfType("culturalModel.CreateCulturalRequest")).Return(expectedResponse, nil)

		handler.HandleCreateCultural(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("use case error", func(t *testing.T) {
		mockUseCase := new(MockCulturalUseCase)
		handler := handlers.NewCulturalHandler(mockUseCase)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		jsonData, _ := json.Marshal(culturalModel.CreateCulturalRequest{
			Title: "Test Error",
			Type:  "event",
		})
		_ = writer.WriteField("data", string(jsonData))
		writer.Close()

		req, _ := http.NewRequest("POST", "/cultural", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		rr := httptest.NewRecorder()

		mockUseCase.On("CreateCultural", mock.Anything, mock.AnythingOfType("culturalModel.CreateCulturalRequest")).Return(nil, errors.New("some internal error"))

		handler.HandleCreateCultural(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestHandleGetCultural(t *testing.T) {
	t.Run("success getting event", func(t *testing.T) {
		mockUseCase := new(MockCulturalUseCase)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/event/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "event")
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		expectedResponse := map[string]interface{}{"name": "Test Event"}
		mockUseCase.On("GetCultural", mock.Anything, 1, "event").Return(expectedResponse, nil)

		handler.HandleGetCultural(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, "Test Event", resp["name"])
		mockUseCase.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUseCase := new(MockCulturalUseCase)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/event/99", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "event")
		chiCtx.URLParams.Add("id", "99")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.On("GetCultural", mock.Anything, 99, "event").Return(nil, errors.New("cultural event not found"))

		handler.HandleGetCultural(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("invalid id", func(t *testing.T) {
		mockUseCase := new(MockCulturalUseCase)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/event/abc", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "event")
		chiCtx.URLParams.Add("id", "abc")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		handler.HandleGetCultural(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("invalid type", func(t *testing.T) {
		mockUseCase := new(MockCulturalUseCase)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/invalidtype/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "invalidtype")
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		handler.HandleGetCultural(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
