package mysql

import (
	"context"
	"database/sql"
	"strings"
	_ "embed"

	"github.com/nleof/goyesql"

	"poc2/back/domain/user"
	"poc2/back/lib/errors"
)

var (
	//go:embed queries/users.sql
	userEmbed []byte
	userQueries goyesql.Queries
)

type userRepository struct {
	db *sql.DB
}

func init() {
	userQueries = goyesql.MustParseBytes(userEmbed)
}

// NewUserRepository creates a new MySQL user repository
func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Save(ctx context.Context, user *user.User) error {
	documentType := user.GetDocumentType()
	
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

func (r *userRepository) FindByID(ctx context.Context, id int) (*user.User, error) {
	var u user.User
	err := r.db.QueryRowContext(ctx, userQueries["fetch-user-by-id"], id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Document,
		&u.CompanyName,
		&u.Type,
		&u.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}
	
	return &u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := r.db.QueryRowContext(ctx, userQueries["fetch-user-by-email"], email).Scan(
		&u.ID,
		&u.Type,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}
	
	return &u, nil
}

func (r *userRepository) Update(ctx context.Context, user *user.User) error {
	_, err := r.db.ExecContext(ctx, userQueries["update-user-profile"],
		user.Name,
		user.Email,
		user.CompanyName,
		user.ID,
	)
	return err
}

func (r *userRepository) DeleteUserFavorites(ctx context.Context, userID int) error {
	_, err := r.db.ExecContext(ctx, userQueries["delete-user-favorites-by-id"], userID)
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, userQueries["delete-user-by-id"], id)
	return err
}

func (r *userRepository) CheckPassword(ctx context.Context, userID int, password string) (bool, error) {
	var storedPassword string
	err := r.db.QueryRowContext(ctx, userQueries["fetch-user-password-by-id"], userID).Scan(&storedPassword)
	if err != nil {
		return false, err
	}
	return storedPassword == password, nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, userID int, newPassword string) error {
	_, err := r.db.ExecContext(ctx, userQueries["update-user-password"],
		newPassword,
		userID,
	)
	return err
}

func (r *userRepository) RemoveEvent(ctx context.Context, eventIDs []int) error {
	_, err := r.db.ExecContext(ctx, userQueries["delete-users-favorites-by-event-id"], eventIDs)
	return err
}

func (r *userRepository) RemoveTouristAttraction(ctx context.Context, touristAttractionIDs []int) error {
	_, err := r.db.ExecContext(ctx, userQueries["delete-users-favorites-by-tourist-attraction-id"], touristAttractionIDs)
	return err
}

func (r *userRepository) AddFavorite(ctx context.Context, userID int, culturalType string, culturalID int) error {
	_, err := r.db.ExecContext(ctx, userQueries["add-user-favorite"],
		userID,
		culturalType,
		culturalID,
	)
	return err
}

func (r *userRepository) RemoveFavorite(ctx context.Context, userID int, culturalType string, culturalID int) error {
	_, err := r.db.ExecContext(ctx, userQueries["remove-user-favorite"],
		userID,
		culturalType,
		culturalID,
	)
	return err
}

func (r *userRepository) GetFavoritesByUserID(ctx context.Context, userID int) ([]user.CulturalList, error) {
	rows, err := r.db.QueryContext(ctx, userQueries["fetch-user-favorites-by-id"], userID, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	favorites := make([]user.CulturalList, 0)
	for rows.Next() {
		var favorite user.CulturalList
		if err := rows.Scan(&favorite.ID, &favorite.Title, &favorite.Type); err != nil {
			return nil, err
		}

		favorites = append(favorites, favorite)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return favorites, nil
}

func (r *userRepository) UpdateLastSeenFavorite(ctx context.Context, userID, culturalID int, culturalType string) error {
	_, err := r.db.ExecContext(ctx, userQueries["update-last-seen-favorite"],
		userID,
		culturalType,
		culturalID,
	)
	return err
}

func (r *userRepository) GetCulturaisByOrganizerID(ctx context.Context, organizerID int) ([]user.CulturalList, error) {
	rows, err := r.db.QueryContext(ctx, userQueries["fetch-culturals-by-organizer-id"], organizerID, organizerID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	culturais := make([]user.CulturalList, 0)
	for rows.Next() {
		var cultural user.CulturalList
		if err := rows.Scan(&cultural.ID, &cultural.Title, &cultural.Type); err != nil {
			return nil, err
		}

		culturais = append(culturais, cultural)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return culturais, nil
}
