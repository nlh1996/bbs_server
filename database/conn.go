package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

// Session .
var (
	Session *mgo.Session
	Session2 *mgo.Session
	err     error
)

// InitDB 初始化数据库连接.
func InitDB() {
	Session, err = mgo.Dial("mongodb://nlh:111@115.159.77.155:11600?maxPoolSize=100")
	//Session,err = mgo.Dial("mongodb://localhost:27017?maxPoolSize=500")
	if err != nil {
		log.Println(err)
	}
	Session.SetPoolLimit(100)

	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)

	Session2, err = mgo.Dial("mongodb://admin:admin@115.159.77.155:11000?maxPoolSize=100")
	//Session,err = mgo.Dial("mongodb://localhost:27017?maxPoolSize=500")
	if err != nil {
		log.Println(err)
	}
	Session2.SetPoolLimit(100)

	// Optional. Switch the session to a monotonic behavior.
	Session2.SetMode(mgo.Monotonic, true)

}
