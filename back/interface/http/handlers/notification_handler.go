package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
"fmt"

	chi "github.com/go-chi/chi/v5"

	"poc2/back/application/notification"

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
		fmt.Println("Error retrieving notifications:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(notifications)
}