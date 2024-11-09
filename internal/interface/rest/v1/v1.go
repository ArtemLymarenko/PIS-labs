package v1

import (
	"PIS_labs/internal/interface/rest/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	const (
		Docs    = "/docs"
		Swagger = "/swagger/*any"
	)

	apiPublicV1Routes := router.Group(ApiV1)
	{
		auth := apiPublicV1Routes.Group(Auth)
		{
			auth.POST(Login, controller.AuthHandler.Login)
			auth.POST(Register, controller.AuthHandler.Register)
		}

		docs := apiPublicV1Routes.Group(Docs)
		{
			docs.GET(Swagger, ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	const Projects = "/projects"
	const Tasks = "/tasks"

	apiPrivateV1Routes := router.Group(ApiV1)
	{
		projects := apiPrivateV1Routes.Group(Projects)
		{
			projects.GET(ById, controller.ProjectHandler.GetProjectById)
			projects.POST(Root, controller.ProjectHandler.CreateProject)
			projects.PUT(ById, controller.ProjectHandler.UpdateProjectById)
			projects.DELETE(ById, controller.ProjectHandler.DeleteProjectById)
		}
		tasks := apiPrivateV1Routes.Group(Tasks)
		{
			tasks.GET(ById, controller.TaskHandler.GetTaskById)
			tasks.POST(Root, controller.TaskHandler.CreateTask)
			tasks.PUT(ById, controller.TaskHandler.UpdateTaskById)
			tasks.DELETE(ById, controller.TaskHandler.DeleteTaskById)
		}
	}

	return router
}
