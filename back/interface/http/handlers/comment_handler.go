package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"poc2/back/application/comment"
	commentModel "poc2/back/interface/model"
)

type CommentHandler struct {
	commentUseCase comment.UseCase
}

func NewCommentHandler(commentUseCase comment.UseCase) *CommentHandler {
	return &CommentHandler{
		commentUseCase: commentUseCase,
	}
}

// [400] Invalid data
// [405] Invalid HTTP method
// [500] Internal Server Error
// [201] Comment created successfully
// /comments/ [POST]
// HandleCreateComment creates a new comment entry
func (h *CommentHandler) HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request commentModel.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err := h.commentUseCase.CreateComment(r.Context(), request)
	if err != nil {
		if err.Error() == "cultural not found" {
			http.Error(w, "Cultural not found", http.StatusBadRequest)
			return
		}

		http.Error(w, "Error creating comment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Comment created successfully"})
}

// [400] Invalid data
// [404] Comment not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] Comment data recovered successfully
// /comments/{culturalType}/{culturalID} [GET]
// HandleGetComment retrieves comment information
func (h *CommentHandler) HandleGetComment(w http.ResponseWriter, r *http.Request) {
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
	comments, err := h.commentUseCase.GetComments(r.Context(), culturalID, culturalType)
	if err != nil {
		switch err.Error() {
		case "comments not found":
			http.Error(w, "Error retrieving comments: "+err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, "Error retrieving comments: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}
