package users

import (
	"fmt"
	"github.com/kolaveridi/bookstore_users-api/utils/date_utils"
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) GET() *errors.RestError {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf(" user %id not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf(" email %s already exists", user.Email))
		}
		return errors.NewNotFoundError(fmt.Sprintf(" user %id already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	userDB[user.Id] = user

	return nil
}
