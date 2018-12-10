package model

import (
	"bbs_server/database"
	"bbs_server/utils"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
	// "fmt"
)

// User .
type User struct {
	UName      string   `json:"uName"`
	PassWord   string   `json:"password"`
	ReplyNum   uint32   `json:"replyNum"`
	Support    uint32   `json:"support"`
	Exp        uint32   `json:"exp"`
	Integral   uint32   `json:"integral"`
	CreateTime string		`json:"createTime"`
	SigninTime string   `json:"signinTime"`
	MyReply    []string `json:"myReply"`
	MyPosts    []string `json:"myPosts"`
	MyCollect  []string `json:"myCollect"`
	MySupport  []string `json:"mySupport"`
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
	err := c.Find(bson.M{"uname": pUser.UName}).One(result)
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
		if result[index].UName == pUser.UName {
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
		if result[index].UName == pUser.UName {
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
	c.Update(bson.M{"uname": pUser.UName}, bson.M{"$inc": bson.M{"exp": 10}})
	c.Update(bson.M{"uname": pUser.UName}, bson.M{"$inc": bson.M{"integral": 10}})
	err := c.Update(bson.M{"uname": pUser.UName}, bson.M{"$set": bson.M{"signintime": date}})
	if err != nil {
		log.Fatal(err)
	}
}

// SaveSupport 点赞记录贴子id
func (pUser *User) SaveSupport(tid string) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	err := c.Update(bson.M{"uname": pUser.UName}, bson.M{"$push": bson.M{"mysupport": tid}})
	if err != nil {
		return false
	}
	return true
}

// AddSupport 点赞数增加
func (pUser *User) AddSupport() bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	err := c.Update(bson.M{"uname": pUser.UName}, bson.M{"$inc": bson.M{"support": 1}})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// CancelSupport 删除点赞记录
func (pUser *User) CancelSupport(tid string) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	err := c.Update(bson.M{"uname": pUser.UName}, bson.M{"$pull": bson.M{"mysupport": tid}})
	if err != nil {
		return false
	}
	return true
}

// ReduceSupport 点赞数减少
func (pUser *User) ReduceSupport() bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	err := c.Update(bson.M{"uname": pUser.UName}, bson.M{"$inc": bson.M{"support": -1}})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}