package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"poc2/back/interface/http/handlers"
	culturalModel "poc2/back/interface/model"
	mock "poc2/back/mocks"
)

func TestHandleCreateCultural(t *testing.T) {
	t.Run("success without image", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		requestBody, _ := json.Marshal(culturalModel.CreateCulturalRequest{
			Title: "Test Event No Image",
			Type:  "event",
		})
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("data", string(requestBody))
		writer.Close()

		req, _ := http.NewRequest("POST", "/cultural", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		expectedResponse := culturalModel.CreateCulturalResponse{ID: 2, Type: "event"}

		mockUseCase.EXPECT().CreateCultural(gomock.Any(), gomock.Any()).Return(expectedResponse, nil)

		handler.HandleCreateCultural(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("use case error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
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

		mockUseCase.EXPECT().CreateCultural(gomock.Any(), gomock.AssignableToTypeOf(culturalModel.CreateCulturalRequest{})).Return(culturalModel.CreateCulturalResponse{}, errors.New("use case error"))

		handler.HandleCreateCultural(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandleGetCultural(t *testing.T) {
	t.Run("success getting event", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/event/1", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "event")
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		expectedResponse := culturalModel.GetCulturalResponse{
			ID:    1,
			Title: "Test Event",
		}

		mockUseCase.EXPECT().GetCultural(gomock.Any(), 1, "event").Return(expectedResponse, nil)

		handler.HandleGetCultural(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, "Test Event", resp["title"])
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/event/99", nil)
		rr := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "event")
		chiCtx.URLParams.Add("id", "99")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		mockUseCase.EXPECT().GetCultural(gomock.Any(), 99, "event").Return(culturalModel.GetCulturalResponse{}, errors.New("cultural event not found"))

		handler.HandleGetCultural(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
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
}

func TestHandleGetAllCulturais(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural", nil)
		rr := httptest.NewRecorder()

		expectedResponse := culturalModel.GetAllCulturaisResponse{
			Events: []culturalModel.Event{
				{
					ID: 1,
				},
			},
			TouristAttractions: []culturalModel.TouristAttraction{
				{
					ID: 2,
				},
			},
		}

		mockUseCase.EXPECT().GetAllCulturais(gomock.Any()).Return(expectedResponse, nil)

		handler.HandleGetAllCulturais(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp culturalModel.GetAllCulturaisResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Len(t, resp.Events, 1)
		assert.Len(t, resp.TouristAttractions, 1)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural", nil)
		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().GetAllCulturais(gomock.Any()).Return(culturalModel.GetAllCulturaisResponse{}, errors.New("db error"))

		handler.HandleGetAllCulturais(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandleGetHomeCulturais(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/home", nil)
		rr := httptest.NewRecorder()

		expectedResponse := culturalModel.GetAllCulturaisResponse{
			Events: []culturalModel.Event{{ID: 1}},
		}

		mockUseCase.EXPECT().GetHomeCulturais(gomock.Any()).Return(expectedResponse, nil)

		handler.HandleGetHomeCulturais(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("GET", "/cultural/home", nil)
		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().GetHomeCulturais(gomock.Any()).Return(culturalModel.GetAllCulturaisResponse{}, errors.New("db error"))

		handler.HandleGetHomeCulturais(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandleUpdateCultural(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		requestBody, _ := json.Marshal(culturalModel.UpdateCulturalRequest{
			ID:    1,
			Title: "Updated Event",
			Type:  "event",
		})
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("data", string(requestBody))
		writer.Close()

		req, _ := http.NewRequest("PATCH", "/cultural", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateCultural(gomock.Any(), gomock.Any()).Return(nil)

		handler.HandleUpdateCultural(rr, req)
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		requestBody, _ := json.Marshal(culturalModel.UpdateCulturalRequest{
			ID:    1,
			Title: "Updated Event",
			Type:  "event",
		})
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_ = writer.WriteField("data", string(requestBody))
		writer.Close()

		req, _ := http.NewRequest("PATCH", "/cultural", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().UpdateCultural(gomock.Any(), gomock.Any()).Return(errors.New("update error"))

		handler.HandleUpdateCultural(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestHandleDeleteCultural(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("DELETE", "/cultural/event/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "event")
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().DeleteCultural(gomock.Any(), 1, "event").Return(nil)

		handler.HandleDeleteCultural(rr, req)
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUseCase := mock.NewMockCulturalUseCase(ctrl)
		handler := handlers.NewCulturalHandler(mockUseCase)

		req, _ := http.NewRequest("DELETE", "/cultural/event/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("type", "event")
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()

		mockUseCase.EXPECT().DeleteCultural(gomock.Any(), 1, "event").Return(errors.New("delete error"))

		handler.HandleDeleteCultural(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
