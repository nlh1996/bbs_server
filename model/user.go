package model

import (
	"bbs_server/database"
	"log"

	"gopkg.in/mgo.v2/bson"
	// "fmt"
)

// User .
type User struct {
	UserName     string `json:"username"`
	PassWord     string	`json:"password"`
	ReplyNum	   string	`json:"reply_num"`
	SupportNum   string `json:"support_num"`
	BrowseNum    string `json:"browse_num"`
	Exp					 int		`json:"exp"`
	Integral		 int		`json:"integral"`
	ThreadID		 []string`json:"thread_id"`

	// lastLoginAt string
}

// Save .
func (t *User) Save() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}
}

// Validator .
func (t *User) Validator() string {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	result := User{}
	err := c.Find(bson.M{"username": t.UserName}).One(&result)
	var msg string

	if err != nil {
		msg = "没有该账户！"
		return msg
	}

	if result.PassWord != t.PassWord {
		msg = "密码错误！"
		return msg
	}

	msg = "登录成功！"
	return msg
}

// Find .
func (t *User) Find() bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	result := []User{}
	c.Find(bson.M{}).All(&result)
	for index := range result {
		if result[index].UserName == t.UserName {
			return true
		}
	}
	return false
}
