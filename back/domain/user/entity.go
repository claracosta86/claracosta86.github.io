package user

import (
	"errors"
	"strings"
)

// User represents the core user entity in the domain
type User struct {
	ID          int
	Name        string
	Email       string
	Document    string
	CompanyName string
	Password    string
	Type        string
	CreatedAt   string
}

// UserType represents valid user types
type UserType string

const (
	UserTypeCommon    UserType = "common"
	UserTypeOrganizer UserType = "organizer"
)

type CulturalList struct {
	ID    int
	Title string
	Type  string
}

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
	// Validate Name
	nameVO, err := NewName(name)
	if err != nil {
		return nil, err
	}

	// Validate Email
	email = strings.ToLower(strings.TrimSpace(email))
	if !isValidEmailFormat(email) {
		return nil, errors.New("invalid email format")
	}

	// Validate Document
	documentVO, err := NewDocument(document)
	if err != nil {
		return nil, err
	}

	// Validate Password
	passwordVO, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	// Validate UserType
	validUserType, err := userType.Validate()
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:        nameVO.String(),
		Email:       email,
		Document:    documentVO.String(),
		CompanyName: strings.TrimSpace(companyName),
		Password:    passwordVO.String(),
		Type:        string(validUserType),
	}

	return user, nil
}

// GetDocumentType returns the appropriate document type based on user type
func (u *User) GetDocumentType() DocumentType {
	if u.Type == string(UserTypeOrganizer) {
		return DocumentTypeCNPJ
	}
	return DocumentTypeCPF
}

// ChangePassword changes the user's password
func (u *User) ChangePassword(currentPassword, newPassword string) error {
	if u.Password != currentPassword {
		return errors.New("current password is incorrect")
	}

	u.Password = newPassword
	return nil
}
