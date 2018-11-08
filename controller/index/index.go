package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetInfo .
func GetInfo(c *gin.Context) {
	c.String(http.StatusOK, "首页加载成功")
}
