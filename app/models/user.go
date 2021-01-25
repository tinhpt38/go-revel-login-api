package models

import(
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"

)

type User struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword( []byte(password),14)
	return string(bytes), err
}

func CheckPasswordHash(hashed, password string)  error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}


