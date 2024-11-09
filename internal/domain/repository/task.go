package repository

import (
	"PIS_labs/internal/domain/entity"
	"context"
	"errors"
)

// Task defines the interface for managing task data in the repository layer.
// It includes methods for finding, saving, updating, and deleting tasks.
type Task interface {
	FindById(ctx context.Context, id string) (entity.Task, error)
	FindByNameMany(ctx context.Context, name string) ([]entity.Task, error)
	Save(ctx context.Context, task entity.Task) error
	Update(ctx context.Context, task entity.Task) error
	DeleteById(ctx context.Context, id string) error
}

// TaskRepository is an implementation of the Task interface.
// It holds a slice of tasks and provides methods to perform CRUD operations.
type TaskRepository struct {
	tasks []entity.Task
}

// FindById searches for a task by its unique identifier.
// Parameters:
// - ctx: The context for managing the request.
// - id: The unique identifier of the task.
// Returns:
// - The found task, if any, or an error if the task is not found.
func (t *TaskRepository) FindById(ctx context.Context, id string) (entity.Task, error) {
	for _, task := range t.tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return entity.Task{}, errors.New("task not found")
}

// FindByNameMany searches for tasks that match the specified name.
// Parameters:
// - ctx: The context for managing the request.
// - name: The name of the tasks to search for.
// Returns:
// - A slice of matching tasks, or an error if no tasks match.
func (t *TaskRepository) FindByNameMany(ctx context.Context, name string) ([]entity.Task, error) {
	var foundTasks []entity.Task
	for _, task := range t.tasks {
		if task.Name == name {
			foundTasks = append(foundTasks, task)
		}
	}
	if len(foundTasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return foundTasks, nil
}

// Save adds a new task to the repository.
// Parameters:
// - ctx: The context for managing the request.
// - task: The task to be saved.
// Returns:
// - An error, if any, during the saving process.
func (t *TaskRepository) Save(ctx context.Context, task entity.Task) error {
	t.tasks = append(t.tasks, task)
	return nil
}

// Update modifies an existing task in the repository.
// Parameters:
// - ctx: The context for managing the request.
// - task: The updated task data.
// Returns:
// - An error, if the task could not be updated.
func (t *TaskRepository) Update(ctx context.Context, task entity.Task) error {
	for i, tsk := range t.tasks {
		if tsk.Id == task.Id {
			t.tasks[i] = task
			return nil
		}
	}
	return errors.New("task not found")
}

// DeleteById removes a task from the repository by its unique identifier.
// Parameters:
// - ctx: The context for managing the request.
// - id: The unique identifier of the task to be deleted.
// Returns:
// - An error, if any, during the deletion process.
func (t *TaskRepository) DeleteById(ctx context.Context, id string) error {
	for i, task := range t.tasks {
		if task.Id == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
