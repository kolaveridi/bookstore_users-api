package users

import (
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
	"strings"
)

// when working with json this is the json format given on right
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func Validate(user *User) *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil

}
