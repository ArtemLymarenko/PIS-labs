package repository

import (
	"PIS_labs/internal/domain/entity"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectRepository_FindById(t *testing.T) {
	ctx := context.TODO()
	projectRepo := &ProjectRepository{
		projects: []entity.Project{
			{Id: "1", Name: "Project A"},
			{Id: "2", Name: "Project B"},
		},
	}

	t.Run("should find project by ID", func(t *testing.T) {
		project, err := projectRepo.FindById(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, "Project A", project.Name)
	})

	t.Run("should return error when project not found", func(t *testing.T) {
		_, err := projectRepo.FindById(ctx, "3")
		assert.Error(t, err)
		assert.Equal(t, ErrProjectNotFound, err)
	})
}

func TestProjectRepository_FindByNameMany(t *testing.T) {
	ctx := context.TODO()
	projectRepo := &ProjectRepository{
		projects: []entity.Project{
			{Id: "1", Name: "Project A"},
			{Id: "2", Name: "Project A"},
			{Id: "3", Name: "Project B"},
		},
	}

	t.Run("should find multiple projects by name", func(t *testing.T) {
		projects, err := projectRepo.FindByNameMany(ctx, "Project A")
		assert.NoError(t, err)
		assert.Len(t, projects, 2)
	})
}

func TestProjectRepository_Save(t *testing.T) {
	ctx := context.TODO()
	projectRepo := &ProjectRepository{
		projects: []entity.Project{},
	}

	t.Run("should save a new project", func(t *testing.T) {
		newProject := entity.Project{Id: "1", Name: "Project A"}
		err := projectRepo.Save(ctx, newProject)
		assert.NoError(t, err)
		assert.Len(t, projectRepo.projects, 1)
		assert.Equal(t, newProject, projectRepo.projects[0])
	})
}

func TestProjectRepository_Update(t *testing.T) {
	ctx := context.TODO()
	projectRepo := &ProjectRepository{
		projects: []entity.Project{
			{Id: "1", Name: "Project A"},
		},
	}

	t.Run("should update existing project", func(t *testing.T) {
		updatedProject := entity.Project{Id: "1", Name: "Updated Project A"}
		err := projectRepo.Update(ctx, updatedProject)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Project A", projectRepo.projects[0].Name)
	})

	t.Run("should return error when updating non-existent project", func(t *testing.T) {
		nonExistentProject := entity.Project{Id: "2", Name: "Non-existent Project"}
		err := projectRepo.Update(ctx, nonExistentProject)
		assert.Error(t, err)
		assert.Equal(t, ErrProjectNotUpdated, err)
	})
}

func TestProjectRepository_DeleteById(t *testing.T) {
	ctx := context.TODO()
	projectRepo := &ProjectRepository{
		projects: []entity.Project{
			{Id: "1", Name: "Project A"},
			{Id: "2", Name: "Project B"},
		},
	}

	t.Run("should delete project by ID", func(t *testing.T) {
		err := projectRepo.DeleteById(ctx, "1")
		assert.NoError(t, err)
		assert.Len(t, projectRepo.projects, 1)
		assert.Equal(t, "Project B", projectRepo.projects[0].Name)
	})

	t.Run("should not return error when deleting non-existent project", func(t *testing.T) {
		err := projectRepo.DeleteById(ctx, "3")
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
}
