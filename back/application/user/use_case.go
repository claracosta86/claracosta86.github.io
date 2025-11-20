package user

import (
	"context"
	"fmt"
	"time"

	"poc2/back/domain/cultural"
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
	DeleteUser(ctx context.Context, userID int, userType string) error

	// ToggleFavorite adds or removes a cultural item from user's favorites
	ToggleFavorite(ctx context.Context, userID int, request model.FavoriteRequest) error

	// GetUserFavorites retrieves a user's favorite cultural items
	GetUserFavorites(ctx context.Context, userID int) ([]model.CulturalList, error)

	// UpdateLastSeenFavorite updates the last seen timestamp of a favorite cultural item
	UpdateLastSeenFavorite(ctx context.Context, userID int, request model.FavoriteRequest) error

	// GetOrganizerCulturais retrieves cultural items associated with an organizer
	GetOrganizerCulturais(ctx context.Context, organizerID int) ([]model.CulturalList, error)

	// GetOrganizerInfo retrieves organizer information
	GetOrganizerInfo(ctx context.Context, organizerID int) (*model.GetOrganizerInfoResponse, error)
}

type useCase struct {
	userService     user.Service
	culturalService cultural.Service
}

// NewUseCase creates a new user use case
func NewUseCase(userService user.Service, culturalService cultural.Service) UseCase {
	return &useCase{
		userService:     userService,
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
		eventsIDs, err := uc.culturalService.GetEventsIDsByOrganizer(ctx, userID)
		if err != nil {
			return err
		}

		attractionsIDs, err := uc.culturalService.GetTouristAttractionsIDsByOrganizer(ctx, userID)
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
func (uc *useCase) GetUserFavorites(ctx context.Context, userID int) ([]model.CulturalList, error) {
	favorites, err := uc.userService.GetUserFavorites(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]model.CulturalList, 0)
	for _, favorite := range favorites {
		result = append(result, model.CulturalList{
			ID:   favorite.ID,
			Type: favorite.Type,
		})
	}

	return result, nil
}

// UpdateLastSeenFavorite updates the last seen timestamp of a favorite cultural item
func (uc *useCase) UpdateLastSeenFavorite(ctx context.Context, userID int, request model.FavoriteRequest) error {
	return uc.userService.UpdateLastSeenFavorite(ctx, userID, request.CulturalID, request.CulturalType)
}

// GetOrganizerCulturais retrieves cultural items associated with an organizer
func (uc *useCase) GetOrganizerCulturais(ctx context.Context, organizerID int) ([]model.CulturalList, error) {
	culturais, err := uc.userService.GetCulturaisByOrganizerID(ctx, organizerID)
	if err != nil {
		return nil, err
	}

	result := make([]model.CulturalList, 0)
	for _, cultural := range culturais {
		result = append(result, model.CulturalList{
			ID:   cultural.ID,
			Type: cultural.Type,
		})
	}

	return result, nil
}

// GetOrganizerInfo retrieves organizer information
func (uc *useCase) GetOrganizerInfo(ctx context.Context, organizerID int) (*model.GetOrganizerInfoResponse, error) {
	organizer, err := uc.userService.GetUserByID(ctx, organizerID)
	if err != nil {
		return nil, err
	}

	culturalItems, err := uc.GetOrganizerCulturais(ctx, organizerID)
	if err != nil {
		return nil, err
	}

	years, err := time.Parse("2006-01-02 15:04:05", organizer.CreatedAt)
	if err != nil {
		return nil, err
	}
	yearsSince := formatTimeSince(years)

	return &model.GetOrganizerInfoResponse{
		Name:           organizer.Name,
		Email:          organizer.Email,
		ID:             organizer.ID,
		CulturalItems:  culturalItems,
		OrganizerSince: yearsSince,
	}, nil
}

func formatTimeSince(t time.Time) string {
	now := time.Now()

	years := now.Year() - t.Year()
	months := int(now.Month() - t.Month())
	days := now.Day() - t.Day()

	if days < 0 {
		months--
		days += time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, time.UTC).Day()
	}
	if months < 0 {
		years--
		months += 12
	}

	if years > 0 {
		if years == 1 {
			return "1 ano"
		}
		return fmt.Sprintf("%d anos", years)
	}

	if months > 0 {
		if months == 1 {
			return "1 mês"
		}
		return fmt.Sprintf("%d meses", months)
	}

	if days == 1 {
		return "1 dia"
	}
	return fmt.Sprintf("%d dias", days)
}
