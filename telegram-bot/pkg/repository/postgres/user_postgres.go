package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sync-muz-bot/pkg/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(id int64, username string) error {
	query := fmt.Sprintf("INSERT INTO %s (id, username) values ($1, $2)", usersTable)
	row := r.db.QueryRow(query, id, username)
	err := row.Err()
	return err
}

func (r *UserPostgres) GetUser(id int64) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *UserPostgres) ExistsUser(id int64) (bool, error) {
	var exists = false
	query := fmt.Sprintf("SELECT EXISTS (SELECT id FROM %s WHERE id=$1)", usersTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&exists)
	return exists, err
}
