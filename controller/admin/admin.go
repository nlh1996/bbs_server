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
