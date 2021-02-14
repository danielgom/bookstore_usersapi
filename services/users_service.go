package services

import (
	"github.com/danielgom/bookstore_usersapi/domain/users"
	"github.com/danielgom/bookstore_usersapi/gateways/usergateway"
	"github.com/danielgom/bookstore_usersapi/utils/errors"
)

func GetUser(uId int64) (*users.User, *errors.RestErr) {

	// Use this is we do not enable negative user ids
	if uId < 0 {
		return nil, errors.NewBadRequestError("Invalid user id")
	}

	user, err := usergateway.Get(uId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(u *users.User) (*users.User, *errors.RestErr) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := usergateway.Save(u); err != nil {
		return nil, err
	}

	return u, nil
}
