package errors

import (
	"fmt"

)

var (
	ErrUserNotFound          = fmt.Errorf("user not found")
	ErrInvalidCredentials    = fmt.Errorf("invalid credentials")
	ErrUserAlreadyExists     = fmt.Errorf("user already exists")
)