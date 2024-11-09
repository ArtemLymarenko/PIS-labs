package entity

type Task struct {
	Id          string
	Name        string
	Description string
	ProjectId   string
}

func NewTask(id, name, description, projectId string) Task {
	return Task{
		Id:          id,
		Name:        name,
		Description: description,
		ProjectId:   projectId,
	}
}

func (t *Task) Validate() error {
	return nil
}
