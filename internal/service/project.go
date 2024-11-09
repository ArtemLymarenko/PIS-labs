package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
	"errors"
)

var (
	ErrInvalidProjectEntity = errors.New("failed to validate project entity")
)

type Project struct {
	projectRepo repository.Project
}

func NewProject(projectRepo repository.Project) *Project {
	return &Project{
		projectRepo: projectRepo,
	}
}

func (service *Project) FindProjectById(ctx context.Context, id string) (entity.Project, error) {
	return service.projectRepo.FindById(ctx, id)
}

func (service *Project) FindProjectsByName(ctx context.Context, name string) ([]entity.Project, error) {
	return service.projectRepo.FindByNameMany(ctx, name)
}

func (service *Project) SaveProject(ctx context.Context, project entity.Project) error {
	err := project.Validate()
	if err != nil {
		return ErrInvalidProjectEntity
	}

	//function to generate id (example: google/uuid)
	id := "123456789"
	project.SetId(id)

	return service.projectRepo.Save(ctx, project)
}

func (service *Project) UpdateProject(ctx context.Context, project entity.Project) error {
	err := project.Validate()
	if err != nil {
		return ErrInvalidProjectEntity
	}

	return service.projectRepo.Update(ctx, project)
}

func (service *Project) DeleteProjectById(ctx context.Context, id string) error {
	return service.projectRepo.DeleteById(ctx, id)
}
