package model

import (
	"bbs_server/database"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Admin .
type Admin struct {
	UName    string `json:"uName"`
	PassWord string `json:"password"`
}

// TodayMsg 今日用户数据.
type TodayMsg struct {
	Today         string
	TodayAccess   uint32
	TodayLogin    uint32
	TodayRegister uint32
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

// LoginSave 每天用户登录数据保存
func (msg *TodayMsg) LoginSave() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_count")
	result := &TodayMsg{}
	err := c.Find(bson.M{"today": msg.Today}).One(result)
	if err != nil {
		c.Insert(msg)
		return
	}
	c.Update(bson.M{"today": msg.Today}, bson.M{"$inc": bson.M{"todaylogin": 1}})
}

// RegisterSave 每天用户注册数据保存
func (msg *TodayMsg) RegisterSave() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_count")
	result := &TodayMsg{}
	err := c.Find(bson.M{"today": msg.Today}).One(result)
	if err != nil {
		c.Insert(msg)
		return
	}
	c.Update(bson.M{"today": msg.Today}, bson.M{"$inc": bson.M{"todayregister": 1}})
}

// AccessSave 每天论坛访问量保存
func (msg *TodayMsg) AccessSave() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_count")
	result := &TodayMsg{}
	err := c.Find(bson.M{"today": msg.Today}).One(result)
	if err != nil {
		c.Insert(msg)
		return
	}
	c.Update(bson.M{"today": msg.Today}, bson.M{"$inc": bson.M{"todayaccess": 1}})
}

//Search 查询今日统计结果
func (msg *TodayMsg) Search() *TodayMsg {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_count")
	result := &TodayMsg{}
	err := c.Find(bson.M{"today": msg.Today}).One(result)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return result
}

//Count 统计总注册用户人数
func (msg *TodayMsg) Count() int {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	num,err:= c.Find(nil).Count()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return num
}