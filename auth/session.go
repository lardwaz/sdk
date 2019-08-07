package auth

import (
	"context"
)

// SessionContextKey is the auth session context key
const SessionContextKey = "AuthSession"

// Session represents signature for a session
type Session interface {
	// User returns the user info
	User() *User

	// IsAllowed returns whether user can perform an action
	IsAllowed(action string) bool
}

// GetSessionFromCtx returns session from context
func GetSessionFromCtx(ctx context.Context) (Session, error) {
	ctxVal := ctx.Value(SessionContextKey)
	if ctxVal == nil {
		// Session ctx was never set most likely
		// so auth is not enabled
		return nil, nil
	}

	session, ok := ctxVal.(Session)
	if !ok || session == nil {
		return nil, ErrInvalidSession
	}

	return session, nil
}
