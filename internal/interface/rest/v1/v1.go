package v1

import (
	"PIS_labs/internal/interface/rest/handlers"
	"github.com/gin-gonic/gin"
)

func MustGetGinRouter(controller *handlers.Handlers) *gin.Engine {
	const ApiV1 = "/api/v1"

	const (
		Root = "/"
		ById = "/:id"
	)

	const Projects = "/projects"

	const Tasks = "/tasks"

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

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
