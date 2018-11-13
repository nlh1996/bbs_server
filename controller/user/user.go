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
	header := c.Request.Header["Authorization"]
	headerToken := header[0]
	user := &model.User{}
	name,ok := common.TokenMap[headerToken]
	if (ok) {
		user.UserName = name
		user = user.Search()
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":  user.UserName,
			"exp":   user.Exp,
			"jifen": user.Integral,
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
		c.String(http.StatusOK, "0")
	} else {
		c.String(http.StatusOK, "1")
	}

	c.String(http.StatusOK, "Register successful !!!")
	fmt.Println(*user)
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
	msg, result := user.Validator()
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
		//将用户token加入map缓存中
		common.TokenMap[tokenString] = user.UserName

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token": tokenString,
				"user":  user.UserName,
				"exp":   user.Exp,
				"jifen": user.Integral,
			},
			"msg":         msg,
		})
		return
	}
	c.String(http.StatusForbidden, msg)

}
