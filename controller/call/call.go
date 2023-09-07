package call

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	menuTree, _ := c.Get("menuTree")
	c.HTML(http.StatusOK, "call/index", gin.H{
		"title": "会員管理",
		"menu":  menuTree,
		"uri":   "call/list",
	})
}
