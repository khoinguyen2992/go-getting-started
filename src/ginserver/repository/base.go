package repository

import (
	"ginserver/database/mongodb"
	"ginserver/env"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	db = mongodb.Session.DB(env.Get("MONGODB_DATABASE"))
}

type BaseRepository struct {
	Collection *mgo.Collection
}
