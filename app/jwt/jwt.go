package jwt

import (
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func Create(username string) (string, error) {
 claims := jwt.MapClaims{}
 claims["authorized"] = true	
 claims["username"] = username
 //token het han
 claims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}