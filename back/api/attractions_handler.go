package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"poc2/back/model"
	"poc2/back/service"
	"poc2/back/lib/logging"

)

type AttractionHandler struct {
	attractionService service.AttractionService
}

func NewAttractionHandler() *AttractionHandler {
	return &AttractionHandler{
		attractionService: service.NewAttractionService(),
	}
}

// @ Register a new tourist attraction handler
// @ Accept json
// /attractions/register [POST]
func (h *AttractionHandler) HandleRegisterAttraction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var attraction model.TouristAttraction
	err := json.NewDecoder(r.Body).Decode(&attraction)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = h.attractionService.RegisterAttraction(attraction)
	if err != nil {
		http.Error(w, "Erro ao salvar attractiono", http.StatusInternalServerError)
		return
	}

	rec := &logging.StatusRecorder{ResponseWriter: w, Status: http.StatusNoContent}
	rec.WriteHeader(http.StatusNoContent)
}

// @ Gets all tourist attractions available
// /attractions/all [GET]
func (h *AttractionHandler) HandleGetAllAttractions(w http.ResponseWriter, r *http.Request) {
	attractions, err := h.attractionService.GetAllAttractions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(attractions)

	rec := &logging.StatusRecorder{ResponseWriter: w, Status: http.StatusOK}
	rec.WriteHeader(http.StatusOK)
}

// @ Gets a tourist attraction information by its ID
// /attractions/:id [GET]
func (h *AttractionHandler) HandleGetAttractionByID(w http.ResponseWriter, r *http.Request) {
	attractionID := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(attractionID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	attraction, err := h.attractionService.GetAttractionByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(attraction)

	rec := &logging.StatusRecorder{ResponseWriter: w, Status: http.StatusOK}
	rec.WriteHeader(http.StatusOK)
}

// @ Updates a tourist attraction information
// @ Accept json
// /attractions/update [PUT]
func (h *AttractionHandler) HandleUpdateAttraction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var attraction model.TouristAttraction
	err := json.NewDecoder(r.Body).Decode(&attraction)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = h.attractionService.UpdateAttraction(attraction)
	if err != nil {
		http.Error(w, "Erro ao atualizar attractiono", http.StatusInternalServerError)
		return
	}

	rec := &logging.StatusRecorder{ResponseWriter: w, Status: http.StatusNoContent}
	rec.WriteHeader(http.StatusNoContent)
}

// @ Deletes a tourist attraction by its ID
// attractions/delete/:id [DELETE]
func (h *AttractionHandler) HandleDeleteAttraction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	attractionID := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(attractionID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.attractionService.DeleteAttractionByID(ID)
	if err != nil {
		http.Error(w, "Erro ao deletar attractiono", http.StatusInternalServerError)
		return
	}

	rec := &logging.StatusRecorder{ResponseWriter: w, Status: http.StatusNoContent}
	rec.WriteHeader(http.StatusNoContent)
}