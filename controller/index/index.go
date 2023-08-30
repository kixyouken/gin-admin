package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	menuTree, _ := c.Get("menuTree")
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Main website",
		"menu":  menuTree,
	})
}
