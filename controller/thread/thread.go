package thread

import (
	"bbs_server/model"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

//Publish 发帖请求
func Publish(c *gin.Context) {
	thread := &model.TopStorey{}
	if err := c.Bind(thread); err != nil {
		fmt.Println(err.Error())
		return
	}

	thread.CurrentTime = time.Now()
	
	//图片解码，保存至文件服务器
	if len(thread.ImgList)!= 0 {
		var enc = base64.StdEncoding
		for index,img := range thread.ImgList {
			str := img[23:]
			data, err := enc.DecodeString(str)
			if err != nil {
				fmt.Println(err.Error())
			}

			path := fmt.Sprintf("D://image/%d.jpg",index)
			//写入新文件
			f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer f.Close()
			f.Write(data)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": *thread,
	})
}
