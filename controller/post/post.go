package post

import (
	"bbs_server/common"
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
	post := &model.Post{}
	topStorey := &post.TopStorey
	if err := c.Bind(topStorey); err != nil {
		fmt.Println(err.Error())
		return
	}

	now := time.Now()
	topStorey.CreateTime = now.Format("2006-01-02 15:04:05")
	// 将时间格式化为字符串
	rename := now.Format("20060102150405")
	// 帖子的唯一id
	topStorey.TID = topStorey.UID + rename
	// 图片解码，保存至文件服务器
	if len(post.ImgList) != 0 {
		var (
			enc  = base64.StdEncoding	
			path string
		)
		for index, img := range post.ImgList {
			if img[11] == 'j' {
				img = img[23:]
				path = fmt.Sprintf("/img/%s%d.jpg", rename, index)
			} else if img[11] == 'p' {
				img = img[22:]
				path = fmt.Sprintf("/img/%s%d.png", rename, index)
			} else if img[11] == 'g' {
				img = img[22:]
				path = fmt.Sprintf("/img/%s%d.gif", rename, index)
			} else {
				fmt.Println("不支持该文件类型")
			}

			data, err := enc.DecodeString(img)
			if err != nil {
				fmt.Println(err.Error())
			}
			//图片写入文件
			f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer f.Close()
			f.Write(data)
			path = "http://115.159.77.155:12000" + path 
			post.ImgList[index] = path
		}
	}

	//记录贴子更新时间
	post.UpdateTime = topStorey.CreateTime
	post.TID = topStorey.TID
	if(post.Save()) {
		c.String(http.StatusOK, "success")
	}
}

// GetPosts 获取所有贴子
func GetPosts(c *gin.Context) {
	model.UpdatePosts(common.PostsPool)
	topStoreys := []model.TopStorey{}
	for _,value := range *common.PostsPool {
		topStoreys = append(topStoreys, value.TopStorey)
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": topStoreys,
		"msg": "scessue",
	})
}

// GetPost 获取单个贴子详情
func GetPost(c *gin.Context) {
	post := &model.Post{}
	tid := c.Query("tid")
	post.Get(tid)
	c.JSON(http.StatusOK, gin.H{
		"post": *post,
		"msg": "scessue",
	})
}