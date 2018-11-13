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

// Publish 发帖请求
func Publish(c *gin.Context) {
	thread := &model.Thread{}
	topStorey := &thread.TopStorey
	if err := c.Bind(topStorey); err != nil {
		fmt.Println(err.Error())
		return
	}

	topStorey.CreateTime = time.Now()
	// 将时间格式化为字符串
	rename := topStorey.CreateTime.Format("20060102150405")

	// 图片解码，保存至文件服务器
	if len(thread.ImgList) != 0 {
		var (
			enc  = base64.StdEncoding
			path string
		)
		for index, img := range thread.ImgList {
			if img[11] == 'j' {
				img = img[23:]
				path = fmt.Sprintf("D://image/%s%d.jpg", rename, index)
			} else if img[11] == 'p' {
				img = img[22:]
				path = fmt.Sprintf("D://image/%s%d.png", rename, index)
			} else if img[11] == 'g' {
				img = img[22:]
				path = fmt.Sprintf("D://image/%s%d.gif", rename, index)
			} else {
				fmt.Println("不支持该文件类型")
			}

			data, err := enc.DecodeString(img)
			if err != nil {
				fmt.Println(err.Error())
			}
			//写入新文件
			f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer f.Close()
			f.Write(data)
			thread.ImgList[index] = path
		}
	}

	thread.Save()

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": *thread,
	})
}
