package entity

import (
	"time"
)

type ProjectStatus string

const (
	Idle   ProjectStatus = "idle"
	Active ProjectStatus = "active"
	Closed ProjectStatus = "closed"
)

type Project struct {
	Id                string
	Name              string
	Description       string
	Status            ProjectStatus
	ProductionStartAt time.Time
	ProductionEndAt   time.Time
}

func NewProject(
	id, name, description string,
	status ProjectStatus,
	productionStartAt, productionEndAt time.Time,
) Project {
	return Project{
		Id:                id,
		Name:              name,
		Description:       description,
		Status:            status,
		ProductionStartAt: productionStartAt,
		ProductionEndAt:   productionEndAt,
	}
}

func (project *Project) SetId(id string) {
	project.Id = id
}

func (project *Project) Validate() error {
	return nil
}
