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

type Group interface {
	CreateGroup(group model.Group) (string, error)
	GetGroup(id string) (model.Group, error)
	GetAllGroups() ([]model.Group, error)
	UpdateGroup(group model.Group) (string, error)
	DeleteGroup(id string) error
}

type Repository struct {
	Users
	Group
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUsersPostgres(db),
		Group: NewGroupPostgres(db),
	}
}
