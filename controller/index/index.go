package index

import (
	"bbs_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetInfo .
func GetInfo(c *gin.Context) {
	var topic = &model.Topic{}
	topic.TopicNum = 100
	topic.TopicImg = "https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2564997198,4187947589&fm=58"
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": *topic,
	})
}
