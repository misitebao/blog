package controller

import (
	"github.com/gin-gonic/gin"
)

type PageController struct{}

func (PageController) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Main website",
	})
}