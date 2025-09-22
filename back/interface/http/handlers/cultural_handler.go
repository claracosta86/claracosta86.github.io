package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

	chi "github.com/go-chi/chi/v5"

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

	err := r.ParseMultipartForm(10 << 20) // Limite de 10MB para o formulário
	if err != nil {
		http.Error(w, "Error parsing form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	data := r.FormValue("data")
	if data == "" {
		http.Error(w, "Missing data field", http.StatusBadRequest)
		return
	}
	var createReq culturalModel.CreateCulturalRequest
	err = json.Unmarshal([]byte(data), &createReq)
	if err != nil {
		http.Error(w, "Error parsing JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		// Processar o arquivo da imagem conforme necessário
		// Por exemplo, salvar em um serviço de armazenamento ou banco de dados
		// Aqui, apenas simulamos que a imagem foi processada
		createReq.Image = "processed_image_path_or_url"
	}

	err = h.culturalUseCase.CreateCultural(r.Context(), createReq)
	if err != nil {
		http.Error(w, "Error creating cultural: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cultural created successfully"))
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