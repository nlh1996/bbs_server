package database

import (
	"gopkg.in/mgo.v2"
)

// Session .
var Session *mgo.Session
var err error

// InitDB 初始化数据库连接.
func InitDB() {
	Session,err = mgo.Dial("mongodb://admin:admin@115.159.77.155:11000?maxPoolSize=100")
	if err != nil {
		panic(err)
	}
	Session.SetPoolLimit(100)

	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)

}



