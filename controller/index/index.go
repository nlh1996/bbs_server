package index

import (
	"bbs_server/model"
	"bbs_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetInfo .
func GetInfo(c *gin.Context) {
	topics := &[]model.Topic{}
	if err := model.GetTopics(topics); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	msg := &model.TodayMsg{}
	msg.Today = utils.GetDateStr()
	msg.AccessSave()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": *topics,
	})
}
