package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
	"errors"
)

type Project struct {
	repo repository.ProjectRepository
}

func NewProject(repo repository.ProjectRepository) *Project {
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
	if project.Name == "" {
		return errors.New("project name cannot be empty")
	}

	return s.repo.Save(ctx, project)
}

func (s *Project) Update(ctx context.Context, project entity.Project) error {
	existingProject, err := s.repo.FindById(ctx, project.Id)
	if err != nil {
		return err
	}

	// Викликаємо метод з репозиторію для оновлення
	return s.repo.Update(ctx, existingProject)
}

func (s *Project) DeleteById(ctx context.Context, id string) error {
	return s.repo.DeleteById(ctx, id)
}
