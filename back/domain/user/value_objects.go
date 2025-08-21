package user

import (
	"errors"
	"strings"
)

// Email represents an email address value object
type Email struct {
	value string
}

// String returns the email as a string
func (e *Email) String() string {
	return e.value
}

// isValidEmailFormat performs basic email format validation
func isValidEmailFormat(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	localPart := parts[0]
	domainPart := parts[1]
	
	if localPart == "" || domainPart == "" {
		return false
	}
	
	if !strings.Contains(domainPart, ".") {
		return false
	}
	
	return true
}

// Password represents a password value object
type Password struct {
	value string
}

// NewPassword creates a new password value object
func NewPassword(password string) (*Password, error) {
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	
	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}
	
	return &Password{value: password}, nil
}

// String returns the password as a string
func (p *Password) String() string {
	return p.value
}

// Verify checks if the provided password matches
func (p *Password) Verify(password string) bool {
	return p.value == password
}

// Name represents a name value object
type Name struct {
	value string
}

// NewName creates a new name value object
func NewName(name string) (*Name, error) {
	name = strings.TrimSpace(name)
	
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	
	if len(name) < 2 {
		return nil, errors.New("name must be at least 2 characters")
	}
	
	return &Name{value: name}, nil
}

// String returns the name as a string
func (n *Name) String() string {
	return n.value
}

// Document represents a document value object
type Document struct {
	value string
}

// NewDocument creates a new document value object
func NewDocument(document string) (*Document, error) {
	document = strings.TrimSpace(document)
	
	if document == "" {
		return nil, errors.New("document cannot be empty")
	}
	
	return &Document{value: document}, nil
}

// String returns the document as a string
func (d *Document) String() string {
	return d.value
}
