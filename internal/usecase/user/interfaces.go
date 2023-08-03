package user

import (
	"context"

	"sarkor-test/internal/entity"
	"sarkor-test/internal/repository/phone"
	"sarkor-test/internal/repository/user"
)

type User interface {
	GetAuthDetail(ctx context.Context, detail user.CookieAuth) error
	GetByName(ctx context.Context, name string) (entity.User, error)
	Create(ctx context.Context, data user.Create) error
}

type Phone interface {
	GetAll(ctx context.Context, filter phone.Filter) ([]phone.List, error)
	Create(ctx context.Context, data phone.Create) error
	Update(ctx context.Context, data phone.Update) error
	Delete(ctx context.Context, id int) error
}
