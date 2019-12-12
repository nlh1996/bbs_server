package model

import (
	"bbs_server/config"
	"bbs_server/database"
	"sync"

	"gopkg.in/mgo.v2/bson"
)

// Gift .
type Gift struct {
	Comment      string
	Jifen        uint32
	GiftPackName string
	GiftPackNum  int
}

type MyGift struct {
	GiftPackName string
	Code         string
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
	err := c.Find(filter).All(res)
	return res, err
}

// Del .
func (g *Gift) Del(filter interface{}) error {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_gift")
	err := c.Remove(filter)
	return err
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
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_gift")
	filter := bson.M{"giftpackname": g.GiftPackName}
	err := c.Update(filter, bson.M{"$inc": bson.M{"giftpacknum": -1}})
	mx.Unlock()
	return err
}

var mutex = sync.Mutex{}

// FindOne .
func (code *RedeemCode) FindOne(filter interface{}) error {
	mutex.Lock()
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.GM).C("code")
	err := c.Find(filter).One(code)
	mutex.Unlock()
	return err
}

// Update .
func (code *RedeemCode) Update(filter interface{}) error {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.GM).C("code")
	return c.Update(bson.M{"_id": code.Code}, filter)
}

// Count .
func (rc *RedeemCode) Count(filter interface{}) (int, error) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.GM).C("code")
	return c.Find(filter).Count()
}