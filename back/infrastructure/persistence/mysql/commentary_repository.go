package mysql


import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/nleof/goyesql"

	"poc2/back/domain/commentary"

)

var (
	//go:embed queries/commentary.sql
	commentaryEmbed []byte
	commentaryQueries goyesql.Queries
)

type commentaryRepository struct {
	db *sql.DB
}

func init() {
	commentaryQueries = goyesql.MustParseBytes(commentaryEmbed)
}

// NewCommentaryRepository creates a new MySQL commentary repository
func NewCommentaryRepository(db *sql.DB) commentary.Repository {
	return &commentaryRepository{
		db: db,
	}
}


func (r *commentaryRepository) SaveCommentary(ctx context.Context, culturalID int, culturalType string, userID int, commentary string) error {
	fmt.Printf("Saving commentary for culturalID: %d, culturalType: %s, userID: %d\n", culturalID, culturalType, userID)
	_, err := r.db.ExecContext(ctx, commentaryQueries["save-commentary"],
		culturalID,
		culturalType,
		userID,
		commentary,
	)
	if err != nil {
		return fmt.Errorf("failed to save commentary: %w", err)
	}
	return nil
}

func (r *commentaryRepository) FindCommentariesByCultural(ctx context.Context, culturalID int, culturalType string) (commentary.Commentaries, error) {
	rows, err := r.db.QueryContext(ctx, commentaryQueries["fetch-commentaries-by-cultural"], culturalID, culturalType)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commentaries: %w", err)
	}
	defer rows.Close()

	var commentaries []commentary.Commentary
	for rows.Next() {
		var c commentary.Commentary
		if err := rows.Scan(&c.ID, &c.CulturalID, &c.CulturalType, &c.UserName, &c.Commentary, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan commentary: %w", err)
		}
		commentaries = append(commentaries, c)
	}

	fmt.Println("Fetched commentaries:", commentaries)

	return commentaries, nil
}