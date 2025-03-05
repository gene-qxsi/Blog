package http

import "github.com/gin-gonic/gin"

type UserHandlers interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}
