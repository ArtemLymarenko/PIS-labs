package repository

import (
	"PIS_labs/internal/domain/entity"
	"context"
	"errors"
)

var (
	ErrProjectNotFound   = errors.New("failed to find project")
	ErrProjectNotUpdated = errors.New("failed to update project")
)

type Project interface {
	FindById(ctx context.Context, id string) (entity.Project, error)
	FindByNameMany(ctx context.Context, name string) ([]entity.Project, error)
	Save(ctx context.Context, project entity.Project) error
	Update(ctx context.Context, project entity.Project) error
	DeleteById(ctx context.Context, id string) error
}

type ProjectRepository struct {
	projects []entity.Project
}

func (repository *ProjectRepository) FindById(ctx context.Context, id string) (entity.Project, error) {
	for _, project := range repository.projects {
		if project.Id == id {
			return project, nil
		}
	}
	return entity.Project{}, ErrProjectNotFound
}

func (repository *ProjectRepository) FindByNameMany(ctx context.Context, name string) ([]entity.Project, error) {
	var foundProjects []entity.Project
	for _, project := range repository.projects {
		if project.Name == name {
			foundProjects = append(foundProjects, project)
		}
	}

	return foundProjects, nil
}

func (repository *ProjectRepository) Save(ctx context.Context, newProject entity.Project) error {
	repository.projects = append(repository.projects, newProject)
	return nil
}

func (repository *ProjectRepository) Update(ctx context.Context, newProject entity.Project) error {
	for i, project := range repository.projects {
		if project.Id == newProject.Id {
			repository.projects[i] = newProject
			return nil
		}
	}

	return ErrProjectNotUpdated
}

func (repository *ProjectRepository) DeleteById(ctx context.Context, id string) error {
	for i, project := range repository.projects {
		if project.Id == id {
			repository.projects = append(repository.projects[:i], repository.projects[i+1:]...)
			return nil
		}
	}

	return nil
}
