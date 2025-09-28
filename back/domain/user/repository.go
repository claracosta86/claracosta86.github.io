package user

import "context"

// Repository defines the contract for user data persistence
type Repository interface {
	// Save saves a new user to the repository
	Save(ctx context.Context, user *User) error
	
	// FindByID finds a user by their ID
	FindByID(ctx context.Context, id int) (*User, error)
	
	// FindByEmail finds a user by their email address
	FindByEmail(ctx context.Context, email string) (*User, error)
	
	// Update updates an existing user
	Update(ctx context.Context, user *User) error

	// DeleteUserFavorites removes all favorites of a user
	DeleteUserFavorites(ctx context.Context, userID int) error

	// DeleteUser removes a user from the repository
	DeleteUser(ctx context.Context, id int) error
	
	// CheckPassword verifies if a user's password matches
	CheckPassword(ctx context.Context, userID int, password string) (bool, error)
	
	// UpdatePassword updates a user's password
	UpdatePassword(ctx context.Context, userID int, newPassword string) error

	// RemoveEvent removes an event from all users' favorites
	RemoveEvent(ctx context.Context, eventIDs []int) error

	// RemoveTouristAttraction removes a tourist attraction from all users' favorites
	RemoveTouristAttraction(ctx context.Context, touristAttractionIDs []int) error

	// AddFavorite adds a cultural item to user's favorites
	AddFavorite(ctx context.Context, userID int, culturalType string, culturalID int) error

	// RemoveFavorite removes a cultural item from user's favorites
	RemoveFavorite(ctx context.Context, userID int, culturalType string, culturalID int) error

	// GetFavoritesByUserID retrieves all favorite cultural items of a user
	GetFavoritesByUserID(ctx context.Context, userID int) ([]FavoriteCulturalList, error)
}
