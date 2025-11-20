package user

import (
	"context"
	"errors"
	"fmt"
)

// Service defines the business logic for user operations
type Service interface {
	// RegisterUser registers a new user in the system
	RegisterUser(ctx context.Context, name, email, document, companyName, password string, userType UserType) error

	// AuthenticateUser authenticates a user with email and password
	AuthenticateUser(ctx context.Context, email, password string) (*User, error)

	// GetUserByID retrieves a user by their ID
	GetUserByID(ctx context.Context, id int) (*User, error)

	// UpdateUserProfile updates a user's profile information
	UpdateUserProfile(ctx context.Context, id int, name, email, companyName string) error

	// ChangeUserPassword changes a user's password
	ChangeUserPassword(ctx context.Context, id int, currentPassword, newPassword string) error

	// RemoveEventFromAllUsers removes an event from all users' favorites
	RemoveEventFromAllUsers(ctx context.Context, eventIDs []int) error

	// RemoveTouristAttractionFromAllUsers removes a tourist attraction from all users' favorites
	RemoveTouristAttractionFromAllUsers(ctx context.Context, attractionIDs []int) error

	// DeleteUser removes a user from the system
	DeleteUser(ctx context.Context, id int) error

	// ToggleFavorite adds or removes a cultural item from user's favorites
	ToggleFavorite(ctx context.Context, userID int, culturalType string, culturalID int, isFavorite bool) error

	// GetUserFavorites retrieves a user's favorite cultural items
	GetUserFavorites(ctx context.Context, userID int) ([]CulturalList, error)

	// UpdateLastSeenFavorite updates the last seen timestamp of a favorite cultural item
	UpdateLastSeenFavorite(ctx context.Context, userID, culturalID int, culturalType string) error

	// GetCulturaisByOrganizerID retrieves cultural items associated with an organizer
	GetCulturaisByOrganizerID(ctx context.Context, organizerID int) ([]CulturalList, error)
}

type service struct {
	repository Repository
}

// NewService creates a new user service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterUser(ctx context.Context, name, email, document, companyName, password string, userType UserType) error {
	// Create a new user
	user, err := NewUser(name, email, document, companyName, password, userType)
	if err != nil {
		return err
	}

	// Check if user already exists
	existingUser, err := s.repository.FindByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists with this email")
	}

	// Save the user
	return s.repository.Save(ctx, user)
}

func (s *service) AuthenticateUser(ctx context.Context, email, password string) (*User, error) {
	// Find user by email
	user, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Verify password
	valid, err := s.repository.CheckPassword(ctx, user.ID, password)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, id int) (*User, error) {
	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *service) UpdateUserProfile(ctx context.Context, id int, name, email, companyName string) error {
	// Get existing user
	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	user.Name = name
	user.Email = email
	user.CompanyName = companyName

	// Save changes
	return s.repository.Update(ctx, user)
}

func (s *service) ChangeUserPassword(ctx context.Context, id int, currentPassword, newPassword string) error {
	// Change password
	if valid, err := s.repository.CheckPassword(ctx, id, currentPassword); err != nil {
		return err
	} else if !valid {
		return errors.New("current password is incorrect")
	}

	// Save changes
	return s.repository.UpdatePassword(ctx, id, newPassword)
}

func (s *service) RemoveEventFromAllUsers(ctx context.Context, eventIDs []int) error {
	return s.repository.RemoveEvent(ctx, eventIDs)
}

func (s *service) RemoveTouristAttractionFromAllUsers(ctx context.Context, attractionIDs []int) error {
	return s.repository.RemoveTouristAttraction(ctx, attractionIDs)
}

func (s *service) DeleteUser(ctx context.Context, id int) error {
	// Check if user exists
	_, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	// Delete user favorites
	if err := s.repository.DeleteUserFavorites(ctx, id); err != nil {
		return err
	}

	// Delete user
	return s.repository.DeleteUser(ctx, id)
}

func (s *service) ToggleFavorite(ctx context.Context, userID int, culturalType string, culturalID int, isFavorite bool) error {
	// Check if user exists
	_, err := s.repository.FindByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	if isFavorite {
		return s.repository.AddFavorite(ctx, userID, culturalType, culturalID)
	} else {
		return s.repository.RemoveFavorite(ctx, userID, culturalType, culturalID)
	}
}

func (s *service) GetUserFavorites(ctx context.Context, userID int) ([]CulturalList, error) {
	_, err := s.repository.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return s.repository.GetFavoritesByUserID(ctx, userID)
}

func (s *service) UpdateLastSeenFavorite(ctx context.Context, userID, culturalID int, culturalType string) error {
	_, err := s.repository.FindByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repository.UpdateLastSeenFavorite(ctx, userID, culturalID, culturalType)
}

func (s *service) GetCulturaisByOrganizerID(ctx context.Context, organizerID int) ([]CulturalList, error) {
	// Check if organizer exists
	organizer, err := s.repository.FindByID(ctx, organizerID)
	if err != nil {
		return nil, errors.New("organizer not found")
	}
	if organizer.Type != string(UserTypeOrganizer) {
		return nil, fmt.Errorf("user with ID %d is not an organizer", organizerID)
	}

	return s.repository.GetCulturaisByOrganizerID(ctx, organizerID)
}
