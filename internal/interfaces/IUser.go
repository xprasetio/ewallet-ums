package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	DeleteUserSession(ctx context.Context, token string) error
}