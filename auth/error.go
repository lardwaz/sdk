package auth

import "errors"

// Auth errors
var (
	ErrTokenNotFound  = errors.New("Error finding token in request")
	ErrFetchUserData  = errors.New("Error fetching user data")
	ErrInvalidSession = errors.New("Invalid session")
)
