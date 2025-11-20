package notification

import "context"

// Repository defines the contract for user data persistence
type Repository interface {
	// Search notifications by user ID
	FindByUserID(ctx context.Context, userID int) ([]NotificationCulturalList, error)

	// Mark notifications as seen by user ID and notification IDs
	MarkAsSeen(ctx context.Context, userID int, notificationIDs []int) error
}
