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
	Title      	string        `json:"title"`
	ImgList    	[]string      `json:"imgList"`
	ReplyNum    uint32        `json:"replyNum"`
	Support     uint32        `json:"support"`
	ReadNum     uint32        `json:"readNum"`
	ShareMsg		`bson:",inline"`
}

// Reply1 .
type Reply1 struct {
	ID        	bson.ObjectId `json:"id"`
	Show				bool					`json:"show"`
	ShareMsg 		`bson:",inline"`
}

// Reply2 .
type Reply2 struct {
	ID        	bson.ObjectId `json:"id"`
	RID     		bson.ObjectId `json:"rid"`
	RName 			string 			  `json:"rName"`
	Show				bool					`json:"show"`
	ShareMsg		`bson:",inline"`
}

// ShareMsg .
type ShareMsg struct {
	HeadImg			string				`json:"headImg"`
	UName 			string 				`json:"uName"`
	CreateTime 	string				`json:"createTime"`
	Content    	string    		`json:"content"`
	TID        	bson.ObjectId `json:"tid"`
}

// Save 保存贴子信息.
func (p *Post) Save() bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	c.Update(bson.M{"uname": p.TopStorey.UName}, bson.M{"$inc": bson.M{ "exp": 15 }})
	c.Update(bson.M{"uname": p.TopStorey.UName}, bson.M{"$inc": bson.M{ "integral": 15 }})
	c = session.DB("test").C("bbs_posts")
	err := c.Insert(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// UpdatePosts 获取所有贴子.
func UpdatePosts(postsPool *[]Post) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Find(bson.M{}).Sort("-_id").All(postsPool)
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

// Save 保存一级回复信息.
func (reply1 *Reply1) Save(tid bson.ObjectId) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	c.Update(bson.M{"username": reply1.UName}, bson.M{"$inc": bson.M{ "exp": 5 }})
	c.Update(bson.M{"username": reply1.UName}, bson.M{"$inc": bson.M{ "integral": 5 }})
	c = session.DB("test").C("bbs_posts")
	err := c.Update(bson.M{"tid": tid}, bson.M{"$push": bson.M{ "relist1": reply1 }})
	err = c.Update(bson.M{"tid": tid}, bson.M{"$inc": bson.M{ "topstorey.replynum": 1 }})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Save 保存二级回复信息
func (reply2 *Reply2) Save(id bson.ObjectId) bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_user")
	c.Update(bson.M{"username": reply2.UName}, bson.M{"$inc": bson.M{ "exp": 5 }})
	c.Update(bson.M{"username": reply2.UName}, bson.M{"$inc": bson.M{ "integral": 5 }})
	c = session.DB("test").C("bbs_posts")
	err := c.Update(bson.M{"tid": id}, bson.M{"$push": bson.M{ "relist2": reply2 }})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Del 删除贴子
func (p *Post) Del(tid bson.ObjectId,name string) bool{
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Find(bson.M{"tid": tid}).One(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if p.TopStorey.UName == name {
		c.Remove(bson.M{"tid": tid})
		return true
	}
	return false
}

// AddSupport 增加点赞数
func (p *Post) AddSupport() bool {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("bbs_posts")
	err := c.Update(bson.M{"tid": p.TID}, bson.M{"$inc": bson.M{ "topstorey.support": 1 }})
	if err != nil {
		fmt.Println("333")
		return false
	}
	return true
}