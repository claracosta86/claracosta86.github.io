package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/nleof/goyesql"

	"poc2/back/domain/comment"
)

var (
	//go:embed queries/comment.sql
	commentEmbed   []byte
	commentQueries goyesql.Queries
)

type commentRepository struct {
	db *sql.DB
}

func init() {
	commentQueries = goyesql.MustParseBytes(commentEmbed)
}

// NewCommentRepository creates a new MySQL comment repository
func NewCommentRepository(db *sql.DB) comment.Repository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) SaveComment(ctx context.Context, culturalID int, culturalType string, userID int, comment comment.CommentContent) error {
	fmt.Printf("Saving comment for culturalID: %d, culturalType: %s, userID: %d\n", culturalID, culturalType, userID)
	_, err := r.db.ExecContext(ctx, commentQueries["save-comment"],
		culturalID,
		culturalType,
		userID,
		comment,
	)
	if err != nil {
		return fmt.Errorf("failed to save comment: %w", err)
	}
	return nil
}

func (r *commentRepository) FindCommentsByCultural(ctx context.Context, culturalID int, culturalType string) (comment.Comments, error) {
	rows, err := r.db.QueryContext(ctx, commentQueries["fetch-comments-by-cultural"], culturalID, culturalType)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer rows.Close()

	var comments []comment.Comment
	for rows.Next() {
		var c comment.Comment
		if err := rows.Scan(&c.ID, &c.CulturalID, &c.CulturalType, &c.UserName, &c.Comment, &c.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, c)
	}

	return comments, nil
}
