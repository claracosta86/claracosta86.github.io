package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	chi "github.com/go-chi/chi/v5"

	"poc2/back/application/user"
	userModel "poc2/back/interface/model"
)

type UserHandler struct {
	userUseCase user.UseCase
}

// NewUserHandler creates a new user HTTP handler
func NewUserHandler(userUseCase user.UseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// @Accept json
// [400] Invalid data
// [405] Invalid HTTP method
// [409] Conflict - User already exists
// [500] Internal Server Error
// [201] User registered successfully
// /users/register [POST]
// HandleRegisterUser handles user registration
func (h *UserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request userModel.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err := h.userUseCase.RegisterUser(r.Context(), request)
	if err != nil {
		if strings.Contains(err.Error(), "user already exists") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		log.Printf("Error registering user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// @Accept json
// [400] Invalid data
// [401] Invalid credentials
// [403] User not authorized
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] User logged in successfully
// /users/login [POST]
// HandleUserLogin handles user authentication
func (h *UserHandler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request userModel.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	response, err := h.userUseCase.LoginUser(r.Context(), request)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "invalid password") {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
		log.Printf("Error during login: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] User data recovered successfully
// /users/{userID}/profile/ [GET]
// HandleGetUserProfile retrieves user profile information
func (h *UserHandler) HandleGetUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	response, err := h.userUseCase.GetUserProfile(r.Context(), userID)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error getting user profile: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// @Accept json
// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] User profile edited in successfully
// /users/{userID}/profile/edit [PATCH]
// HandleEditUserProfile updates user profile information
func (h *UserHandler) HandleEditUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var request userModel.UpdateUserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err = h.userUseCase.UpdateUserProfile(r.Context(), userID, request)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error updating user profile: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Accept json
// [400] Invalid data
// [401] Incorrect password
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] User password edited in successfully
// /users/{userID}/profile/change-password [PATCH]
// HandleChangeUserPassword changes user password
func (h *UserHandler) HandleChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var request userModel.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err = h.userUseCase.ChangePassword(r.Context(), userID, request)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "current password is incorrect") {
			http.Error(w, "Current password is incorrect", http.StatusUnauthorized)
			return
		}
		log.Printf("Error changing password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] User deleted successfully
// /users/{userID}/profile/delete [DELETE]
// HandleDeleteUser removes a user account
func (h *UserHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userType := chi.URLParam(r, "userType")
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.userUseCase.DeleteUser(r.Context(), userID, userType)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"status": "user deleted successfully"})
}

// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] User favorites saved successfully
// /users/{userID}/profile/favorites [PATCH]
func (h *UserHandler) HandleUpdateFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var request userModel.FavoriteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err = h.userUseCase.UpdateFavorites(r.Context(), userID, request)
	if err != nil {
		log.Printf("Error in UpdateFavorites: %v", err)
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error updating favorites: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"status": "favorites updated successfully"})
}

// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] User favorites recovered successfully
// /users/{userID}/favorites [GET]
func (h *UserHandler) HandleGetUserFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	response, err := h.userUseCase.GetUserFavorites(r.Context(), userID)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error getting user favorites: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] Organizer recovered successfully
// /users/favorites/last-seen [PATCH]
func (h *UserHandler) HandleLastSeenFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var request userModel.FavoriteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err = h.userUseCase.UpdateLastSeenFavorite(r.Context(), userID, request)
	if err != nil {
		log.Printf("Error in UpdateLastSeenFavorite: %v", err)
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error updating last seen favorite: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"status": "last seen favorite updated successfully"})
}

// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] User favorites recovered successfully
// /users/culturais/ [GET]
func (h *UserHandler) HandleGetOrganizerCulturais(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	response, err := h.userUseCase.GetOrganizerCulturais(r.Context(), userID)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error getting user favorites: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] User favorites recovered successfully
// HandleGetOrganizerInfo
func (h *UserHandler) HandleGetOrganizerInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	response, err := h.userUseCase.GetOrganizerInfo(r.Context(), userID)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error getting organizer info: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
