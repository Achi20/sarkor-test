package sqlite

import (
	"database/sql"
)

func New() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./gopher.db")
	if err != nil {
		return nil, err
	}

	statement, err := database.Prepare("create table if not exists users (user_id integer primary key, " +
		"login text not null, password text not null, name text not null, age integer not null, created_at text default current_timestamp)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	statement.Exec()

	statement, err = database.Prepare("create table if not exists phones (phone_id integer primary key, " +
		"phone text not null unique, description text, is_fax boolean not null, user_id integer not null, created_at text default current_timestamp, " +
		"foreign key (user_id) references users (user_id))")
	if err != nil {
		return nil, err
	}
	statement.Exec()

	// to improve performance on auth queries
	statement, err = database.Prepare("create index if not exists idx_users_login on users (login)")
	if err != nil {
		return nil, err
	}
	statement.Exec()

	if err = statement.Close(); err != nil {
		return nil, err
	}

	return database, nil
}
