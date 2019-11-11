package admin

import (
	"bbs_server/common"
	"bbs_server/model"
	"bbs_server/utils"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Login 管理员登录
func Login(c *gin.Context) {
	admin := &model.Admin{}
	err := c.Bind(admin)
	if err != nil {
		fmt.Println(err)
	}
	newPwd := utils.Jiami(&admin.PassWord, &admin.UName)
	admin.PassWord = newPwd
	msg, result := admin.Validator()
	if result {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": admin.UName,
		})
		tokenString, err := token.SignedString([]byte("321"))
		if err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusOK, "内部错误")
			return
		}
		//将用户token以键值对的方式加入map缓存中
		common.TokenMap[tokenString] = admin.UName

		c.JSON(http.StatusOK, gin.H{
			"user":   admin,
			"isLoad": 2,
			"token":  tokenString,
			"msg":    msg,
		})
		return
	}
	c.String(http.StatusForbidden, msg)
}

// Count 用户统计。
func Count(c *gin.Context) {
	msg := &model.TodayMsg{}
	msg.Today = utils.GetDateStr()
	msg = msg.Search()
	num := msg.Count()
	c.JSON(http.StatusOK, gin.H{
		"count":   *msg,
		"userNum": num,
	})
}

// UserSearch 用户搜索。
func UserSearch(c *gin.Context) {
	user := &model.User{}
	user.UName = c.PostForm("name")
	var result bool
	result, user = user.Search()
	if result {
		c.JSON(http.StatusOK, gin.H{
			"name":  user.UName,
			"level": user.Exp,
			"jifen": user.Integral,
			"time":  user.CreateTime,
		})
	} else {
		c.String(http.StatusNoContent, "没有该用户！")
	}
}

// AddBlackList 加入黑名单
func AddBlackList(c *gin.Context) {
	user := &model.BlackName{}
	user.UName = c.PostForm("name")
	user.Time = utils.GetTimeStr()
	result := user.BlackNameSave()
	if result == true {
		c.String(http.StatusOK, "scess")
		common.BlackList = user.BlackList()
	} else {
		c.String(http.StatusAccepted, "用户已经拉黑")
	}
}

// RemoveBlackList 移出黑名单
func RemoveBlackList(c *gin.Context) {
	user := &model.BlackName{}
	user.UName = c.PostForm("name")
	result := user.BlackNameRemove()
	if result == true {
		c.String(http.StatusOK, "删除成功！")
		common.BlackList = user.BlackList()
	} else {
		c.String(http.StatusOK, "内部出错！")
	}
}

// GetBlackList 获取黑名单
func GetBlackList(c *gin.Context) {
	user := &model.BlackName{}
	list := user.BlackList()
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}

// GetFeedList0 获取未处理的用户反馈信息
func GetFeedList0(c *gin.Context) {
	msg := &model.Complaint{}
	list := msg.FeedList0()
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}

// GetFeedList1 获取未处理的用户反馈信息
func GetFeedList1(c *gin.Context) {
	msg := &model.Complaint{}
	list := msg.FeedList1()
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}

// DelFeedBack 删除用户反馈信息
func DelFeedBack(c *gin.Context) {
	tid := c.PostForm("tid")
	feedBack := &model.Complaint{}
	result := feedBack.Del(tid)
	if result {
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusNoContent, "")
	}
}

// AgreeFeedBack 同意用户反馈信息
func AgreeFeedBack(c *gin.Context) {
	tid := c.PostForm("tid")
	feedBack := &model.Complaint{}
	result := feedBack.Agree(tid)
	if result {
		post := &model.Post{}
		post.Del(bson.ObjectIdHex(tid), "admin")
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusNoContent, "")
	}
}

// ZhiDing 同意贴子置顶
func ZhiDing(c *gin.Context) {
	tid := c.PostForm("tid")
	post := &model.Post{}
	result := post.AgreeZhiDIng(tid)
	if result {
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusNoContent, "")
	}
}

// AddNotice 添加公告
func AddNotice(c *gin.Context) {
	notice := &model.Notice{}
	notice.Createtime = utils.GetTimeStr()
	if err := c.Bind(notice); err != nil {
		log.Println(err.Error())
		return
	}
	result := notice.Save()
	if result {
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusNoContent, "")
	}
}

// SendGiftPack .
func SendGiftPack(c *gin.Context) {
	gift := &model.Gift{}
	if err := c.Bind(gift); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"channel": gift.Channel, "area": gift.Area, "giftpackname": gift.GiftPackName}
	if err := gift.FindOne(filter); err == nil {
		c.String(http.StatusBadRequest, "请不要重复发送该礼包！")
		return
	}
	if err := gift.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}

// AddTopic .
func AddTopic(c *gin.Context) {
	topic := &model.Topic{}
	if err := c.Bind(topic); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var (
		enc  = base64.StdEncoding
		path string
	)
	var img = topic.ImgURL
	if img[11] == 'j' {
		img = img[23:]
		path = fmt.Sprintf("/img/%s.jpg", topic.Name)
	} else if img[11] == 'p' {
		img = img[22:]
		path = fmt.Sprintf("/img/%s.png", topic.Name)
	} else if img[11] == 'g' {
		img = img[22:]
		path = fmt.Sprintf("/img/%s.gif", topic.Name)
	} else {
		fmt.Println("不支持该文件类型")
	}

	data, err := enc.DecodeString(img)
	if err != nil {
		log.Println(err.Error())
	}
	// 图片写入文件
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		log.Println(err)
	}
	//记录图片保存的地址
	path = "http://www.yinghuo2018.com" + path
	topic.ImgURL = path
	if err := topic.Save(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}
