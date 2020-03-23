package services

import (
	"github.com/kolaveridi/bookstore_users-api/domain/users"
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
	"strings"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	// one of user or error should be kept nil
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return nil, errors.NewBadRequestError("Invalid email address")
	}

	return nil, nil

}
