package main

import (
	"bbs_server/database"
	"bbs_server/router"
)

func main() {
	database.InitDB()
	router.Init()

}
