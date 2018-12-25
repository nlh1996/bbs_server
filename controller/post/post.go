package post

import (
	"bbs_server/common"
	"bbs_server/model"
	"bbs_server/utils"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Publish 发帖请求
func Publish(c *gin.Context) {
	post := &model.Post{}
	topStorey := &post.TopStorey
	if err := c.Bind(topStorey); err != nil {
		fmt.Println(err.Error())
		return
	}

	// 贴子创建时间
	topStorey.CreateTime = utils.GetTimeStr()
	// 帖子的唯一id
	topStorey.TID = bson.NewObjectId()
	// 图片解码，保存至文件服务器
	if len(topStorey.ImgList) != 0 {
		var (
			enc  = base64.StdEncoding
			path string
		)
		for index, img := range topStorey.ImgList {
			if img[11] == 'j' {
				img = img[23:]
				path = fmt.Sprintf("d:/image/%x%d.jpg", string(topStorey.TID), index)
			} else if img[11] == 'p' {
				img = img[22:]
				path = fmt.Sprintf("d:/image/%x%d.png", string(topStorey.TID), index)
			} else if img[11] == 'g' {
				img = img[22:]
				path = fmt.Sprintf("d:/image/%x%d.gif", string(topStorey.TID), index)
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
			//记录图片保存的地址
			path = "http://115.159.77.155:12000" + path
			topStorey.ImgList[index] = path
		}
	}

	//记录贴子更新时间
	post.UpdateTime = topStorey.CreateTime
	//记录帖子的唯一id
	post.TID = topStorey.TID
	//将贴子保存到数据库
	if post.Save() {
		c.String(http.StatusOK, "success")
	}else{
		c.String(http.StatusOK, "未能成功保存")
	}
}

// GetPosts 获取所有贴子
func GetPosts(c *gin.Context) {
	model.UpdatePosts(common.PostsPool)
	topStoreys := []model.TopStorey{}
	for _, value := range *common.PostsPool {
		topStoreys = append(topStoreys, value.TopStorey)
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": topStoreys,
		"msg":   "scessue",
	})
}

// GetPost 获取单个贴子详情
func GetPost(c *gin.Context) {
	post := &model.Post{}
	tid := c.Query("tid")
	if post.Get(bson.ObjectIdHex(tid)) {
		c.JSON(http.StatusOK, gin.H{
			"post": *post,
			"msg":  "scessue",
		})
	} else {
		c.String(http.StatusOK, "error,未正确获取到贴子!")
	}
}

// Reply1 一级回复
func Reply1(c *gin.Context) {
	reply1 := &model.Reply1{}
	if err := c.Bind(reply1); err != nil {
		fmt.Println(err.Error())
		return
	}
	reply1.UName = c.Request.Header["Authorization"][0]
	reply1.CreateTime = utils.GetTimeStr()
	reply1.ID = bson.NewObjectId()
	if reply1.Save(reply1.TID) {
		c.JSON(http.StatusOK, gin.H{
			"reply": *reply1,
		})
	} else {
		c.String(http.StatusOK, "内部错误")
	}
}

// Reply2 二级回复
func Reply2(c *gin.Context) {
	reply2 := &model.Reply2{}
	if err := c.Bind(reply2); err != nil {
		fmt.Println(err.Error())
		return
	}
	reply2.UName = c.Request.Header["Authorization"][0]
	reply2.CreateTime = utils.GetTimeStr()
	reply2.ID = bson.NewObjectId()
	if reply2.Save(reply2.TID) {
		c.JSON(http.StatusOK, gin.H{
			"reply": *reply2,
		})
	} else {
		c.String(http.StatusOK, "内部错误")
	}
}

// DelPost 删除贴子
func DelPost(c *gin.Context) {
	name := c.Request.Header["Authorization"][0]
	post := &model.Post{}
	if err := c.Bind(post); err != nil {
		fmt.Println(err.Error())
		return
	}
	if post.Del(post.TID,name) {
		c.String(http.StatusOK, "删除成功")
	} else {
		c.String(http.StatusOK, "删除失败")
	}
}

// Support 点赞行为
func Support(c *gin.Context) {
	//点赞用户，记录点赞的帖子
	name := &model.User{}
	post := &model.Post{}
	name.UName = c.Request.Header["Authorization"][0]
	tid := c.PostForm("tid")
	name.SaveSupport(tid)
	//被点赞用户，增加点赞数
	name.UName = c.PostForm("name")
	name.AddSupport()
	//被点赞贴子，增加点赞数
	post.TID = bson.ObjectIdHex(tid)
	post.AddSupport()
	c.String(http.StatusOK,"")
}

// Cancel 取消赞
func Cancel(c *gin.Context) {
	//点赞用户，删除点赞的帖子
	name := &model.User{}
	post := &model.Post{}
	name.UName = c.Request.Header["Authorization"][0]
	tid := c.PostForm("tid")
	name.CancelSupport(tid)
	//被点赞用户，减少点赞数
	name.UName = c.PostForm("name")
	name.ReduceSupport()
	//被点赞贴子，减少点赞数
	post.TID = bson.ObjectIdHex(tid)
	post.ReduceSupport()
}

