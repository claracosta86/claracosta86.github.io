package notification

import "context"

// Repository defines the contract for user data persistence
type Repository interface {
	// Search notifications by user ID
	FindByUserID(ctx context.Context, userID int) ([]NotificationCulturalList, error)
}
