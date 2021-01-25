package models

type Message struct {
	Name string `json:"name" bson:"name"`
	Body interface{} `json:"data" bson:"data"`
}