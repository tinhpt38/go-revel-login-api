package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"login-token/app/database"
	"login-token/app/models"
	"github.com/revel/revel"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) Index() revel.Result {
	result := []models.User{}
	if err := database.UserCollection.Find(bson.M{}).All(&result); err != nil {
		c.Response.Status = http.StatusInternalServerError
		var body interface{}
		body = "Internal Server Error"
		message := &models.Message{Name: "err",Body:body}
		c.RenderJSON(message)
	}
	c.Response.Status = http.StatusOK
	message := &models.Message{Name: "success",Body:result}
	return c.RenderJSON(message)
}

func (c UserController) Delete(id string) revel.Result{
	if !bson.IsObjectIdHex(id){
		c.Response.Status = http.StatusBadRequest
		var body interface{}
		body = "Status bad request"
		message := &models.Message{Name: "err",Body:body}	
		return c.RenderJSON(message)
	}

	obj := bson.ObjectIdHex(id)
	if !obj.Valid(){
		c.Response.Status = http.StatusBadRequest
		message := &models.Message{Name: "err",Body:"ID is not valid"}	
		return c.RenderJSON(message)
	}

	err:= database.UserCollection.Remove(bson.M{"_id":obj})
	c.Response.Status = http.StatusOK
	if err != nil{
		message := &models.Message{Name:"err",Body:err.Error()}
		return c.RenderJSON(message)
	}
	message := &models.Message{Name:"err",Body:"ID was removed successfuly"}
	return c.RenderJSON(message);

}