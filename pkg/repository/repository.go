package repository

import (
	"github.com/jmoiron/sqlx"
	"sync-muz-bot/pkg/models"
	"sync-muz-bot/pkg/repository/postgres"
)

type User interface {
	CreateUser(user models.User) (int, error)
	GetUser(id int64) (models.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: postgres.NewUserPostgres(db),
	}
}
