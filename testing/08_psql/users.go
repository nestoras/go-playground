package psql

import (
	"database/sql"

	"github.com/pkg/errors"
)


var (
	ErrNotFound = errors.New("psql: resource could not be located")
)

type User struct {
	ID    int
	Name  string
	Email string
}

// UserStore is used to interact with our user store.
type UserStore struct {
	sql interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		QueryRow(query string, args ...interface{}) *sql.Row
	}
}

func (us *UserStore) Find(id int) (*User, error) {
	const query = `SELECT id, name, email FROM users WHERE id=$1;`
	row := us.sql.QueryRow(query, id)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	switch err {
	case sql.ErrNoRows:
		return nil, ErrNotFound
	case nil:
		return &user, nil
	default:
		return nil, errors.Wrap(err, "psql: error querying for user by id")
	}
}


func (us *UserStore) Create(user *User) error {
	const query = `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := us.sql.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return errors.Wrap(err, "psql: error creating new user")
	}
	return nil
}


func (us *UserStore) Delete(id int) error {
	const query = `DELETE FROM users WHERE id=$1;`
	_, err := us.sql.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "psql: error deleting user")
	}
	return nil
}