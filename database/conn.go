package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

// Session .
var (
	Session *mgo.Session
	err     error
)

// InitDB 初始化数据库连接.
func InitDB() {
	Session, err = mgo.Dial("mongodb://212.129.149.224:32768?maxPoolSize=100")
	if err != nil {
		log.Println(err)
	}
	Session.SetPoolLimit(100)

	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)

}
