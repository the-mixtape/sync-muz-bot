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

func (r *UserPostgres) CreateUser(user models.User) (int, error) {
	var id = 0
	query := fmt.Sprintf("INSERT INTO %s (id, username) values ($1, $2)", usersTable)

	row := r.db.QueryRow(query, user.Id, user.Username)
	err := row.Scan(&id)
	return id, err
}

func (r *UserPostgres) GetUser(id int64) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}
