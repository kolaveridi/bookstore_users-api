package services

import (
	"github.com/kolaveridi/bookstore_users-api/datasources/mysql/users_db"
	"github.com/kolaveridi/bookstore_users-api/domain/users"
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	result := &users.User{Id: userId}
	if err := result.GET(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	//Email validated properly

	return &user, nil

}
