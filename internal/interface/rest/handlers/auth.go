package handlers

import "github.com/gin-gonic/gin"

type AuthHandlerImpl interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
