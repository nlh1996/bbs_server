package thread

import (
	"bbs_server/model"

	"github.com/gin-gonic/gin"
)

//Publish 发帖请求
func Publish(c *gin.Context) {
	thread := &model.Thread{}
	if err := c.Bind(thread); err != nil{

	}

}
