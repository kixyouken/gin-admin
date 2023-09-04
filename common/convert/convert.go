package convert

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetInt 获取Int参数
//
//	@param c
//	@param param
//	@return int
func GetInt(c *gin.Context, param string) int {
	paramStr := c.Param(param)
	paramInt, _ := strconv.Atoi(paramStr)
	return paramInt
}

// GetUint 获取Uint参数
//
//	@param c
//	@param param
//	@return uint
func GetUint(c *gin.Context, param string) uint {
	return uint(GetInt(c, param))
}
