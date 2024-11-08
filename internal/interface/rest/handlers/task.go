package handlers

import "github.com/gin-gonic/gin"

type TaskHandlerImpl interface {
	GetTaskById(c *gin.Context)
	SaveTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTaskById(c *gin.Context)
}

//implementation of task handler
