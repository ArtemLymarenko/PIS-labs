package handlers

import "github.com/gin-gonic/gin"

type TaskHandlerImpl interface {
	GetTaskById(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTaskById(c *gin.Context)
	DeleteTaskById(c *gin.Context)
}

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

// GetTaskById godoc
//
//	@Summary		Get task by ID
//	@Description	Retrieve a task by its unique ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string				true	"Task ID"
//	@Success		200	{object}	dto.TaskResponse	"Task details"
//	@Failure		404	{object}	dto.ResponseErr		"Task not found"
//	@Failure		500	{object}	dto.ResponseErr		"Internal Server Error"
//	@Router			/tasks/{id} [get]
func (t *TaskHandler) GetTaskById(c *gin.Context) {}

// CreateTask godoc
//
//	@Summary		Create a new task
//	@Description	Create a task with the provided data
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			task	body		dto.CreateTaskRequest	true	"Task data"
//	@Success		201		{object}	dto.TaskResponse		"Task created successfully"
//	@Failure		400		{object}	dto.ResponseErr			"Invalid task data"
//	@Failure		500		{object}	dto.ResponseErr			"Internal Server Error"
//	@Router			/tasks [post]
func (t *TaskHandler) CreateTask(c *gin.Context) {}

// UpdateTaskById godoc
//
//	@Summary		Update an existing task
//	@Description	Update a task's information by its ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			task	body		dto.UpdateTaskRequest	true	"Updated task data"
//	@Success		200		{object}	dto.TaskResponse		"Task updated successfully"
//	@Failure		400		{object}	dto.ResponseErr			"Invalid task data"
//	@Failure		404		{object}	dto.ResponseErr			"Task not found"
//	@Failure		500		{object}	dto.ResponseErr			"Internal Server Error"
//	@Router			/tasks/{id} [put]
func (t *TaskHandler) UpdateTaskById(c *gin.Context) {}

// DeleteTaskById godoc
//
//	@Summary		Delete task by ID
//	@Description	Delete a task using its unique ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string			true	"Task ID"
//	@Success		200	{object}	nil				"Task deleted successfully"
//	@Failure		404	{object}	dto.ResponseErr	"Task not found"
//	@Failure		500	{object}	dto.ResponseErr	"Internal Server Error"
//	@Router			/tasks/{id} [delete]
func (t *TaskHandler) DeleteTaskById(c *gin.Context) {}
