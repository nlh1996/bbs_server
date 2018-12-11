package feedback

import (
	"bbs_server/model"
	"bbs_server/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Complaint 用户反馈
func Complaint(c *gin.Context) {
	msg := &model.Complaint{}
	msg.CreateTime = utils.GetTimeStr()
	if err := c.Bind(msg); err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, "内部出错")
	}
	msg.Save()
	c.String(http.StatusOK, "success")
}
