package model

import (
	"bbs_server/database"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Post 贴子结构
type Post struct {
	TopStorey  `json:"topStorey"`
	Reply1     `json:"reply1"`
	Reply2     `json:"reply2"`
	UpdateTime string `json:"time"`
	TID				 string	`json:"tid"`
}

// TopStorey .
type TopStorey struct {
	UID        string    `json:"uid"`
	TID				 string		 `json:"tid"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ReadNum    int32     `json:"readNum"`
	Support    int32     `json:"support"`
	ReplyNum   int32     `json:"replyNum"`
	CreateTime string		 `json:"createTime"`
	ImgList    []string  `json:"imgList"`
}

// Reply1 .
type Reply1 []struct {
	UID        string    `json:"uid"`
	Index      int32     `json:"index"`
	Content    string    `json:"content"`
	CreateTime string		 `json:"createTime"`
}

// Reply2 .
type Reply2 []struct {
	UID        string    `json:"uid"`
	Index      int32     `json:"index"`
	Content    string    `json:"content"`
	CreateTime string 	 `json:"createTime"`
}

// Save .
func (p *Post) Save() bool{
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Insert(p)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdatePosts .
func UpdatePosts(postsPool *[]Post) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
 	err := c.Find(bson.M{}).All(postsPool) 
	if err != nil {
		log.Fatal(err)
	} 
}

// Get .
func (p *Post) Get(tid string) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	c.Find(bson.M{"tid": tid}).One(p)

}
