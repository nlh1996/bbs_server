package model

import (
	"bbs_server/database"
	"bbs_server/utils"
	"log"

	"gopkg.in/mgo.v2/bson"
	// "fmt"
)

// User .
type User struct {
	UserName   string   `json:"username"`
	PassWord   string   `json:"password"`
	ReplyNum   int32    `json:"replyNum"`
	Support    int32    `json:"supportNum"`
	Exp        int32    `json:"exp"`
	Integral   int32    `json:"integral"`
	SigninTime string   `json:"signinTime"`
	MyReply    []string `json:"myReply"`
	MyPosts    []string `json:"myPosts"`
	MyCollect  []string `json:"myCollect"`
	// lastLoginAt string
}

// Save .
func (pUser *User) Save() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	err := c.Insert(pUser)
	if err != nil {
		log.Fatal(err)
	}
}

// Validator .
func (pUser *User) Validator() (*User, string, bool) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	result := &User{}
	err := c.Find(bson.M{"username": pUser.UserName}).One(result)
	var msg string
	if err != nil {
		msg = "没有该账户！"
		return nil, msg, false
	}

	if result.PassWord != pUser.PassWord {
		msg = "密码错误！"
		return nil, msg, false
	}

	msg = "登录成功！"
	return result, msg, true
}

// Find .
func (pUser *User) Find() bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	result := []User{}
	c.Find(bson.M{}).All(&result)
	for index := range result {
		if result[index].UserName == pUser.UserName {
			return true
		}
	}
	return false
}

// Search .
func (pUser *User) Search() *User {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	result := []User{}
	c.Find(bson.M{}).All(&result)
	for index := range result {
		if result[index].UserName == pUser.UserName {
			return &result[index]
		}
	}
	return pUser
}

// IsSignin 判断是否签到
func (pUser *User) IsSignin() bool {
	if pUser.SigninTime == utils.GetDateStr() {
		return true
	}
	return false
}

// InsertDate 插入签到日期
func (pUser *User) InsertDate(date string) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	c.Update(bson.M{"username": pUser.UserName}, bson.M{"$inc": bson.M{ "exp": 10 }})
	c.Update(bson.M{"username": pUser.UserName}, bson.M{"$inc": bson.M{ "integral": 10 }})
	err := c.Update(bson.M{"username": pUser.UserName}, bson.M{"$set": bson.M{"signintime": date}})
	if err != nil {
		log.Fatal(err)
	}
}
