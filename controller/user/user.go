package user

import (
	"bbs_server/model"
	"bbs_server/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	user := &model.User{}
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

	// cookie := &http.Cookie{
	// 	Name:     "session_id",
	// 	Value:    "123",
	// 	Path:     "/",
	// 	HttpOnly: true,
	// }
	// http.SetCookie(c.Writer, cookie)
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
	msg := user.Validator()
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"status_code": http.StatusOK,
			"status": "OK",
		},
		"data": user.UserName,
		"msg": msg,
	})

}
