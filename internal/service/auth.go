package service

import (
	"Diploma-project-backend/internal/model"
	"Diploma-project-backend/internal/repository"
)

type AuthService struct {
	repository repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

func (auth *AuthService) CreateUser(user model.Users) (string, error) {
	return auth.repository.CreateUser(user)
}

func (auth *AuthService) GetUser(email, password string) (model.Users, error) {
	return auth.repository.GetUser(email, password)
}

func (auth *AuthService) UpdateUser(user model.Users) (string, error) {
	return auth.repository.UpdateUser(user)
}

func (auth *AuthService) DeleteUser(id string, isProfessor bool) error {
	return auth.repository.DeleteUser(id, isProfessor)
}
