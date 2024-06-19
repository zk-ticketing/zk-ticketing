package util

import "context"

type ContextKey string

const (
	UserIDContextKey    ContextKey = "userID"
	UserEmailContextKey ContextKey = "userEmail"
)

func SetUserIDInContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDContextKey, userID)
}

func SetUserEmailInContext(ctx context.Context, userEmail string) context.Context {
	return context.WithValue(ctx, UserEmailContextKey, userEmail)
}

func GetUserIDFromContext(ctx context.Context) string {
	userID := ctx.Value(UserIDContextKey)
	if userID == nil {
		return ""
	}
	return userID.(string)
}

func GetUserEmailFromContext(ctx context.Context) string {
	userEmail := ctx.Value(UserEmailContextKey)
	if userEmail == nil {
		return ""
	}
	return userEmail.(string)
}
