package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
	"errors"
)

var (
	ErrInvalidTaskEntity = errors.New("failed to validate task entity")
)

type Task struct {
	taskRepository repository.Task
}

func NewTask(taskRepository repository.Task) *Task {
	return &Task{
		taskRepository: taskRepository,
	}
}

func (service *Task) FindById(ctx context.Context, id string) (entity.Task, error) {
	return service.taskRepository.FindById(ctx, id)
}

func (service *Task) FindByNameMany(ctx context.Context, name string) ([]entity.Task, error) {
	return service.taskRepository.FindByNameMany(ctx, name)
}

func (service *Task) Save(ctx context.Context, task entity.Task) error {
	err := task.Validate()
	if err != nil {
		return ErrInvalidTaskEntity
	}

	return service.taskRepository.Save(ctx, task)
}

func (service *Task) Update(ctx context.Context, task entity.Task) error {
	err := task.Validate()
	if err != nil {
		return ErrInvalidTaskEntity
	}

	return service.taskRepository.Update(ctx, task)
}

func (service *Task) DeleteById(ctx context.Context, id string) error {
	return service.taskRepository.DeleteById(ctx, id)
}
