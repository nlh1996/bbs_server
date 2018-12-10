package main

import (
	"bbs_server/common"
	"bbs_server/database"
	"bbs_server/model"
	"bbs_server/router"
)

func main() {
	database.InitDB()
	//获取黑名单
	blackName := &model.BlackName{}
	common.BlackList = blackName.BlackList()
	router.Init()

}
