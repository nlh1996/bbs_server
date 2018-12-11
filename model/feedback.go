package model

import (
	"bbs_server/database"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Complaint .
type Complaint struct {
	Theme      string        `json:"theme"`
	UName      string        `json:"name"`
	CreateTime string        `json:"createTime"`
	Commit     string        `json:"commit"`
	TID        bson.ObjectId `json:"tid"`
}

// Save 保存用户反馈
func (p *Complaint) Save() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_feedback")
	err := c.Insert(p)
	if err != nil {
		log.Fatal(err)
	}
}
