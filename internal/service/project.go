package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
)

type Project struct {
	repo repository.Project
}

func NewProject(repo repository.Project) *Project {
	return &Project{
		repo: repo,
	}
}

func (s *Project) FindById(ctx context.Context, id string) (entity.Project, error) {
	return s.repo.FindById(ctx, id)
}

func (s *Project) FindByNameMany(ctx context.Context, name string) ([]entity.Project, error) {
	return s.repo.FindByNameMany(ctx, name)
}

func (s *Project) Save(ctx context.Context, project entity.Project) error {
	err := project.Validate()
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, project)
}

func (s *Project) Update(ctx context.Context, project entity.Project) error {
	err := project.Validate()
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, project)
}

func (s *Project) DeleteById(ctx context.Context, id string) error {
	return s.repo.DeleteById(ctx, id)
}
