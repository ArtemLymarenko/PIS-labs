package repository

import (
	"PIS_labs/internal/domain/entity"
	"context"
	"errors"
)

type Project interface {
	FindById(ctx context.Context, id string) (entity.Project, error)
	FindByNameMany(ctx context.Context, name string) ([]entity.Project, error)
	Save(ctx context.Context, project entity.Project) error
	Update(ctx context.Context, project entity.Project) error
	DeleteById(ctx context.Context, id string) error
}

type ProjectRepository struct {
	projects []entity.Project
}

func (p *ProjectRepository) FindById(ctx context.Context, id string) (entity.Project, error) {
	for _, project := range p.projects {
		if project.Id == id {
			return project, nil
		}
	}
	return entity.Project{}, errors.New("project not found")
}

func (p *ProjectRepository) FindByNameMany(ctx context.Context, name string) ([]entity.Project, error) {
	var foundProjects []entity.Project
	for _, project := range p.projects {
		if project.Name == name {
			foundProjects = append(foundProjects, project)
		}
	}
	if len(foundProjects) == 0 {
		return nil, errors.New("no projects found")
	}
	return foundProjects, nil
}

func (p *ProjectRepository) Save(ctx context.Context, project entity.Project) error {
	p.projects = append(p.projects, project)
	return nil
}

func (p *ProjectRepository) Update(ctx context.Context, project entity.Project) error {
	// Оновлюємо проект за ID
	for i, v := range p.projects {
		if v.Id == project.Id {
			p.projects[i] = project
			return nil
		}
	}
	return errors.New("project not found")
}

func (p *ProjectRepository) DeleteById(ctx context.Context, id string) error {
	for i, project := range p.projects {
		if project.Id == id {
			p.projects = append(p.projects[:i], p.projects[i+1:]...)
			return nil
		}
	}
	return errors.New("project not found")
}
