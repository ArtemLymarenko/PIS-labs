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
	repo repository.Project
}

func NewProject(repo repository.Project) *Project {
	return &Project{
		repo: repo,
	}
}

func (s *Project) FindProjectById(ctx context.Context, id string) (entity.Project, error) {
	return s.repo.FindById(ctx, id)
}

func (s *Project) FindProjectsByName(ctx context.Context, name string) ([]entity.Project, error) {
	return s.repo.FindByNameMany(ctx, name)
}

func (s *Project) SaveProject(ctx context.Context, project entity.Project) error {
	err := project.Validate()
	if err != nil {
		return ErrInvalidProjectEntity
	}

	return s.repo.Save(ctx, project)
}

func (s *Project) UpdateProject(ctx context.Context, project entity.Project) error {
	err := project.Validate()
	if err != nil {
		return ErrInvalidProjectEntity
	}

	return s.repo.Update(ctx, project)
}

func (s *Project) DeleteProjectById(ctx context.Context, id string) error {
	return s.repo.DeleteById(ctx, id)
}
