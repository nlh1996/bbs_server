package model

import (
	"bbs_server/database"
	"log"
	"bbs_server/config"
	"gopkg.in/mgo.v2/bson"
)

// Complaint 用户反馈信息的结构体.
type Complaint struct {
	Theme      string        `json:"theme"`
	UName      string        `json:"name"`
	CreateTime string        `json:"createTime"`
	Commit     string        `json:"commit"`
	TID        bson.ObjectId `json:"tid"`
	Status     int8          `json:"status"`
}

// Save 保存用户反馈
func (p *Complaint) Save() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_feedback")
	err := c.Insert(p)
	if err != nil {
		log.Fatal(err)
	}
}

// FeedList0 获取未处理用户反馈信息
func (p *Complaint) FeedList0() *[]Complaint {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_feedback")
	list := &[]Complaint{}
	err := c.Find(bson.M{"status": 0}).All(list)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return list
}

// FeedList1 获取已处理用户反馈信息
func (p *Complaint) FeedList1() *[]Complaint {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_feedback")
	list := &[]Complaint{}
	err := c.Find(bson.M{"status": 1}).All(list)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return list
}

// Del 删除用户反馈
func (p *Complaint) Del(tid string) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_feedback")
	id := bson.ObjectIdHex(tid)
	err := c.Remove(bson.M{"tid": id})
	if err != nil {
		return false
	}
	return true
}

// Agree 同意用户反馈
func (p *Complaint) Agree(tid string) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_feedback")
	id := bson.ObjectIdHex(tid)
	_, err := c.UpdateAll(bson.M{"tid": id}, bson.M{"$set": bson.M{"status": 1}})
	if err != nil {
		return false
	}
	return true
}
