package commentary

import (
	"context"

)

type Service interface {
	// CreateCommentary creates a new commentary entry
	CreateCommentary(ctx context.Context, culturalID int, culturalType string, userID int, commentary string) error
	
	// GetCommentaries retrieves commentaries for a specific cultural entry
	GetCommentaries(ctx context.Context, culturalID int, culturalType string) (Commentaries, error)
}


type service struct {
	repository Repository
}

// NewService creates a new cultural service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateCommentary(ctx context.Context, culturalID int, culturalType string, userID int, commentary string) error {
	return s.repository.SaveCommentary(
		ctx,
		culturalID,
		culturalType,
		userID,
		commentary,
	)
}

func (s *service) GetCommentaries(ctx context.Context, culturalID int, culturalType string) (Commentaries, error) {
	return s.repository.FindCommentariesByCultural(ctx, culturalID, culturalType)
}