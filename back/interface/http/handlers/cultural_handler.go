package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	chi "github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"poc2/back/application/cultural"
	culturalModel "poc2/back/interface/model"
)

type CulturalHandler struct {
	culturalUseCase cultural.UseCase
}

func NewCulturalHandler(culturalUseCase cultural.UseCase) *CulturalHandler {
	return &CulturalHandler{
		culturalUseCase: culturalUseCase,
	}
}

// [400] Invalid data
// [405] Invalid HTTP method
// [500] Internal Server Error
// [201] Cultural created successfully
// /cultural [POST]
// HandleCreateCultural creates a new cultural entry
func (h *CulturalHandler) HandleCreateCultural(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	createReq, err := parseCreateCulturalRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	imageName, err := processImageUpload(r, "./static/culturalthumbs")
	if err != nil && err != http.ErrMissingFile {
		fmt.Printf("Error processing image upload: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createReq.Image = imageName

	result, err := h.culturalUseCase.CreateCultural(r.Context(), createReq)
	if err != nil {
		fmt.Printf("Error creating cultural: %v\n", err)
		http.Error(w, "Error creating cultural: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Cultural created successfully", "id": strconv.Itoa(result.ID), "type": result.Type})
}

// HandleGetAllCulturais retrieves all cultural events and attractions
func (h *CulturalHandler) HandleGetAllCulturais(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	culturals, err := h.culturalUseCase.GetAllCulturais(r.Context())
	if err != nil {
		http.Error(w, "Error retrieving culturals: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(culturals)
}

// HandleGetAllCulturais retrieves all cultural events and attractions
func (h *CulturalHandler) HandleGetHomeCulturais(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	culturals, err := h.culturalUseCase.GetHomeCulturais(r.Context())
	if err != nil {
		http.Error(w, "Error retrieving culturals: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(culturals)
}

// [400] Invalid data
// [404] Cultural not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] Cultural data recovered successfully
// /cultural/{type}/{id} [GET]
// HandleGetCultural retrieves cultural information
func (h *CulturalHandler) HandleGetCultural(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	culturalID := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(culturalID)
	if err != nil {
		http.Error(w, "Invalid cultural ID", http.StatusBadRequest)
		return
	}

	culturalType := chi.URLParam(r, "type")
	if culturalType != cultural.CulturalTypeEvent && culturalType != cultural.CulturalTypeTouristAttraction {
		http.Error(w, "Invalid cultural type", http.StatusBadRequest)
		return
	}

	culturalData, err := h.culturalUseCase.GetCultural(r.Context(), ID, culturalType)
	if err != nil {
		switch err.Error() {
		case "cultural attraction not found", "cultural event not found":
			http.Error(w, "Error retrieving cultural data: "+err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, "Error retrieving cultural data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(culturalData)
}

func (h *CulturalHandler) HandleUpdateCultural(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	updateReq, err := parseUpdateCulturalRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	imageName, err := processImageUpload(r, "./static/culturalthumbs")
	if err != nil && err != http.ErrMissingFile {
		fmt.Printf("Error processing image upload: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updateReq.Image = imageName

	err = h.culturalUseCase.UpdateCultural(r.Context(), updateReq)
	if err != nil {
		fmt.Printf("Error updating cultural: %v\n", err)
		http.Error(w, "Error updating cultural: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Cultural created successfully", "id": strconv.Itoa(updateReq.ID), "type": updateReq.Type})
}

func (h *CulturalHandler) HandleDeleteCultural(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	culturalID := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(culturalID)
	if err != nil {
		http.Error(w, "Invalid cultural ID", http.StatusBadRequest)
		return
	}

	culturalType := chi.URLParam(r, "type")
	if culturalType != cultural.CulturalTypeEvent && culturalType != cultural.CulturalTypeTouristAttraction {
		http.Error(w, "Invalid cultural type", http.StatusBadRequest)
		return
	}

	err = h.culturalUseCase.DeleteCultural(r.Context(), ID, culturalType)
	if err != nil {
		fmt.Printf("Error deleting cultural: %v\n", err)
		http.Error(w, "Error deleting cultural: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseCreateCulturalRequest(r *http.Request) (culturalModel.CreateCulturalRequest, error) {
	var req culturalModel.CreateCulturalRequest

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return req, fmt.Errorf("formato do request inválido: %w", err)
	}

	dataJSON := r.FormValue("data")
	if dataJSON == "" {
		return req, errors.New("campo 'data' do formulário não encontrado")
	}

	if err := json.Unmarshal([]byte(dataJSON), &req); err != nil {
		return req, fmt.Errorf("erro ao decodificar JSON do campo 'data': %w", err)
	}

	return req, nil
}

func parseUpdateCulturalRequest(r *http.Request) (culturalModel.UpdateCulturalRequest, error) {
	var req culturalModel.UpdateCulturalRequest

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return req, fmt.Errorf("formato do request inválido: %w", err)
	}

	dataJSON := r.FormValue("data")
	if dataJSON == "" {
		return req, errors.New("campo 'data' do formulário não encontrado")
	}

	if err := json.Unmarshal([]byte(dataJSON), &req); err != nil {
		return req, fmt.Errorf("erro ao decodificar JSON do campo 'data': %w", err)
	}

	return req, nil
}

func processImageUpload(r *http.Request, destinationPath string) (string, error) {
	file, header, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return "", http.ErrMissingFile
		}
		return "", fmt.Errorf("erro ao extrair arquivo: %w", err)
	}
	defer file.Close()

	fileName := uuid.New().String() + filepath.Ext(header.Filename)
	fullPath := filepath.Join(destinationPath, fileName)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("não foi possível criar arquivo no servidor: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("não foi possível salvar o arquivo: %w", err)
	}

	return fileName, nil
}
