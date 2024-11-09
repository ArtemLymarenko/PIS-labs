package handlers

import "github.com/gin-gonic/gin"

type ProjectHandlerImpl interface {
	GetProjectById(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProjectById(c *gin.Context)
	DeleteProjectById(c *gin.Context)
}

type ProjectHandler struct{}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

// GetProjectById godoc
//
//	@Summary		Getting project by id
//	@Description	get project by ID
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"Project ID"
//	@Success		200	{object}	dto.ProjectResponse
//	@Failure		400	{object}	dto.ResponseErr
//	@Failure		404	{object}	dto.ResponseErr
//	@Failure		500	{object}	dto.ResponseErr
//	@Router			/projects/{id} [get]
func (p *ProjectHandler) GetProjectById(c *gin.Context) {}

// CreateProject godoc
//
//	@Summary		Creates new project by id
//	@Description	post project
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			project	body		dto.CreateProjectRequest	true	"project data"
//	@Success		201		{object}	dto.ProjectResponse
//	@Failure		400		{object}	dto.ResponseErr
//	@Failure		404		{object}	dto.ResponseErr
//	@Failure		500		{object}	dto.ResponseErr
//	@Router			/projects [post]
func (p *ProjectHandler) CreateProject(c *gin.Context) {}

// UpdateProjectById godoc
//
//	@Summary		Updates project by id
//	@Description	put project
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int							true	"Project ID"
//	@Param			project	body		dto.UpdateProjectRequest	true	"project data"
//	@Success		200		{object}	dto.ProjectResponse
//	@Failure		400		{object}	dto.ResponseErr
//	@Failure		404		{object}	dto.ResponseErr
//	@Failure		500		{object}	dto.ResponseErr
//	@Router			/projects/{id} [put]
func (p *ProjectHandler) UpdateProjectById(c *gin.Context) {}

// DeleteProjectById godoc
//
//	@Summary		Deletes project by id
//	@Description	delete project
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"Project ID"
//	@Success		200	{object}	nil	"Successful response"
//	@Failure		400	{object}	dto.ResponseErr
//	@Failure		404	{object}	dto.ResponseErr
//	@Failure		500	{object}	dto.ResponseErr
//	@Router			/projects/{id} [delete]
func (p *ProjectHandler) DeleteProjectById(c *gin.Context) {}
