package user

import (
	"errors"
	"strings"
)

// User represents the core user entity in the domain
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Document    string `json:"document"`
	CompanyName string `json:"companyName"`
	Password    string `json:"password"`
	Type        string `json:"type"`
}

// UserType represents valid user types
type UserType string

const (
	UserTypeCommon   UserType = "common"
	UserTypeOrganizer UserType = "organizer"
)

// Validate checks if the UserType is valid
func (ut UserType) Validate() (UserType, error) {
	switch ut {
	case UserTypeCommon, UserTypeOrganizer:
		return ut, nil
	default:
		return "", errors.New("invalid user type")
	}
}

// DocumentType represents valid document types
type DocumentType string

const (
	DocumentTypeCPF  DocumentType = "CPF"
	DocumentTypeCNPJ DocumentType = "CNPJ"
)

// NewUser creates a new user with validation
func NewUser(name, email, document, companyName, password string, userType UserType) (*User, error) {
	if err := validateUserData(name, email, document, password, userType); err != nil {
		return nil, err
	}

	user := &User{
		Name:        strings.TrimSpace(name),
		Email:       strings.ToLower(strings.TrimSpace(email)),
		Document:    strings.TrimSpace(document),
		CompanyName: strings.TrimSpace(companyName),
		Password:    password,
		Type:        string(userType),
	}

	return user, nil
}

// validateUserData validates user input data
func validateUserData(name, email, document, password string, userType UserType) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}

	if !isValidEmail(email) {
		return errors.New("invalid email format")
	}

	if strings.TrimSpace(document) == "" {
		return errors.New("document cannot be empty")
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("password cannot be empty")
	}

	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if userType != UserTypeCommon && userType != UserTypeOrganizer {
		return errors.New("invalid user type")
	}

	return nil
}

// isValidEmail performs basic email validation
func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// GetDocumentType returns the appropriate document type based on user type
func (u *User) GetDocumentType() DocumentType {
	if u.Type == string(UserTypeOrganizer) {
		return DocumentTypeCNPJ
	}
	return DocumentTypeCPF
}

// UpdateProfile updates user profile information
func (u *User) UpdateProfile(name, email, companyName string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}

	if !isValidEmail(email) {
		return errors.New("invalid email format")
	}

	u.Name = strings.TrimSpace(name)
	u.Email = strings.ToLower(strings.TrimSpace(email))
	u.CompanyName = strings.TrimSpace(companyName)

	return nil
}

// ChangePassword changes the user's password
func (u *User) ChangePassword(currentPassword, newPassword string) error {
	if u.Password != currentPassword {
		return errors.New("current password is incorrect")
	}

	if strings.TrimSpace(newPassword) == "" {
		return errors.New("new password cannot be empty")
	}

	if len(newPassword) < 6 {
		return errors.New("new password must be at least 6 characters")
	}

	u.Password = newPassword
	return nil
}
