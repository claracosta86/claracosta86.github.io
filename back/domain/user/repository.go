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
	
	// Delete removes a user from the repository
	Delete(ctx context.Context, id int) error
	
	// CheckPassword verifies if a user's password matches
	CheckPassword(ctx context.Context, userID int, password string) (bool, error)
	
	// UpdatePassword updates a user's password
	UpdatePassword(ctx context.Context, userID int, newPassword string) error
}
