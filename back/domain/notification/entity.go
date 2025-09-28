package notification

// NotificationCulturalList represents the core notification entity in the domain
type NotificationCulturalList struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	CulturalType  string `json:"culturalType"`
	CulturalID   int    `json:"culturalID"`
	Type         string `json:"type"`
}
