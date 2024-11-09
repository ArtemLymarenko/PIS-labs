package entity

import (
	"time"
)

type Status string

const (
	IDLE   Status = "idle"
	ACTIVE Status = "active"
	CLOSED Status = "closed"
)

type Project struct {
	Id                string
	Name              string
	Description       string
	Status            Status
	ProductionStartAt time.Time
	ProductionEndAt   time.Time
}

func NewProject(
	id, name, description string,
	status Status,
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

func (p *Project) Validate() error {
	return nil
}
