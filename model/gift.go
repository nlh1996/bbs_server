package model

import (
	"bbs_server/config"
	"bbs_server/database"
	"sync"

	"gopkg.in/mgo.v2/bson"
)

// Gift .
type Gift struct {
	Channel      string
	Area         string
	GiftPackName string
	GiftPackNum  int
}

// RedeemCode 兑换码
type RedeemCode struct {
	Code string `bson:"_id"`
	// 是否使用了
	Used bool
	// 是否被领取
	Geted bool
	// 对应的礼包名
	GiftPackName string
	GiftPackId   string
	//所属渠道
	Channel string
	//所属区
	Area  string
	Start string
	End   string
}

// Save .
func (g *Gift) Save() error {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_gift")
	return c.Insert(g)
}

// Search .
func (g *Gift) Search(filter interface{}) (*[]Gift, error) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_gift")
	res := &[]Gift{}
	return res, c.Find(filter).All(res)
}

// FindOne .
func (g *Gift) FindOne(filter interface{}) error {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_gift")
	return c.Find(filter).One(g)
}

var mx = sync.Mutex{}

// Update .
func (g *Gift) Update() error {
	mx.Lock()
	defer mx.Unlock()
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_gift")
	filter := bson.M{"channel": g.Channel, "area": g.Area, "giftpackname": g.GiftPackName}
	return c.Update(filter, bson.M{"$inc": bson.M{"giftpacknum": -1}})
}

var mutex = sync.Mutex{}

// FindOne .
func (code *RedeemCode) FindOne(filter interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()
	session := database.Session2.Clone()
	defer session.Close()
	c := session.DB(config.GM).C("code")
	return c.Find(filter).One(code)
}

// Update .
func (code *RedeemCode) Update(filter interface{}) error {
	session := database.Session2.Clone()
	defer session.Close()
	c := session.DB(config.GM).C("code")
	return c.Update(bson.M{"_id": code.Code}, filter)
}
