package services

import (
	"github.com/kolaveridi/bookstore_users-api/domain/users"
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	// one of user or error should be kept nil
	return &user, nil
	return &user, &errors.RestError{
		Message: "",
		Status:  http.StatusInternalServerError,
		Error:   "",
	}
}
