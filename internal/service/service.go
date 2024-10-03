package service

import (
	"Diploma-project-backend/internal/model"
	"Diploma-project-backend/internal/repository"
)

type Users interface {
	CreateUser(user model.Users) (string, error)
	GetUser(email, password string) (model.Users, error)
	UpdateUser(user model.Users) (string, error)
	DeleteUser(id string, isProfessor bool) error
}

type Service struct {
	Users
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repository.Users),
	}
}
