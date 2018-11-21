package model

import (
	"bbs_server/database"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// Post 贴子结构
type Post struct {
	TopStorey  							  `json:"topStorey"`
	ReList1     []Reply1      `json:"reList1"`
	ReList2     []Reply2      `json:"reList2"`
	UpdateTime  string        `json:"time"`
	TID         bson.ObjectId `json:"tid"`
}

// TopStorey .
type TopStorey struct {
	TID         bson.ObjectId `json:"tid"`
	Title      	string        `json:"title"`
	ImgList    	[]string      `json:"imgList"`
	HeadImg			string				`json:"headImg"`
	UName 			string 				`json:"uName"`
	CreateTime 	string				`json:"createTime"`
	Content    	string    		`json:"content"`
	ReplyNum    uint32        `json:"replyNum"`
	Support     uint32        `json:"support"`
	ReadNum     uint32        `json:"readNum"`
}

// Reply1 .
type Reply1 struct {
	ID        	bson.ObjectId `json:"id"`
	HeadImg			string				`json:"headImg"`
	UName 			string 				`json:"uName"`
	CreateTime 	string				`json:"createTime"`
	Content    	string    		`json:"content"`
	TID        	bson.ObjectId `json:"tid"`
}

// Reply2 .
type Reply2 struct {
	ID        	bson.ObjectId `json:"id"`
	RID     		bson.ObjectId `json:"rid"`
	RName 			string 			  `json:"rName"`
	HeadImg			string				`json:"headImg"`
	UName 			string 				`json:"uName"`
	CreateTime 	string				`json:"createTime"`
	Content    	string    		`json:"content"`
	TID        	bson.ObjectId `json:"tid"`
}

// // ShareMsg .
// type ShareMsg struct {
// 	HeadImg			string				`json:"headImg"`
// 	UName 			string 				`json:"uName"`
// 	CreateTime 	string				`json:"createTime"`
// 	Content    	string    		`json:"content"`
// 	TID        	bson.ObjectId `json:"tid"`
// }

// Save 保存贴子信息.
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

// UpdatePosts 获取指定数量贴子.
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

// Get 获取单个贴子详情.
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

// Save 保存回复信息.
func (reply1 *Reply1) Save(tid bson.ObjectId) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Update(bson.M{"tid": tid}, bson.M{"$push": bson.M{ "relist1": reply1 }})
	err = c.Update(bson.M{"tid": tid}, bson.M{"$inc": bson.M{ "topstorey.replynum": 1 }})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
