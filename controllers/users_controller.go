package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kolaveridi/bookstore_users-api/domain/users"
	"github.com/kolaveridi/bookstore_users-api/services"
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// when we try to populate the model
	if err := c.ShouldBindJSON(&user); err != nil {
		//handling json error and return bad request to the caller
		fmt.Println(err)
		restErr := errors.NewBadRequestError("invalid json body")
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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		//handle usercreation error
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement to be ")
}
