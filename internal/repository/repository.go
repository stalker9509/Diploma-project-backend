package repository

import (
	"Diploma-project-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	CreateUser(user model.Users) (string, error)
	GetUser(email, password string) (model.Users, error)
	UpdateUser(user model.Users) (string, error)
	DeleteUser(id string, isProfessor bool) error
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUsersPostgres(db),
	}
}
