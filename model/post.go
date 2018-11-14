package model

import (
	"bbs_server/database"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Post 贴子结构
type Post struct {
	TopStorey  `json:"topStorey"`
	Reply1     `json:"reply1"`
	Reply2     `json:"reply2"`
	UpdateTime time.Time `json:"time"`
}

// TopStorey .
type TopStorey struct {
	UID        string    `json:"uid"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ReadNum    int32     `json:"readNum"`
	Support    int32     `json:"support"`
	ReplyNum   int32     `json:"replyNum"`
	CreateTime time.Time `json:"createTime"`
	ImgList    []string  `json:"imgList"`
}

// Reply1 .
type Reply1 []struct {
	UID        string    `json:"uid"`
	Index      int32     `json:"index"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
}

// Reply2 .
type Reply2 []struct {
	UID        string    `json:"uid"`
	Index      int32     `json:"index"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
}

// Save .
func (t *Post) Save() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdatePosts .
func UpdatePosts(postsPool *[]Post) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	c.Find(bson.M{}).All(postsPool)
}
