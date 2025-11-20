package comment

import (
	"context"
)

type Repository interface {
	// SaveComment saves a new comment entry
	SaveComment(ctx context.Context, culturalID int, culturalType string, userID int, comment string) error

	// FindCommentsByCultural retrieves comments for a specific cultural entry
	FindCommentsByCultural(ctx context.Context, culturalID int, culturalType string) (Comments, error)
}
