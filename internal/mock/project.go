package mock

import (
	"PIS_labs/internal/domain/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type ProjectRepository struct {
	mock.Mock
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{mock.Mock{}}
}

func (m *ProjectRepository) FindById(ctx context.Context, id string) (entity.Project, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.Project), args.Error(1)
}

func (m *ProjectRepository) FindByNameMany(ctx context.Context, name string) ([]entity.Project, error) {
	args := m.Called(ctx, name)
	return args.Get(0).([]entity.Project), args.Error(1)
}

func (m *ProjectRepository) Save(ctx context.Context, project entity.Project) error {
	args := m.Called(ctx, project)
	return args.Error(0)
}

func (m *ProjectRepository) Update(ctx context.Context, project entity.Project) error {
	args := m.Called(ctx, project)
	return args.Error(0)
}

func (m *ProjectRepository) DeleteById(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
