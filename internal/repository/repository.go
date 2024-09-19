package repository

import (
	"Diploma-project-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.Users) (string, error)
	GetUser(email, password string) (model.Users, error)
	UpdateUser(user model.Users) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
