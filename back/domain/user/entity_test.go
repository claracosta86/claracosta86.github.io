package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		email       string
		document    string
		companyName string
		password    string
		userType    UserType
		expectError bool
	}{
		{
			name:        "valid common user",
			userName:    "John Doe",
			email:       "john@example.com",
			document:    "12345678901",
			companyName: "",
			password:    "password123",
			userType:    UserTypeCommon,
			expectError: false,
		},
		{
			name:        "valid organizer user",
			userName:    "Company Inc",
			email:       "contact@company.com",
			document:    "12345678000199",
			companyName: "Company Inc",
			password:    "password123",
			userType:    UserTypeOrganizer,
			expectError: false,
		},
		{
			name:        "empty name",
			userName:    "",
			email:       "john@example.com",
			document:    "12345678901",
			companyName: "",
			password:    "password123",
			userType:    UserTypeCommon,
			expectError: true,
		},
		{
			name:        "invalid email",
			userName:    "John Doe",
			email:       "invalid-email",
			document:    "12345678901",
			companyName: "",
			password:    "password123",
			userType:    UserTypeCommon,
			expectError: true,
		},
		{
			name:        "short password",
			userName:    "John Doe",
			email:       "john@example.com",
			document:    "12345678901",
			companyName: "",
			password:    "123",
			userType:    UserTypeCommon,
			expectError: true,
		},
		{
			name:        "invalid user type",
			userName:    "John Doe",
			email:       "john@example.com",
			document:    "12345678901",
			companyName: "",
			password:    "password123",
			userType:    "invalid",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.userName, tt.email, tt.document, tt.companyName, tt.password, tt.userType)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if user == nil {
				t.Errorf("Expected user but got nil")
				return
			}
			
			// Verify the user was created with correct values
			if user.Name != tt.userName {
				t.Errorf("Expected name %s, got %s", tt.userName, user.Name)
			}
			
			if user.Email != tt.email {
				t.Errorf("Expected email %s, got %s", tt.email, user.Email)
			}
			
			if user.Type != string(tt.userType) {
				t.Errorf("Expected type %s, got %s", tt.userType, user.Type)
			}
		})
	}
}

func TestUserType_Validate(t *testing.T) {
	tests := []struct {
		name        string
		userType    UserType
		expectError bool
	}{
		{"valid common", UserTypeCommon, false},
		{"valid organizer", UserTypeOrganizer, false},
		{"invalid type", "invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.userType.Validate()
			
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestUser_GetDocumentType(t *testing.T) {
	commonUser := &User{Type: string(UserTypeCommon)}
	organizerUser := &User{Type: string(UserTypeOrganizer)}
	
	if commonUser.GetDocumentType() != DocumentTypeCPF {
		t.Errorf("Common user should have CPF document type")
	}
	
	if organizerUser.GetDocumentType() != DocumentTypeCNPJ {
		t.Errorf("Organizer user should have CNPJ document type")
	}
}
