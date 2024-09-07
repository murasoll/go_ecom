// services/user_context.go

package services

import (
	"context"
	"ecomerce/models"
	"errors"
)

type contextKey string

const userContextKey contextKey = "user"

// ContextWithUser adds a user to the given context
func ContextWithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// UserFromContext retrieves the user from the given context
func UserFromContext(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		return nil, errors.New("user not found in context")
	}
	return user, nil
}
