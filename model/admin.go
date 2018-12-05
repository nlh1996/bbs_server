package model

import (
	"bbs_server/database"

	"gopkg.in/mgo.v2/bson"
)

// Admin .
type Admin struct {
	UName    string `json:"uName"`
	PassWord string `json:"password"`
}

// Validator .
func (admin *Admin) Validator() (string, bool) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	result := &Admin{}
	err := c.Find(bson.M{"uname": admin.UName}).One(result)
	var msg string
	if err != nil {
		msg = "没有该账户！"
		return msg, false
	}

	if result.PassWord != admin.PassWord {
		msg = "密码错误！"
		return msg, false
	}

	msg = "登录成功！"
	return msg, true
}
