package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Main website",
	})
}

func Read(c *gin.Context) {
	c.HTML(http.StatusOK, "read", gin.H{
		"title": "Main website",
	})
}
