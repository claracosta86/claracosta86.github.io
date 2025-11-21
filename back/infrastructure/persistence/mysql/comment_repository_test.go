package mysql

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"poc2/back/domain/comment"
)

func TestCommentRepository_SaveComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewCommentRepository(db)
	ctx := context.Background()

	content, _ := comment.NewCommentContent("Nice!")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO comments")).
			WithArgs(1, "event", 1, content).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.SaveComment(ctx, 1, "event", 1, content)
		assert.NoError(t, err)
	})
}

func TestCommentRepository_FindCommentsByCultural(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewCommentRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "cultural_id", "cultural_type", "name", "content", "created_at"}).
			AddRow(1, 1, "event", "User", "Nice!", time.Now().Format(time.RFC3339))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT c.id, c.cultural_id, c.cultural_type, u.name, c.content, c.created_at FROM comments c")).
			WithArgs(1, "event").
			WillReturnRows(rows)

		comments, err := repo.FindCommentsByCultural(ctx, 1, "event")
		assert.NoError(t, err)
		assert.Len(t, comments, 1)
		assert.Equal(t, "Nice!", comments[0].Comment.String())
	})
}
