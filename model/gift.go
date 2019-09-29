package model

import (
	"bbs_server/config"
	"bbs_server/database"
)

// Gift .
type Gift struct {
	Channel      string
	Area         string
	GiftPackName string
	GiftPackNum  int
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
func (g *Gift) FindOne(filter interface{}) (error) {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_gift")
	return c.Find(filter).One(g)
}
