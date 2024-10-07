package service

import (
	"Diploma-project-backend/internal/model"
	"Diploma-project-backend/internal/repository"
)

type GroupService struct {
	repository repository.Group
}

func NewGroupService(repository repository.Group) *GroupService {
	return &GroupService{repository: repository}
}

func (g *GroupService) CreateGroup(group model.Group) (string, error) {
	return g.repository.CreateGroup(group)
}

func (g *GroupService) GetGroup(id string) (model.Group, error) {
	return g.repository.GetGroup(id)
}

func (g *GroupService) GetAllGroups() ([]model.Group, error) {
	return g.repository.GetAllGroups()
}

func (g *GroupService) UpdateGroup(id model.Group) (string, error) {
	return g.repository.UpdateGroup(id)
}

func (g *GroupService) DeleteGroup(id string) error {
	return g.repository.DeleteGroup(id)
}
