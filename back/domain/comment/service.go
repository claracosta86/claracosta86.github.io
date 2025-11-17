package comment

import (
	"context"
)

type Service interface {
	// CreateComment creates a new comment entry
	CreateComment(ctx context.Context, culturalID int, culturalType string, userID int, comment string) error

	// GetComments retrieves comments for a specific cultural entry
	GetComments(ctx context.Context, culturalID int, culturalType string) (Comments, error)
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

func (s *service) CreateComment(ctx context.Context, culturalID int, culturalType string, userID int, comment string) error {
	return s.repository.SaveComment(
		ctx,
		culturalID,
		culturalType,
		userID,
		comment,
	)
}

func (s *service) GetComments(ctx context.Context, culturalID int, culturalType string) (Comments, error) {
	return s.repository.FindCommentsByCultural(ctx, culturalID, culturalType)
}
