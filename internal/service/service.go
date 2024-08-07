package service

import (
	"Diploma-project-backend/internal/repository"
)

type Service struct {
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
