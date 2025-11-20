package model

type GetNotificationsResponse struct {
	UserID   int                        `json:"userID"`
	Cultural []NotificationCulturalList `json:"culturals"`
}

type NotificationCulturalList struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	CulturalType     string `json:"culturalType"`
	NotificationType string `json:"notificationType"`
	NotificationID   int    `json:"notificationID"`
}

type SeenNotificationPost struct {
	NotificationIDs []int `json:"notificationIDs"`
}
