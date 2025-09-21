package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

	chi "github.com/go-chi/chi/v5"

	"poc2/back/application/cultural"
	// culturalModel "poc2/back/interface/model"

)

type CulturalHandler struct {
	culturalUseCase cultural.UseCase
}

func NewCulturalHandler(culturalUseCase cultural.UseCase) *CulturalHandler {
	return &CulturalHandler{
		culturalUseCase: culturalUseCase,
	}
}

func (h *CulturalHandler) HandleCreateCultural(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating cultural content
}

// [400] Invalid data
// [404] Cultural not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] Cultural data recovered successfully
// /cultural/{type}/{id} [GET]
// HandleGetCulturalEvent retrieves cultural information
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
		case "cultural attraction not found", "cultural event not found" :
			http.Error(w, "Error retrieving cultural data: "+err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, "Error retrieving cultural data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Println("Cultural data retrieved:", culturalData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(culturalData)
}

func (h *CulturalHandler) HandleUpdateCultural(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating cultural content
}

func (h *CulturalHandler) HandleDeleteCultural(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting cultural content
}