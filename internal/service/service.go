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

type Group interface {
	CreateGroup(group model.Group) (string, error)
	GetGroup(id string) (model.Group, error)
	UpdateGroup(group model.Group) (string, error)
	DeleteGroup(id string) error
}

type Service struct {
	Users
	Group
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repository.Users),
		Group: NewGroupService(repository.Group),
	}
}
