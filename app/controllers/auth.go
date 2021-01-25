package controllers

import (
	"encoding/json"
	// "login-token/app/jwt"
	"login-token/app/database"
	"login-token/app/models"
	"github.com/revel/revel"
	"github.com/asaskevich/govalidator"
	"net/http"
	"gopkg.in/mgo.v2/bson" 
)


type AuthController struct {
	*revel.Controller
}


func (c AuthController) Login() revel.Result {
	defer c.Request.Destroy()
	login := &models.Login{}

	err := json.NewDecoder(c.Request.GetBody()).Decode(&login)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(&models.Message{Name:"err",Body:"Could parse request"})
	}

	if govalidator.IsNull(login.Name) || govalidator.IsNull(login.Password) {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(&models.Message{Name:"err",Body:"name or password is empty"})
	}

	collection := database.UserCollection;
	result := &models.User{}
	if err := collection.Find(bson.M{"name":login.Name}).One(&result); err != nil{
		c.Response.Status = http.StatusBadRequest
		var body interface{}
		body = "user not found"
		message := &models.Message{Name:"err",Body:body}
		return c.RenderJSON(message)
	}
	err = models.CheckPasswordHash(result.Password,login.Password)
	if(err != nil){
		c.Response.Status = http.StatusBadRequest
		var body interface{}
		body = "user or password is not correct"
		message := &models.Message{Name:"err",Body:body}
		return c.RenderJSON(message)
	}

	// token, errCreate := jwt.Create(username)
	// if errCreate != nil {
	// 	c.Response.Status = http.StatusBadRequest
	// 	var body interface{}
	// 	body = "t"
	// 	message := &models.Message{Name:"err",Body:body}
	// 	return c.RenderJSON(message)
	// }

	c.Response.Status = http.StatusOK
	var body = map[string]interface{}{
		"result":result,
		"token":"token",
	}
	message := &models.Message{Name:"success",Body:body}
	return c.RenderJSON(message)

}


func (c AuthController) Register() revel.Result{
	defer c.Request.Destroy()
	user := &models.User{}
	err := json.NewDecoder(c.Request.GetBody()).Decode(&user)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(&models.Message{Name:"err",Body:"Could parse request"})
	}

	hashPassword ,err := models.Hash(user.Password)
	if err != nil {
		return c.RenderJSON(&models.Message{Name:"err",Body:"Could hash password"})
	}
	user.Password = hashPassword

	if err := database.UserCollection.Insert(user); err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON( &models.Message{Name:"err",Body:"Internal server error"})
	}
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(&models.Message{Name:"created",Body:user})
}