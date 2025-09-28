package model

type GetNotificationsResponse struct {
	UserID        int       `json:"userID"`
	Cultural  []CulturalList    `json:"culturals"`
}

type CulturalList struct {
	ID    int    `json:"id"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	NotificationType string `json:"notificationType"`
	NotificationID   int    `json:"notificationID"`
}

type SeenNortificationsRequest struct {
	NotificationIDs []int `json:"notificationIDs"`
}