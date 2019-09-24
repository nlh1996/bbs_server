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
