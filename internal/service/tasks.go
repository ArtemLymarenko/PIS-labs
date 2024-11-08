package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
)

type Task struct {
	repo repository.TaskRepository
}

func NewTask(repo repository.TaskRepository) *Task {
	return &Task{
		repo: repo,
	}
}

func (s *Task) FindById(ctx context.Context, id string) (entity.Task, error) {
	return s.repo.FindById(ctx, id)
}

func (s *Task) FindByNameMany(ctx context.Context, name string) ([]entity.Task, error) {
	return s.repo.FindByNameMany(ctx, name)
}

func (s *Task) Save(ctx context.Context, task entity.Task) error {
	err := task.Validate()
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, task)
}

func (s *Task) Update(ctx context.Context, task entity.Task) error {
	err := task.Validate()
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, task)
}

func (s *Task) DeleteById(ctx context.Context, id string) error {
	return s.repo.DeleteById(ctx, id)
}
