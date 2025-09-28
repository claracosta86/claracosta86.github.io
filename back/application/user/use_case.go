package user

import (
	"context"

	"poc2/back/domain/user"
	"poc2/back/domain/cultural"
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
	DeleteUser(ctx context.Context, userID int, userType string) error

	// ToggleFavorite adds or removes a cultural item from user's favorites
	ToggleFavorite(ctx context.Context, userID int, request model.FavoriteRequest) error

	// GetUserFavorites retrieves a user's favorite cultural items
	GetUserFavorites(ctx context.Context, userID int) ([]model.FavoriteCulturalList, error)
}

type useCase struct {
	userService    user.Service
	culturalService cultural.Service
}

// NewUseCase creates a new user use case
func NewUseCase(userService user.Service, culturalService cultural.Service) UseCase {
	return &useCase{
		userService:    userService,
		culturalService: culturalService,
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
		CompanyName: user.CompanyName,
		Type:        user.Type,
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
func (uc *useCase) DeleteUser(ctx context.Context, userID int, userType string) error {
	if userType == "organizer" {
		eventsIDs, err := uc.culturalService.GetEventsIDsByOrganizer(ctx, userID); 
		if err != nil {
			return err
		}

		attractionsIDs, err := uc.culturalService.GetTouristAttractionsIDsByOrganizer(ctx, userID);
		if err != nil {
			return err
		}

		// Delete events
		if len(eventsIDs) > 0 {
			if err := uc.userService.RemoveEventFromAllUsers(ctx, eventsIDs); err != nil {
				return err
			}
		}

		// Delete attractions
		if len(attractionsIDs) > 0 {
			if err := uc.userService.RemoveTouristAttractionFromAllUsers(ctx, attractionsIDs); err != nil {
				return err
			}
		}
	}

	return uc.userService.DeleteUser(ctx, userID)
}

// ToggleFavorite adds or removes a cultural item from user's favorites
func (uc *useCase) ToggleFavorite(ctx context.Context, userID int, request model.FavoriteRequest) error {
	return uc.userService.ToggleFavorite(ctx, userID, request.CulturalType, request.CulturalID, request.IsFavorite)
}

// GetUserFavorites retrieves a user's favorite cultural items
func (uc *useCase) GetUserFavorites(ctx context.Context, userID int) ([]model.FavoriteCulturalList, error) {
	favorites, err := uc.userService.GetUserFavorites(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	result := make([]model.FavoriteCulturalList, len(favorites))
	for _, favorite := range favorites {
		result = append(result, model.FavoriteCulturalList{
			ID:   favorite.ID,
			Type: favorite.Type,
		})
	}
	
	return result, nil
}