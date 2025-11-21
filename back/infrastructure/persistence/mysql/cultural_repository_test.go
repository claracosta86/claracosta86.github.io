package mysql

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"poc2/back/domain/cultural"
)

func TestCulturalRepository_SaveEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewCulturalRepository(db)
	ctx := context.Background()

	price := cultural.NewPrice("10.00")
	location, _ := cultural.NewLocation("Location")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO events")).
			WithArgs("Title", "Desc", location, "Start", "End", "2h", price, true, 1, "img.jpg").
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := repo.SaveEvent(ctx, "Title", "Desc", location, "Start", "End", "2h", price, true, 1, "img.jpg")
		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})
}

func TestCulturalRepository_SaveTouristAttraction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewCulturalRepository(db)
	ctx := context.Background()

	price := cultural.NewPrice("20.00")
	location, _ := cultural.NewLocation("Location")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO tourist_attractions")).
			WithArgs("Title", "Desc", location, "9-5", price, true, 1, "img.jpg").
			WillReturnResult(sqlmock.NewResult(2, 1))

		id, err := repo.SaveTouristAttraction(ctx, "Title", "Desc", location, "9-5", price, true, 1, "img.jpg")
		assert.NoError(t, err)
		assert.Equal(t, 2, id)
	})
}

func TestCulturalRepository_FindEventByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewCulturalRepository(db)
	ctx := context.Background()
	id := 1

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description", "location", "start_date", "end_date", "duration_hours", "price", "is_accessible", "organizer_id", "email", "image"}).
			AddRow(1, "Title", "Desc", "Loc", "Start", "End", "2h", "10.00", true, 1, "org@example.com", "img.jpg")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT e.id, IFNULL(e.title, ''), IFNULL(e.description, ''), IFNULL(e.location, ''), IFNULL(e.start_date, ''), IFNULL(e.end_date, ''), IFNULL(e.duration_hours, ''), e.price, e.is_accessible, IFNULL(e.organizer_id, 0), IFNULL(u.email, ''), IFNULL(e.image, '') FROM events e LEFT JOIN users u ON e.organizer_id = u.id WHERE e.id = ?")).
			WithArgs(id).
			WillReturnRows(rows)

		event, err := repo.FindEventByID(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, "Title", event.Title)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT e.id")).
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.FindEventByID(ctx, id)
		assert.Error(t, err)
		assert.Equal(t, "cultural event not found", err.Error())
	})
}

func TestCulturalRepository_FindTouristAttractionByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewCulturalRepository(db)
	ctx := context.Background()
	id := 1

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description", "location", "working_hours", "price", "is_accessible", "organizer_id", "email", "image"}).
			AddRow(1, "Title", "Desc", "Loc", "9-5", "20.00", true, 1, "org@example.com", "img.jpg")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT ta.id, IFNULL(ta.title, ''), IFNULL(ta.description, ''), IFNULL(ta.location, ''), IFNULL(ta.working_hours, ''), ta.price, ta.is_accessible, IFNULL(ta.organizer_id, 0), IFNULL(u.email, ''), IFNULL(ta.image, '') FROM tourist_attractions ta LEFT JOIN users u ON ta.organizer_id = u.id WHERE ta.id = ?")).
			WithArgs(id).
			WillReturnRows(rows)

		attraction, err := repo.FindTouristAttractionByID(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, "Title", attraction.Title)
	})
}
