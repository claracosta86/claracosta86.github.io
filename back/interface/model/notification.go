package model

type GetNotificationsResponse struct {
	UserID   int                        `json:"userID"`
	Cultural []NotificationCulturalList `json:"culturals"`
}

type NotificationCulturalList struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Type             string `json:"type"`
	NotificationType string `json:"notificationType"`
	NotificationID   int    `json:"notificationID"`
}

type SeenNotificationPost struct {
	NotificationIDs []int `json:"notificationIDs"`
}
