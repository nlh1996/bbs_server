package main

import (
	"fmt"
	"bbs_server/common"
	"bbs_server/database"
	"bbs_server/model"
	"bbs_server/router"
)

func main() {
	fmt.Println("111")
	database.InitDB()
	//获取黑名单
	blackName := &model.BlackName{}
	common.BlackList = blackName.BlackList()
	router.Init()

}
