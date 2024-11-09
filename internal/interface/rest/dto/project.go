package dto

import "time"

type CreateProjectRequest struct {
	Name              string    `json:"name"`
	Description       string    `json:"description,omitempty"`
	Status            string    `json:"status"`
	ProductionStartAt time.Time `json:"productionStartAt"`
	ProductionEndAt   time.Time `json:"productionEndAt"`
}

type UpdateProjectRequest struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description,omitempty"`
	Status            string    `json:"status"`
	ProductionStartAt time.Time `json:"productionStartAt"`
	ProductionEndAt   time.Time `json:"productionEndAt"`
}

type ProjectResponse struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description,omitempty"`
	Status            string    `json:"status"`
	ProductionStartAt time.Time `json:"productionStartAt"`
	ProductionEndAt   time.Time `json:"productionEndAt"`
}
