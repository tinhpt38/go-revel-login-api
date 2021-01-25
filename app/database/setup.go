package database

import (
	"gopkg.in/mgo.v2"
)

var UserCollection *mgo.Collection
var Session *mgo.Session


//Init databse

func  InitDB(uri, dbname string) error{
	session, err := mgo.Dial(uri)
	if(err != nil){
		return err 
	}

	session.SetMode(mgo.Monotonic,true)
	Session = session
	UserCollection = session.DB(dbname).C("user")
	return nil
} 