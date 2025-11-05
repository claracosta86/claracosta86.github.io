package handlers


import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"github.com/go-chi/chi/v5"

	"poc2/back/application/commentary"
	commentaryModel "poc2/back/interface/model"

)

type CommentaryHandler struct {
	commentaryUseCase commentary.UseCase
}

func NewCommentaryHandler(commentaryUseCase commentary.UseCase) *CommentaryHandler {
	return &CommentaryHandler{
		commentaryUseCase: commentaryUseCase,
	}
}

// [400] Invalid data
// [405] Invalid HTTP method
// [500] Internal Server Error
// [201] Commentary created successfully
// /commentarys/ [POST]
// HandleCreateCommentary creates a new commentary entry
func (h *CommentaryHandler) HandleCreateCommentary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request commentaryModel.CreateCommentaryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
fmt.Printf("Received commentary creation request: %+v\n", request)
    err := h.commentaryUseCase.CreateCommentary(r.Context(), request)
    if err != nil {
		if err.Error() == "cultural not found" {
			http.Error(w, "Cultural not found", http.StatusBadRequest)
			return
		}
		fmt.Printf("Error creating commentary: %v\n", err)
        http.Error(w, "Error creating commentary: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Commentary created successfully"})
}

// [400] Invalid data
// [404] Commentary not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] Commentary data recovered successfully
// /commentarys/{culturalType}/{culturalID} [GET]
// HandleGetCommentary retrieves commentary information
func (h *CommentaryHandler) HandleGetCommentary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}


	culturalIDStr := chi.URLParam(r, "culturalID")
	culturalID, err := strconv.Atoi(culturalIDStr)
	if err != nil {
		http.Error(w, "Invalid cultural ID", http.StatusBadRequest)
		return
	}

	culturalType := chi.URLParam(r, "culturalType")
	if culturalType != commentary.CulturalTypeEvent && culturalType != commentary.CulturalTypeTouristAttraction {
		http.Error(w, "Invalid cultural type", http.StatusBadRequest)
		return
	}

	commentaries, err := h.commentaryUseCase.GetCommentaries(r.Context(), culturalID, culturalType)
	if err != nil {
		switch err.Error() {
		case "commentaries not found":
			http.Error(w, "Error retrieving commentaries: "+err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, "Error retrieving commentaries: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commentaries)
}

func (h *CommentaryHandler) HandleUpdateCommentary(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating commentary content
}

func (h *CommentaryHandler) HandleDeleteCommentary(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting commentary content
}