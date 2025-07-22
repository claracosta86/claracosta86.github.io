package api

import (
	"encoding/json"
	"log"
	"net/http"

	"poc2/back/model"
	"poc2/back/service"

)

type AttractionHandler struct {
	attractionService service.AttractionService
}

func AttractionHandler() *AttractionHandler {
	return &AttractionHandler{
		attractionService: service.NewAttractionService(),
	}
}

func (h *AttractionHandler) HandleGetUserAttractions(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	attractions, err := h.attractionService.GetUserAttractions(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(attractions)
}

func (h *AttractionHandler) HandleRegisterAttraction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var attraction model.Attraction
	err := json.NewDecoder(r.Body).Decode(&attraction)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = h.attractionService.Registerattraction(attraction)
	if err != nil {
		http.Error(w, "Erro ao salvar attractiono", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sucesso"}`))
}

func (h *AttractionHandler) HandleGetAllAttractions(w http.ResponseWriter, r *http.Request) {
	attractions, err := h.attractionService.GetAllAttractions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(attractions)
}

func (h *AttractionHandler) HandleGetAttractionByID(w http.ResponseWriter, r *http.Request) {
	attractionID := r.Context().Value("attractionID").(string)
	attraction, err := h.attractionService.GetAttractionByID(attractionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(attraction)
}

func (h *AttractionHandler) HandleUpdateAttraction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var attraction model.Attraction
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sucesso"}`))
}