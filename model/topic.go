package model

import (
	"bbs_server/config"
	"bbs_server/database"

	"gopkg.in/mgo.v2/bson"
)

// Topic .
type Topic struct {
	Name   string
	ImgURL string
	Num    int
}

// Save .
func (g *Topic) Save() error {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_topics")
	return c.Insert(g)
}

// GetTopics .
func GetTopics(Topics *[]Topic) error {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_topics")
	return c.Find(bson.M{}).All(Topics)
}
