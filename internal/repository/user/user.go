package user

import (
	"context"
	"database/sql"
	"fmt"

	"sarkor-test/internal/entity"
)

type Repository struct {
	*sql.DB
}

func New(DB *sql.DB) *Repository {
	return &Repository{DB}
}

func (r Repository) GetAuthDetail(ctx context.Context, detail CookieAuth) error {
	query := fmt.Sprintf(
		`select user_id from users where user_id = '%d' and login = '%s'`,
		detail.UserID,
		detail.Login,
	)

	var intCheck int
	return r.QueryRowContext(ctx, query).Scan(&intCheck)
}

func (r Repository) GetByLogin(ctx context.Context, data Auth) (entity.User, error) {
	var userData entity.User

	query := fmt.Sprintf(
		`select user_id, password from users where login = '%s'`,
		data.Login,
	)

	err := r.QueryRowContext(ctx, query).Scan(&userData.UserID, &userData.Password)

	return userData, err
}

func (r Repository) GetByName(ctx context.Context, name string) (entity.User, error) {
	var userData entity.User

	query := fmt.Sprintf(
		`select user_id, name, age from users where name = '%s'`,
		name,
	)

	err := r.QueryRowContext(ctx, query).Scan(&userData.UserID, &userData.Name, &userData.Age)

	return userData, err
}

func (r Repository) Create(ctx context.Context, data Create) error {
	query := fmt.Sprintf(
		`insert into users (login, password, name, age) values ('%s', '%s', '%s', '%d')`,
		data.Login, data.Password, data.Name, data.Age,
	)

	_, err := r.ExecContext(ctx, query)

	return err
}
