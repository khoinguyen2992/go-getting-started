package mongodb

import (
	"ginserver/env"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session

func init() {
	if Session == nil {
		Session = getConnection()
	}
}

func getConnection() *mgo.Session {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    strings.Split(env.Get("MONGODB_HOST"), ","),
		Timeout:  5 * time.Second,
		Database: env.Get("MONGODB_DATABASE"),
		Username: env.Get("MONGODB_USERNAME"),
		Password: env.Get("MONGODB_PASSWORD"),
	})

	if err != nil {
		panic(err)
	}

	return s
}
