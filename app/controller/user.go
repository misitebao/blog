package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "info",
	})
}
