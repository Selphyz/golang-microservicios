package users

import (
	usersdb "bookstore/db/mysql/users_db"
	dateutils "bookstore/utils/date_utils"
	"bookstore/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := usersdb.DBClient.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = dateutils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
