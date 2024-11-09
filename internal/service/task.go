package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
	"errors"
)

// ErrInvalidTaskEntity is an error returned when a task entity fails validation.
var (
	ErrInvalidTaskEntity = errors.New("failed to validate task entity")
)

// Task represents a service for managing task-related operations, including CRUD actions.
// It interacts with a task repository to perform the actual data operations.
type Task struct {
	taskRepository repository.Task
}

// NewTask creates a new Task service instance with the provided task repository.
// Parameters:
// - taskRepository: The repository responsible for handling task data operations.
// Returns:
// - A pointer to the newly created Task service.
func NewTask(taskRepository repository.Task) *Task {
	return &Task{
		taskRepository: taskRepository,
	}
}

// FindById retrieves a task by its unique identifier from the repository.
// Parameters:
// - ctx: The context for managing the lifecycle of the request.
// - id: The unique identifier of the task to be retrieved.
// Returns:
// - The task entity if found, or an error if not found or an issue occurs.
func (service *Task) FindById(ctx context.Context, id string) (entity.Task, error) {
	return service.taskRepository.FindById(ctx, id)
}

// FindByNameMany retrieves tasks that match the specified name from the repository.
// Parameters:
// - ctx: The context for managing the request.
// - name: The name of the tasks to search for.
// Returns:
// - A slice of task entities that match the given name, or an error if the operation fails.
func (service *Task) FindByNameMany(ctx context.Context, name string) ([]entity.Task, error) {
	return service.taskRepository.FindByNameMany(ctx, name)
}

// Save validates the task entity and saves it to the repository.
// Parameters:
// - ctx: The context for managing the request.
// - task: The task entity to be saved.
// Returns:
// - An error if the task entity is invalid or if saving the task fails.
func (service *Task) Save(ctx context.Context, task entity.Task) error {
	err := task.Validate()
	if err != nil {
		return ErrInvalidTaskEntity
	}

	return service.taskRepository.Save(ctx, task)
}

// UpdateById validates the task entity and updates it in the repository.
// Parameters:
// - ctx: The context for managing the request.
// - task: The task entity to be updated.
// Returns:
// - An error if the task entity is invalid or if updating the task fails.
func (service *Task) UpdateById(ctx context.Context, task entity.Task) error {
	err := task.Validate()
	if err != nil {
		return ErrInvalidTaskEntity
	}

	return service.taskRepository.Update(ctx, task)
}

// DeleteById deletes a task from the repository using its unique identifier.
// Parameters:
// - ctx: The context for managing the request.
// - id: The unique identifier of the task to be deleted.
// Returns:
// - An error if the deletion fails.
func (service *Task) DeleteById(ctx context.Context, id string) error {
	return service.taskRepository.DeleteById(ctx, id)
}
