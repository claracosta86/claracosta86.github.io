package notification

// NotificationCulturalList represents the core notification entity in the domain
type NotificationCulturalList struct {
	ID           int
	Title        string
	CulturalType string
	CulturalID   int
	Type         string
}
