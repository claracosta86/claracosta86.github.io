package mysql

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"poc2/back/domain/user"
)

func TestUserRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	ctx := context.Background()
	u := &user.User{
		Name:        "Test User",
		Email:       "test@example.com",
		Document:    "12345678901",
		CompanyName: "Test Company",
		Password:    "password",
		Type:        "common",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users")).
			WithArgs(u.Name, u.Email, u.Document, u.CompanyName, u.Type, u.Password, user.DocumentTypeCPF).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Save(ctx, u)
		assert.NoError(t, err)
	})

	t.Run("duplicate entry", func(t *testing.T) {
		// Simulate MySQL Error 1062
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users")).
			WithArgs(u.Name, u.Email, u.Document, u.CompanyName, u.Type, u.Password, user.DocumentTypeCPF).
			WillReturnError(errors.New("Error 1062: Duplicate entry"))

		err := repo.Save(ctx, u)
		assert.Error(t, err)
		// Note: The repository checks for "Error 1062" string in the error
	})
}

func TestUserRepository_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()
	id := 1

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "document", "company_name", "type", "created_at"}).
			AddRow(1, "Test User", "test@example.com", "12345678901", "Test Company", "common", time.Now().Format(time.RFC3339))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, document, company_name, type, created_at FROM users WHERE id = ?")).
			WithArgs(id).
			WillReturnRows(rows)

		u, err := repo.FindByID(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, id, u.ID)
		assert.Equal(t, "Test User", u.Name)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, document, company_name, type, created_at FROM users WHERE id = ?")).
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)

		u, err := repo.FindByID(ctx, id)
		assert.Error(t, err)
		assert.Nil(t, u)
	})
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()
	email := "test@example.com"

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "type"}).
			AddRow(1, "common")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, type FROM users WHERE email = ?")).
			WithArgs(email).
			WillReturnRows(rows)

		u, err := repo.FindByEmail(ctx, email)
		assert.NoError(t, err)
		assert.Equal(t, 1, u.ID)
		assert.Equal(t, "common", u.Type)
	})
}

func TestUserRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()
	u := &user.User{
		ID:          1,
		Name:        "Updated Name",
		Email:       "updated@example.com",
		CompanyName: "Updated Company",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, company_name = ? WHERE id = ?")).
			WithArgs(u.Name, u.Email, u.CompanyName, u.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(ctx, u)
		assert.NoError(t, err)
	})
}
