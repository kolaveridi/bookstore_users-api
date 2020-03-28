package users

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/kolaveridi/bookstore_users-api/datasources/mysql/users_db"
	"github.com/kolaveridi/bookstore_users-api/utils/date_utils"
	"github.com/kolaveridi/bookstore_users-api/utils/errors"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created  FROM users WHERE id=?;"
	errorNoRows     = "no rows in result set"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) GET() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		sqlErr, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
		}
		fmt.Println(sqlErr.Number)
		fmt.Println(sqlErr.Message)

		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId

	return nil
}
