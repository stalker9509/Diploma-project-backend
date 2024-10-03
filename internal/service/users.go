package service

import (
	"Diploma-project-backend/internal/model"
	"Diploma-project-backend/internal/repository"
)

type UsersService struct {
	repository repository.Users
}

func NewUsersService(repository repository.Users) *UsersService {
	return &UsersService{repository: repository}
}

func (auth *UsersService) CreateUser(user model.Users) (string, error) {
	return auth.repository.CreateUser(user)
}

func (auth *UsersService) GetUser(email, password string) (model.Users, error) {
	return auth.repository.GetUser(email, password)
}

func (auth *UsersService) UpdateUser(user model.Users) (string, error) {
	return auth.repository.UpdateUser(user)
}

func (auth *UsersService) DeleteUser(id string, isProfessor bool) error {
	return auth.repository.DeleteUser(id, isProfessor)
}
