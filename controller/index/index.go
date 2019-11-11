package index

import (
	"bbs_server/model"
	"bbs_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetInfo .
func GetInfo(c *gin.Context) {
	var topic = &model.Topic{}
	topic.TopicNum = 100
	topic.TopicImg = "https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2564997198,4187947589&fm=58"
	msg := &model.TodayMsg{}
	msg.Today = utils.GetDateStr()
	msg.AccessSave()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": *topic,
	})
}

// GetGames .
func GetGames(c *gin.Context) {
	games := &[]model.Game{}
	if err := model.GetGames(games); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"games": *games,
	})
}
