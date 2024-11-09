package main

import (
	_ "PIS_labs/docs"
	"PIS_labs/internal/app"
	"PIS_labs/internal/interface/rest/handlers"
	"PIS_labs/internal/interface/rest/v1"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"net/http"
)

//	@title			Project Management System API
//	@version		1.0
//	@description	This is a sample server Project Management System server.
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.bearer	BearerAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	controller := handlers.New()
	router := v1.MustGetGinRouter(controller)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	application := app.New(server)
	application.Start()
}
