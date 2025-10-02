package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"

	chi "github.com/go-chi/chi/v5"

	"poc2/back/application/notification"
	notificationModel "poc2/back/interface/model"

)

type NotificationHandler struct {
	notificationUseCase notification.UseCase
}

func NewNotificationHandler(notificationUseCase notification.UseCase) *NotificationHandler {
	return &NotificationHandler{
		notificationUseCase: notificationUseCase,
	}
}

// [400] User ID invalid
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] Notifications retrieved successfully
// /notifications/{userID} [POST]
// HandleGetUserNotifications handles new notifications retrieval
func (h *NotificationHandler) HandleGetUserNotifications(w http.ResponseWriter, r *http.Request) {
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
	notifications, err := h.notificationUseCase.GetNotifications(r.Context(), userID)
	if err != nil {
		log.Printf("Error retrieving notifications: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(notifications)
}

// [400] User ID invalid
// [400] Invalid request body
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] Notifications marked as seen successfully
// /notifications/{userID}/seen [PATCH]
// HandleMarkNotificationsAsSeen handles marking notifications as seen
func (h *NotificationHandler) HandleMarkNotificationsAsSeen(w http.ResponseWriter, r *http.Request) {	
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

	var seenReq notificationModel.SeenNortificationsRequest
	err = json.NewDecoder(r.Body).Decode(&seenReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.notificationUseCase.MarkNotificationsAsSeen(r.Context(), userID, seenReq.NotificationIDs)
	if err != nil {
		log.Printf("Error marking notifications as seen: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Notifications marked as seen"))
}