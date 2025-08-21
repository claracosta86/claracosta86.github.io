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
