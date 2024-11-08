package handlers

import "github.com/gin-gonic/gin"

type ProjectHandlerImpl interface {
	GetProjectById(c *gin.Context)
	SaveProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProjectById(c *gin.Context)
}

//implementation of project handler
