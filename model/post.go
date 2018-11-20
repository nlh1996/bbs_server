package model

import (
	"bbs_server/database"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// Post 贴子结构
type Post struct {
	TopStorey  `json:"topStorey"`
	ReList1    []Reply1      `json:"reList1"`
	ReList2    []Reply2      `json:"reList2"`
	UpdateTime string        `json:"time"`
	TID        bson.ObjectId `json:"tid"`
}

// TopStorey .
type TopStorey struct {
	UID        string        `json:"uid"`
	TID        bson.ObjectId `json:"tid"`
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	ReadNum    int32         `json:"readNum"`
	Support    int32         `json:"support"`
	ReplyNum   int32         `json:"replyNum"`
	CreateTime string        `json:"createTime"`
	ImgList    []string      `json:"imgList"`
}

// Reply1 .
type Reply1 struct {
	UID        string        `json:"uid"`
	RID        bson.ObjectId `json:"rid"`
	TID        bson.ObjectId `json:"tid"`
	Content    string        `json:"content"`
	CreateTime string        `json:"createTime"`
}

// Reply2 .
type Reply2 struct {
	UID        string        `json:"uid"`
	RID        bson.ObjectId `json:"rid"`
	TID        bson.ObjectId `json:"tid"`
	Content    string        `json:"content"`
	CreateTime string        `json:"createTime"`
}

// Save .
func (p *Post) Save() bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Insert(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// UpdatePosts .
func UpdatePosts(postsPool *[]Post) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Find(bson.M{}).All(postsPool)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Get .
func (p *Post) Get(tid bson.ObjectId) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Find(bson.M{"tid": tid}).One(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Save .
func (reply1 *Reply1) Save(tid bson.ObjectId) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Update(bson.M{"tid": tid}, bson.M{"$push": bson.M{ "relist1": reply1 }})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
