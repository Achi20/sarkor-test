package auth

import (
	"context"

	"sarkor-test/internal/repository/user"
)

type User interface {
	GetAuthDetail(ctx context.Context, detail user.CookieAuth) error
}
