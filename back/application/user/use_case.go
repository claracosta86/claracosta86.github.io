package user

import (
	"context"

	"poc2/back/domain/user"
	"poc2/back/interface/model"

)

// UseCase defines the application use cases for user operations
type UseCase interface {
	// RegisterUser handles user registration
	RegisterUser(ctx context.Context, request model.RegisterUserRequest) error

	// LoginUser handles user authentication
	LoginUser(ctx context.Context, request model.LoginUserRequest) (*model.LoginUserResponse, error)

	// GetUserProfile retrieves user profile information
	GetUserProfile(ctx context.Context, userID int) (*model.GetUserProfileResponse, error)

	// UpdateUserProfile updates user profile information
	UpdateUserProfile(ctx context.Context, userID int, request model.UpdateUserProfileRequest) error

	// ChangePassword changes user password
	ChangePassword(ctx context.Context, userID int, request model.ChangePasswordRequest) error

	// DeleteUser removes a user account
	DeleteUser(ctx context.Context, userID int) error
}

type useCase struct {
	userService user.Service
}

// NewUseCase creates a new user use case
func NewUseCase(userService user.Service) UseCase {
	return &useCase{
		userService: userService,
	}
}

// RegisterUser handles user registration
func (uc *useCase) RegisterUser(ctx context.Context, request model.RegisterUserRequest) error {
	userType, err := user.UserType(request.Type).Validate()
	if err != nil {
		return err
	}
	
	return uc.userService.RegisterUser(
		ctx,
		request.Name,
		request.Email,
		request.Document,
		request.CompanyName,
		request.Password,
		userType,
	)
}

// LoginUser handles user authentication
func (uc *useCase) LoginUser(ctx context.Context, request model.LoginUserRequest) (*model.LoginUserResponse, error) {
	user, err := uc.userService.AuthenticateUser(ctx, request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	
	return &model.LoginUserResponse{
		UserID: user.ID,
		Name:   user.Name,
		Type:   user.Type,
	}, nil
}

// GetUserProfile retrieves user profile information
func (uc *useCase) GetUserProfile(ctx context.Context, userID int) (*model.GetUserProfileResponse, error) {
	user, err := uc.userService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	return &model.GetUserProfileResponse{
		UserID:      user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Type:        user.Type,
		CompanyName: user.CompanyName,
	}, nil
}

// UpdateUserProfile updates user profile information
func (uc *useCase) UpdateUserProfile(ctx context.Context, userID int, request model.UpdateUserProfileRequest) error {
	return uc.userService.UpdateUserProfile(
		ctx,
		userID,
		request.Name,
		request.Email,
		request.CompanyName,
	)
}

// ChangePassword changes user password
func (uc *useCase) ChangePassword(ctx context.Context, userID int, request model.ChangePasswordRequest) error {
	return uc.userService.ChangeUserPassword(
		ctx,
		userID,
		request.CurrentPassword,
		request.NewPassword,
	)
}

// DeleteUser removes a user account
func (uc *useCase) DeleteUser(ctx context.Context, userID int) error {
	return uc.userService.DeleteUser(ctx, userID)
}
