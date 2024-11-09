package v1

import (
	"PIS_labs/internal/interface/rest/handlers"
	"github.com/gin-gonic/gin"
)

func MustGetGinRouter(controller *handlers.Handlers) *gin.Engine {
	const (
		ApiV1 = "/api/v1"
		Root  = "/"
		ById  = "/:id"
	)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	const (
		Auth     = "/auth"
		Login    = "/login"
		Register = "/register"
	)

	apiPublicV1Routes := router.Group(ApiV1)
	{
		auth := apiPublicV1Routes.Group(Auth)
		{
			auth.POST(Login, controller.AuthHandler.Login)
			auth.POST(Register, controller.AuthHandler.Register)
		}
	}

	const Projects = "/projects"
	const Tasks = "/tasks"

	apiPrivateV1Routes := router.Group(ApiV1)
	{
		projects := apiPrivateV1Routes.Group(Projects)
		{
			projects.GET(ById, controller.ProjectHandler.GetProjectById)
			projects.POST(Root, controller.ProjectHandler.SaveProject)
			projects.PUT(ById, controller.ProjectHandler.UpdateProject)
			projects.DELETE(ById, controller.ProjectHandler.DeleteProjectById)
		}
		tasks := apiPrivateV1Routes.Group(Tasks)
		{
			tasks.GET(ById, controller.TaskHandler.GetTaskById)
			tasks.POST(Root, controller.TaskHandler.SaveTask)
			tasks.PUT(ById, controller.TaskHandler.UpdateTask)
			tasks.DELETE(ById, controller.TaskHandler.DeleteTaskById)
		}
	}

	return router
}
