package repository

import (
	"errors"
	_ "embed"
	"context"
	"database/sql"
	"log"

    "github.com/nleof/goyesql"

	"poc2/back/model"

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
	UpdateUserData(ctx context.Context, user model.User) error
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
	_, err := r.db.ExecContext(ctx, userQueries["register-user"],
		user.Name,
		user.Email,
		user.Document,
		user.CompanyName,
		user.Type,
		user.Password,
	)
	log.Println(err)
	return err
}

func (r *userRepository) FetchUserDataByID(ctx context.Context, userID int) (model.User, error) {
	return model.User{}, errors.New("not implemented")
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

func (r *userRepository) UpdateUserData(ctx context.Context, user model.User) error {
	return errors.New("not implemented")
}

func (r *userRepository) AddEventToFavorites(ctx context.Context, userID int, event model.Event) error {
	return errors.New("not implemented")
}

func (r *userRepository) AddAttractionToFavorites(ctx context.Context, userID int, attraction model.TouristAttraction) error {
	return errors.New("not implemented")
}

func (r *userRepository) RemoveEventFromFavorites(ctx context.Context, userID, eventID int) error {
	return errors.New("not implemented")
}

func (r *userRepository) RemoveAttractionFromFavorites(ctx context.Context, userID, attractionID int) error {
	return errors.New("not implemented")
}

func (r *userRepository) DeleteUserByID(ctx context.Context, userID int) error {
	return errors.New("not implemented")
}