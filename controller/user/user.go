package user

import (
	"bbs_server/common"
	"bbs_server/model"
	"bbs_server/utils"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// IsLoad .
func IsLoad(c *gin.Context) {
	user := &model.User{}
	user.UName = c.Request.Header["Authorization"][0]
	var isLoad int8
	if user.UName == "admin" {
		isLoad = 2
	} else {
		isLoad = 1
	}
	var result bool
	result, user = user.Search()
	isSignin := user.IsSignin()
	if result {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"user":     *user,
				"isSignin": isSignin,
				"isLoad":   isLoad,
			},
		})
	} else {
		c.String(http.StatusOK, "内部错误！")
	}
}

// Register 用户注册
func Register(c *gin.Context) {
	user := &model.User{}

	err := c.Bind(user)
	if err != nil {
		log.Println(err)
	}
	newPwd := utils.Jiami(&user.PassWord, &user.UName)
	user.PassWord = newPwd
	user.CreateTime = utils.GetTimeStr()
	if user.Find() == false {
		user.Save()
		//统计每天用户注册数量
		msg := &model.TodayMsg{}
		msg.Today = utils.GetDateStr()
		msg.RegisterSave()
		c.String(http.StatusOK, "Register successful !!!")
	} else {
		c.String(http.StatusOK, "用户名存在！")
	}
}

// Login 用户登录
func Login(c *gin.Context) {
	user := &model.User{}
	err := c.Bind(user)
	if err != nil {
		log.Println(err)
	}
	for _, item := range *common.BlackList {
		if item.UName == user.UName {
			c.String(http.StatusAccepted, "该账号已被封禁，请联系管理员解封。")
			return
		}
	}
	newPwd := utils.Jiami(&user.PassWord, &user.UName)
	user.PassWord = newPwd
	pUser, msg, result := user.Validator()
	if result {
		//统计每天用户登录情况
		msg := &model.TodayMsg{}
		msg.Today = utils.GetDateStr()
		msg.LoginSave()
		//返回token给用户
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.UName,
		})
		tokenString, err := token.SignedString([]byte("123"))
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusOK, "内部错误")
			return
		}
		//将用户token以键值对的方式加入map缓存中
		common.TokenMap[tokenString] = user.UName
		//是否签到过
		isSignin := pUser.IsSignin()
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token":    tokenString,
				"user":     *pUser,
				"isSignin": isSignin,
				"isLoad":   1,
			},
			"msg": msg,
		})
		return
	}
	c.String(http.StatusForbidden, msg)
}

//Signin 用户签到
func Signin(c *gin.Context) {
	user := &model.User{}
	user.UName = c.Request.Header["Authorization"][0]
	date := utils.GetDateStr()
	user.InsertDate(date)
	c.String(http.StatusOK, "")
}

// GetNotice 获取最新公告
func GetNotice(c *gin.Context) {
	notice := &model.Notice{}
	notice.Get()
	c.JSON(http.StatusOK, gin.H{
		"msg": notice.Message,
	})
}

// GetZhiDing 获取置顶帖
func GetZhiDing(c *gin.Context) {
	headPost := &model.HeadPost{}
	result := headPost.GetHeadPost()
	c.JSON(http.StatusOK, gin.H{
		"list": result,
	})
}

// GetMyPosts 获取用户自身的贴子
func GetMyPosts(c *gin.Context) {
	user := &model.User{}
	user.UName = c.Request.Header["Authorization"][0]
	myposts := user.Myposts()
	if myposts != nil {
		c.JSON(http.StatusOK, gin.H{
			"myposts": myposts,
		})
	}
}
