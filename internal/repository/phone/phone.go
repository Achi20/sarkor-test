package phone

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

type Repository struct {
	*sql.DB
}

func New(DB *sql.DB) *Repository {
	return &Repository{DB}
}

func (r Repository) GetAll(ctx context.Context, filter Filter) ([]List, error) {
	var list []List

	query := "select user_id, phone, description, is_fax from phones"

	if filter.Phone != nil {
		query += fmt.Sprintf(` where phone like '%s'`, "%"+*filter.Phone+"%")
	}

	rows, err := r.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var detail List

		err = rows.Scan(&detail.UserID, &detail.Phone, &detail.Description, &detail.IsFax)
		if err != nil {
			return nil, err
		}

		list = append(list, detail)
	}

	return list, err
}

func (r Repository) Create(ctx context.Context, data Create) error {
	query := fmt.Sprintf(
		`insert into phones (phone, description, is_fax, user_id) values ('%s', '%s', '%v', '%d')`,
		data.Phone, data.Description, data.IsFax, data.UserID,
	)

	_, err := r.ExecContext(ctx, query)
	if errors.As(err, &sqlite3.ErrConstraintUnique) {
		return errors.New("record already exists")

	}

	return err
}

func (r Repository) Update(ctx context.Context, data Update) error {
	query := fmt.Sprintf(
		`update phones set phone = '%s', description = '%s', is_fax = '%v', user_id = '%d' where phone_id = '%d'`,
		data.Phone, data.Description, data.IsFax, data.UserID, data.PhoneID,
	)

	_, err := r.ExecContext(ctx, query)
	if errors.As(err, &sqlite3.ErrConstraintUnique) {
		return errors.New("record already exists")
	}

	return err
}

func (r Repository) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`delete from phones where phone_id = '%d'`, id)

	_, err := r.ExecContext(ctx, query)

	return err
}
