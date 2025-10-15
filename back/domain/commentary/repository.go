package commentary


import (
	"context"

)

type Repository interface {
	// SaveCommentary saves a new commentary entry
	SaveCommentary(ctx context.Context, culturalID int, culturalType string, userID int, commentary string) error

	// FindCommentariesByCultural retrieves commentaries for a specific cultural entry
	FindCommentariesByCultural(ctx context.Context, culturalID int, culturalType string) (Commentaries, error)
}