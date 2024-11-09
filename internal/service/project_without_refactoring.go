package service

import (
	"PIS_labs/internal/domain/entity"
	"PIS_labs/internal/domain/repository"
	"context"
	"errors"
)

// Naming is bad
type Proj struct {
	//Here is using specific implementation instead of interface
	repo repository.ProjectRepository
}

func NewProj(repo repository.ProjectRepository) *Proj {
	return &Proj{
		repo: repo,
	}
}

// Naming is bad
func (s *Proj) Find(ctx context.Context, id string) (entity.Project, error) {
	//Unnecessary validation, better to move this logic to domain
	if id == "" {
		//Inline errors is bad, better to move to separate package variable
		return entity.Project{}, errors.New("ID cannot be empty")
	}

	project, err := s.repo.FindById(ctx, id)
	if err != nil {
		//Inline errors is bad, better to move to separate package variable
		return entity.Project{}, errors.New("project with the given ID does not exist")
	}
	return project, nil
}

// Naming is bad
func (s *Proj) FindName(ctx context.Context, name string) ([]entity.Project, error) {
	var foundProjects []entity.Project
	projects, err := s.repo.FindByNameMany(ctx, name)
	if err != nil {
		//Inline errors is bad, better to move to separate package variable
		return nil, errors.New("error occurred while searching projects")
	}
	//Unnecessary logic
	if len(projects) > 0 {
		for _, project := range projects {
			if project.Name == name {
				foundProjects = append(foundProjects, project)
			}
		}
	}
	return foundProjects, nil
}

func (s *Proj) Save(ctx context.Context, project entity.Project) error {
	//Unnecessary validation, better to move this logic to domain
	if project.Name == "" {
		return errors.New("project name cannot be empty")
	}

	if project.Id == "" {
		return errors.New("project ID cannot be empty")
	}

	return s.repo.Save(ctx, project)
}

func (s *Proj) Update(ctx context.Context, project entity.Project) error {
	if project.Name == "" {
		//Inline errors is bad, better to move to separate package variable
		return errors.New("project name cannot be empty")
	}

	if project.Id == "" {
		//Inline errors is bad, better to move to separate package variable
		return errors.New("project ID cannot be empty")
	}

	return s.repo.Update(ctx, project)
}

// Naming is bad
func (s *Proj) Delete(ctx context.Context, id string) error {
	//Unnecessary validation
	if id == "" {
		//Inline errors is bad, better to move to separate package variable
		return errors.New("ID cannot be empty")
	}

	err := s.repo.DeleteById(ctx, id)
	if err != nil {
		//Inline errors is bad, better to move to separate package variable
		return errors.New("error deleting project by ID")
	}

	return nil
}
