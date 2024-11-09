package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/mock"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProject_FindById(t *testing.T) {
	ctx := context.TODO()
	mockRepo := mock.NewProjectRepository()
	projectService := NewProject(mockRepo)

	t.Run("should find project by ID", func(t *testing.T) {
		expectedProject := entity.Project{Id: "1", Name: "Project A"}
		mockRepo.On("FindById", ctx, "1").Return(expectedProject, nil)

		project, err := projectService.FindProjectById(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, expectedProject, project)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when project not found", func(t *testing.T) {
		mockRepo.On("FindById", ctx, "2").Return(entity.Project{}, errors.New("project not found"))

		_, err := projectService.FindProjectById(ctx, "2")
		assert.Error(t, err)
		assert.Equal(t, "project not found", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestProject_FindByNameMany(t *testing.T) {
	ctx := context.TODO()
	mockRepo := mock.NewProjectRepository()
	projectService := NewProject(mockRepo)

	t.Run("should find multiple projects by name", func(t *testing.T) {
		expectedProjects := []entity.Project{
			{Id: "1", Name: "Project A"},
			{Id: "2", Name: "Project A"},
		}
		mockRepo.On("FindByNameMany", ctx, "Project A").Return(expectedProjects, nil)

		projects, err := projectService.FindProjectsByName(ctx, "Project A")
		assert.NoError(t, err)
		assert.Len(t, projects, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when no projects found", func(t *testing.T) {
		mockRepo.On("FindByNameMany", ctx, "Project C").Return([]entity.Project{}, errors.New("no projects found"))

		_, err := projectService.FindProjectsByName(ctx, "Project C")
		assert.Error(t, err)
		assert.Equal(t, "no projects found", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestProject_Save(t *testing.T) {
	ctx := context.TODO()
	mockRepo := mock.NewProjectRepository()
	projectService := NewProject(mockRepo)

	t.Run("should save project successfully", func(t *testing.T) {
		newProject := entity.Project{Id: "1", Name: "Project A"}
		mockRepo.On("Save", ctx, newProject).Return(nil)

		err := projectService.SaveProject(ctx, newProject)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestProject_Update(t *testing.T) {
	ctx := context.TODO()
	mockRepo := mock.NewProjectRepository()
	projectService := NewProject(mockRepo)

	t.Run("should update project successfully", func(t *testing.T) {
		project := entity.Project{Id: "1", Name: "Updated Project A"}
		mockRepo.On("Update", ctx, project).Return(nil)

		err := projectService.UpdateProject(ctx, project)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestProject_DeleteById(t *testing.T) {
	ctx := context.TODO()
	mockRepo := mock.NewProjectRepository()
	projectService := NewProject(mockRepo)

	t.Run("should delete project by ID", func(t *testing.T) {
		mockRepo.On("DeleteById", ctx, "1").Return(nil)

		err := projectService.DeleteProjectById(ctx, "1")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when deleting non-existent project", func(t *testing.T) {
		mockRepo.On("DeleteById", ctx, "2").Return(errors.New("project not found"))

		err := projectService.DeleteProjectById(ctx, "2")
		assert.Error(t, err)
		assert.Equal(t, "project not found", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
