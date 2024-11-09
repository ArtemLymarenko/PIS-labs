package dto

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ProjectId   string `json:"projectId"`
}

type UpdateTaskRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ProjectId   string `json:"projectId"`
}

type TaskResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ProjectId   string `json:"projectId"`
}
