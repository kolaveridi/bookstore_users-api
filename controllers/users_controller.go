package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kolaveridi/bookstore_users-api/domain/users"
	"github.com/kolaveridi/bookstore_users-api/services"
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// when we try to populate the model
	if err := c.ShouldBindJSON(&user); err != nil {
		//handling json error and return bad request to the caller
		fmt.Println(err)
		restErr := errors.RestError{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad request",
		}
		c.JSON(restErr.Status, restErr)
		return

	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//handle usercreation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement to be ")
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement to be ")
}
