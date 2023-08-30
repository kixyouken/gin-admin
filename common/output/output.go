package output

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Successful 成功返回
//
//	@param c
//	@param message
func Successful(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
	})
}

// Errorful 错误返回
//
//	@param c
//	@param code
//	@param message
func Errorful(c *gin.Context, code int, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": message,
	})
}

// Dataful 数据返回
//
//	@param c
//	@param data
func Dataful(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    data,
		"message": "success",
	})
}

// Pageful 分页返回
//
//	@param c
//	@param data
//	@param count
func Pageful(c *gin.Context, data interface{}, count int64) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    data,
		"count":   count,
		"message": "success",
	})
}
