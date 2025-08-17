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
	
	// DeleteUser removes a user from the system
	DeleteUser(ctx context.Context, id int) error
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
	// Create a new user with validation
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
	fmt.Println(email, password)
	if err != nil {
			fmt.Println("hew")

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
	
	// Update profile
	if err := user.UpdateProfile(name, email, companyName); err != nil {
		return err
	}
	
	// Save changes
	return s.repository.Update(ctx, user)
}

func (s *service) ChangeUserPassword(ctx context.Context, id int, currentPassword, newPassword string) error {
	// Get existing user
	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}
	
	// Change password
	if err := user.ChangePassword(currentPassword, newPassword); err != nil {
		return err
	}
	
	// Save changes
	return s.repository.UpdatePassword(ctx, id, newPassword)
}

func (s *service) DeleteUser(ctx context.Context, id int) error {
	// Check if user exists
	_, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}
	
	// Delete user
	return s.repository.Delete(ctx, id)
}
