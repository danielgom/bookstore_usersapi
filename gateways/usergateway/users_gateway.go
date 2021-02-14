package usergateway

import (
	"fmt"
	"github.com/danielgom/bookstore_usersapi/domain/users"
	"github.com/danielgom/bookstore_usersapi/utils/dateutils"
	"github.com/danielgom/bookstore_usersapi/utils/errors"
)

var (
	usersDB = make(map[int64]*users.User)
)

func Get(uId int64) (*users.User, *errors.RestErr) {
	user := usersDB[uId]
	if user == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User %d not found", uId))
	}
	return user, nil
}

func Save(u *users.User) *errors.RestErr {
	result := usersDB[u.Id]
	if result != nil {
		if result.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", u.Id))
	}
	u.DateCreated = dateutils.GetNowMxString()
	usersDB[u.Id] = u
	return nil
}
