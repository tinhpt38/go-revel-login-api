package database

import (
	"time"
	"gopkg.in/mgo.v2"
)

var UserCollection *mgo.Collection
var Session *mgo.Session


//Init databse

func  InitDB(uri, dbname, user, password string) error{
	// session, err := mgo.Dial(uri)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
        Addrs:    []string{uri},
        Username: user,
        Password: password,
        Timeout:  100 * time.Second,
    })
	if(err != nil){
		return err 
	}

	session.SetMode(mgo.Monotonic,true)
	Session = session
	UserCollection = session.DB(dbname).C("user")
	return nil
} 