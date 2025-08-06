package repository

import (
	stderrors "errors"
	_ "embed"
	"context"
	"database/sql"
	"log"
	"strings"

    "github.com/nleof/goyesql"

	"poc2/back/model"
	"poc2/back/lib/errors"

)

var (
	//go:embed queries/users.sql
	userEmbed []byte
	userQueries goyesql.Queries
)

type UserRepository interface {
	SaveUserData(ctx context.Context, user model.User) error
	FetchUserDataByID(ctx context.Context, userID int) (model.User, error)
	FetchUserIDByEmail(ctx context.Context, email string) (int, error)
	CheckUserPassword(ctx context.Context, userID int, password string) (bool, error)
	UpdateUserProfile(ctx context.Context, user model.User) error
	UpdateUserPassword(ctx context.Context, userID int, newPassword string) error
	AddEventToFavorites(ctx context.Context, userID int, event model.Event) error
	AddAttractionToFavorites(ctx context.Context, userID int, attraction model.TouristAttraction) error
	RemoveEventFromFavorites(ctx context.Context, userID int, eventID int) error
	RemoveAttractionFromFavorites(ctx context.Context, userID int, attractionID int) error
	DeleteUserByID(ctx context.Context, userID int) error
}

type userRepository struct {
	db *sql.DB
}

func init() {
    userQueries = goyesql.MustParseBytes(userEmbed)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) SaveUserData(ctx context.Context, user model.User) error {
	documentType := "CPF"
	if user.Type == "organizer" {
		documentType = "CNPJ"
	}

	_, err := r.db.ExecContext(ctx, userQueries["register-user"],
		user.Name,
		user.Email,
		user.Document,
		user.CompanyName,
		user.Type,
		user.Password,
		documentType,
	)

	if err != nil && strings.Contains(err.Error(), "Error 1062") {
		return errors.ErrUserAlreadyExists
	}

	return err
}

func (r *userRepository) FetchUserDataByID(ctx context.Context, userID int) (model.User, error) {
	return model.User{}, stderrors.New("not implemented")
}

func (r *userRepository) FetchUserIDByEmail(ctx context.Context, userEmail string) (int, error) {
	var userID int
	err := r.db.QueryRowContext(ctx, userQueries["fetch-user-id-by-email"], userEmail).Scan(&userID)
	if err != nil {
		log.Printf("Error fetching user ID by email: %v", err)
		return 0, err
	}
	return userID, nil
}

func (r *userRepository) CheckUserPassword(ctx context.Context, userID int, password string) (bool, error) {
	var storedPassword string
	err := r.db.QueryRowContext(ctx, userQueries["fetch-user-password-by-id"], userID).Scan(&storedPassword)
	if err != nil {
		log.Printf("Error fetching user password by ID: %v", err)
		return false, err
	}
	return storedPassword == password, nil
}

func (r *userRepository) UpdateUserProfile(ctx context.Context, user model.User) error {
	return stderrors.New("not implemented")
}

func (r *userRepository) UpdateUserPassword(ctx context.Context, userID int, newPassword string) error {
	_, err := r.db.ExecContext(ctx, userQueries["update-user-password"],
		newPassword,
		userID,
	)
	return err
}

func (r *userRepository) AddEventToFavorites(ctx context.Context, userID int, event model.Event) error {
	return stderrors.New("not implemented")
}

func (r *userRepository) AddAttractionToFavorites(ctx context.Context, userID int, attraction model.TouristAttraction) error {
	return stderrors.New("not implemented")
}

func (r *userRepository) RemoveEventFromFavorites(ctx context.Context, userID, eventID int) error {
	return stderrors.New("not implemented")
}

func (r *userRepository) RemoveAttractionFromFavorites(ctx context.Context, userID, attractionID int) error {
	return stderrors.New("not implemented")
}

func (r *userRepository) DeleteUserByID(ctx context.Context, userID int) error {
	return stderrors.New("not implemented")
}