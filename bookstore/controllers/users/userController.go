package users

import (
	"bookstore/entities"
	"bookstore/services"
	"bookstore/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context)  {
	var user entities.User
	if err:= c.ShouldBindJSON(&user); err!=nil{
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "getUser to be implemented")
}