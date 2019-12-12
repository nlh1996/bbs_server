package main

import (
	"bbs_server/common"
	"bbs_server/config"
	"bbs_server/database"
	"bbs_server/model"
	"bbs_server/router"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	database.InitDB()
	// 获取黑名单
	blackName := &model.BlackName{}
	common.BlackList = blackName.BlackList()
	router.Init()
	result := []model.User{}
	if err := test(&result); err != nil {
		log.Println(err)
	}
	fmt.Println(result)
}

func test(result *[]model.User) error {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB(config.DbName).C("bbs_user")
	filter := bson.M{"$or": []bson.M{bson.M{"uname": "111"}, bson.M{"uname": "191118022753465"}}}
	return c.Find(filter).All(result)
}
