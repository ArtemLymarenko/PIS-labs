package repository

import (
	"PIS_labs/internal/domain/entity"
	"context"
	"errors"
)

// ErrProjectNotFound is returned when a project cannot be found by its unique identifier.
// ErrProjectNotUpdated is returned when a project cannot be updated by its unique identifier.
var (
	ErrProjectNotFound   = errors.New("failed to find project")
	ErrProjectNotUpdated = errors.New("failed to update project")
)

// Project defines the interface for managing project data in the repository layer.
// It includes methods for finding, saving, updating, and deleting projects.
type Project interface {
	FindById(ctx context.Context, id string) (entity.Project, error)
	FindByNameMany(ctx context.Context, name string) ([]entity.Project, error)
	Save(ctx context.Context, project entity.Project) error
	Update(ctx context.Context, project entity.Project) error
	DeleteById(ctx context.Context, id string) error
}

// ProjectRepository is an implementation of the Project interface.
// It holds a slice of projects and provides methods to perform CRUD operations.
type ProjectRepository struct {
	projects []entity.Project
}

// FindById searches for a project by its unique identifier.
// Parameters:
// - ctx: The context for managing the request.
// - id: The unique identifier of the project.
// Returns:
// - The found project, if any, or an error if the project is not found.
func (repository *ProjectRepository) FindById(ctx context.Context, id string) (entity.Project, error) {
	for _, project := range repository.projects {
		if project.Id == id {
			return project, nil
		}
	}
	return entity.Project{}, ErrProjectNotFound
}

// FindByNameMany searches for projects that match the specified name.
// Parameters:
// - ctx: The context for managing the request.
// - name: The name of the projects to search for.
// Returns:
// - A slice of matching projects, or an empty slice if no projects match.
func (repository *ProjectRepository) FindByNameMany(ctx context.Context, name string) ([]entity.Project, error) {
	var foundProjects []entity.Project
	for _, project := range repository.projects {
		if project.Name == name {
			foundProjects = append(foundProjects, project)
		}
	}

	return foundProjects, nil
}

// Save adds a new project to the repository.
// Parameters:
// - ctx: The context for managing the request.
// - newProject: The project to be saved.
// Returns:
// - An error, if any, during the saving process.
func (repository *ProjectRepository) Save(ctx context.Context, newProject entity.Project) error {
	repository.projects = append(repository.projects, newProject)
	return nil
}

// Update modifies an existing project in the repository.
// Parameters:
// - ctx: The context for managing the request.
// - newProject: The updated project data.
// Returns:
// - An error, if the project could not be updated.
func (repository *ProjectRepository) Update(ctx context.Context, newProject entity.Project) error {
	for i, project := range repository.projects {
		if project.Id == newProject.Id {
			repository.projects[i] = newProject
			return nil
		}
	}

	return ErrProjectNotUpdated
}

// DeleteById removes a project from the repository by its unique identifier.
// Parameters:
// - ctx: The context for managing the request.
// - id: The unique identifier of the project to be deleted.
// Returns:
// - An error, if any, during the deletion process.
func (repository *ProjectRepository) DeleteById(ctx context.Context, id string) error {
	for i, project := range repository.projects {
		if project.Id == id {
			repository.projects = append(repository.projects[:i], repository.projects[i+1:]...)
			return nil
		}
	}

	return nil
}
