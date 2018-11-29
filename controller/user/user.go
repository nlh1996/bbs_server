package user

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

// IsLoad .
func IsLoad(c *gin.Context) {
	user := &model.User{}
	user.UserName = c.Request.Header["Authorization"][0]
	user = user.Search()
	isSignin := user.IsSignin()
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":  user.UserName,
			"exp":   user.Exp,
			"jifen": user.Integral,
			"isSignin": isSignin,
		},
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	user := &model.User{}
	fmt.Printf("%x\n", &user)

	err := c.Bind(user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	newPwd := utils.Jiami(&user.PassWord, &user.UserName)
	user.PassWord = newPwd

	if user.Find() == false {
		user.Save()
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
		fmt.Println(err)
		log.Fatal(err)
	}
	newPwd := utils.Jiami(&user.PassWord, &user.UserName)
	user.PassWord = newPwd
	pUser, msg, result := user.Validator()
	if result {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.UserName,
		})
		tokenString, err := token.SignedString([]byte("123"))
		if err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusOK, "内部错误")
			return
		}
		//将用户token以键值对的方式加入map缓存中
		common.TokenMap[tokenString] = user.UserName
		//是否签到过
		isSignin := pUser.IsSignin()
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token": tokenString,
				"user":  pUser.UserName,
				"exp":   pUser.Exp,
				"jifen": pUser.Integral,
				"isSignin": isSignin,
			},
			"msg":         msg,
		})
		return
	}
	c.String(http.StatusForbidden, msg)
}

//Signin 用户签到
func Signin(c *gin.Context) {
	user := &model.User{}
	user.UserName = c.Request.Header["Authorization"][0]
	date := utils.GetDateStr()
	user.InsertDate(date)
	c.String(http.StatusOK,"")
}
