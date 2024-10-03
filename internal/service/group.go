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

func (group *GroupService) CreateGroup(group model.Group) (string, error) {
	return group.repository.CreateGroup(group)
}

func (group *GroupService) GetGroup(id string) (model.Group, error) {
	return group.repository.GetGroup(id)
}

func (group *GroupService) UpdateGroup(id model.Group) (string, error) {
	return group.repository.UpdateGroup(id)
}

func (group *GroupService) DeleteGroup(id string) error {
	return group.repository.DeleteGroup(id)
}
