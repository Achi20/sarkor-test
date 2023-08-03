package auth

import (
	"context"

	"sarkor-test/internal/entity"
	"sarkor-test/internal/repository/user"
)

type User interface {
	GetAuthDetail(ctx context.Context, detail user.CookieAuth) error
	GetByLogin(ctx context.Context, data user.Auth) (entity.User, error)
}
