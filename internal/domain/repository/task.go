package repository

import (
	"PIS_labs/internal/domain/entity"
	"context"
	"errors"
)

type Task interface {
	FindById(ctx context.Context, id string) (entity.Task, error)
	FindByNameMany(ctx context.Context, name string) ([]entity.Task, error)
	Save(ctx context.Context, task entity.Task) error
	Update(ctx context.Context, task entity.Task) error
	DeleteById(ctx context.Context, id string) error
}

type TaskRepository struct {
	tasks []entity.Task
}

func (t *TaskRepository) FindById(ctx context.Context, id string) (entity.Task, error) {
	for _, task := range t.tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return entity.Task{}, errors.New("task not found")
}

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

func (t *TaskRepository) Save(ctx context.Context, task entity.Task) error {
	t.tasks = append(t.tasks, task)
	return nil
}

func (t *TaskRepository) Update(ctx context.Context, task entity.Task) error {
	for i, tsk := range t.tasks {
		if tsk.Id == task.Id {
			t.tasks[i] = task
			return nil
		}
	}
	return errors.New("task not found")
}

func (t *TaskRepository) DeleteById(ctx context.Context, id string) error {
	for i, task := range t.tasks {
		if task.Id == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
