package models


type Login struct {
	Name string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}