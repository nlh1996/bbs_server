package admin

import (
	"bbs_server/common"
	"bbs_server/model"
	"bbs_server/utils"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login 管理员登录
func Login(c *gin.Context) {
	admin := &model.Admin{}
	err := c.Bind(admin)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
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
			"data": gin.H{
				"token": tokenString,
			},
			"msg": msg,
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
		"count": *msg,
		"userNum": num,
	})
}


// UserSearch 用户搜索。
func UserSearch (c *gin.Context) {
	user := &model.User{}
	user.UName = c.PostForm("name")
	user = user.Search()
	c.JSON(http.StatusOK, gin.H{
		"name": user.UName,
		"level": user.Exp,
		"jifen": user.Integral,
		"time": user.CreateTime,
	})
}

// AddBlackList 加入黑名单
func AddBlackList(c *gin.Context) {
	user := &model.BlackName{}
	user.UName = c.PostForm("name")
	user.Time = utils.GetTimeStr()
	result := user.BlackNameSave()
	if result == true {
		c.String(http.StatusOK,"scess")
		common.BlackList = user.BlackList()
	}else {
		c.String(http.StatusAccepted,"用户已经拉黑")
	}
}

// RemoveBlackList 移出黑名单
func RemoveBlackList(c *gin.Context) {
	user := &model.BlackName{}
	user.UName = c.PostForm("name")
	result := user.BlackNameRemove()
	if result == true {
		c.String(http.StatusOK,"删除成功！")
		common.BlackList = user.BlackList()
	}else{
		c.String(http.StatusOK,"内部出错！")
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
