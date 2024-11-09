package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
	"errors"
)

// ErrInvalidProjectEntity is an error returned when a project entity fails validation.
var (
	ErrInvalidProjectEntity = errors.New("failed to validate project entity")
)

// Project represents a service for managing project-related operations, including CRUD actions.
// It uses a repository to perform the actual data operations.
type Project struct {
	projectRepo repository.Project
}

// NewProject creates a new Project service instance with the provided repository.
// Parameters:
// - projectRepo: The repository responsible for handling project data operations.
// Returns:
// - A pointer to the newly created Project service.
func NewProject(projectRepo repository.Project) *Project {
	return &Project{
		projectRepo: projectRepo,
	}
}

// FindProjectById retrieves a project by its unique identifier from the repository.
// Parameters:
// - ctx: The context for managing the lifecycle of the request.
// - id: The unique identifier of the project to be retrieved.
// Returns:
// - The project entity if found, or an error if not found or an issue occurs.
func (service *Project) FindProjectById(ctx context.Context, id string) (entity.Project, error) {
	return service.projectRepo.FindById(ctx, id)
}

// FindProjectsByName retrieves all projects that match the specified name from the repository.
// Parameters:
// - ctx: The context for managing the request.
// - name: The name of the projects to search for.
// Returns:
// - A slice of project entities that match the given name, or an error if the operation fails.
func (service *Project) FindProjectsByName(ctx context.Context, name string) ([]entity.Project, error) {
	return service.projectRepo.FindByNameMany(ctx, name)
}

// SaveProject validates the project entity and saves it to the repository.
// Parameters:
// - ctx: The context for managing the request.
// - project: The project entity to be saved.
// Returns:
// - An error if the project entity is invalid or if saving the project fails.
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

// UpdateProject validates the project entity and updates it in the repository.
// Parameters:
// - ctx: The context for managing the request.
// - project: The project entity to be updated.
// Returns:
// - An error if the project entity is invalid or if updating the project fails.
func (service *Project) UpdateProject(ctx context.Context, project entity.Project) error {
	err := project.Validate()
	if err != nil {
		return ErrInvalidProjectEntity
	}

	return service.projectRepo.Update(ctx, project)
}

// DeleteProjectById deletes a project from the repository using its unique identifier.
// Parameters:
// - ctx: The context for managing the request.
// - id: The unique identifier of the project to be deleted.
// Returns:
// - An error if the deletion fails.
func (service *Project) DeleteProjectById(ctx context.Context, id string) error {
	return service.projectRepo.DeleteById(ctx, id)
}
