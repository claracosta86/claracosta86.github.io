package model


// RegisterUserRequest represents the request for user registration
type RegisterUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Document    string `json:"document"`
	CompanyName string `json:"companyName"`
	Password    string `json:"password"`
	Type        string `json:"type"`
}

// LoginUserRequest represents the request for user login
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUserResponse represents the response for user login
type LoginUserResponse struct {
	UserID int    `json:"userID"`
	Type   string `json:"type"`
}

// GetUserProfileResponse represents the response for user profile
type GetUserProfileResponse struct {
	UserID      int    `json:"userID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CompanyName string `json:"companyName"`
	Type        string `json:"type"`
}

// UpdateUserProfileRequest represents the request for updating user profile
type UpdateUserProfileRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CompanyName string `json:"companyName"`
}

// ChangePasswordRequest represents the request for changing password
type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// Type represents the type of user for the application
type Type struct {
	Type string `json:"userType"`
}

// Information represents the user information for the application
type Information struct {
	Type  string `json:"userType"`
	ID    string `json:"userID"`
}