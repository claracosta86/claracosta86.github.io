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
	commentaryEmbed   []byte
	commentaryQueries goyesql.Queries
)

type commentaryRepository struct {
	db *sql.DB
}

func init() {
	commentaryQueries = goyesql.MustParseBytes(commentaryEmbed)
}

// NewCommentRepository creates a new MySQL comment repository
func NewCommentRepository(db *sql.DB) comment.Repository {
	return &commentaryRepository{
		db: db,
	}
}

func (r *commentaryRepository) SaveComment(ctx context.Context, culturalID int, culturalType string, userID int, comment string) error {
	fmt.Printf("Saving comment for culturalID: %d, culturalType: %s, userID: %d\n", culturalID, culturalType, userID)
	_, err := r.db.ExecContext(ctx, commentaryQueries["save-comment"],
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

func (r *commentaryRepository) FindCommentsByCultural(ctx context.Context, culturalID int, culturalType string) (comment.Comments, error) {
	rows, err := r.db.QueryContext(ctx, commentaryQueries["fetch-commentaries-by-cultural"], culturalID, culturalType)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commentaries: %w", err)
	}
	defer rows.Close()

	var commentaries []comment.Comment
	for rows.Next() {
		var c comment.Comment
		if err := rows.Scan(&c.ID, &c.CulturalID, &c.CulturalType, &c.UserName, &c.Comment, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		commentaries = append(commentaries, c)
	}

	fmt.Println("Fetched commentaries:", commentaries)

	return commentaries, nil
}
