package services

import (
	"github.com/danielgom/bookstore_usersapi/domain/users"
	"github.com/danielgom/bookstore_usersapi/gateways/usergateway"
	"github.com/danielgom/bookstore_usersapi/utils/errors"
	"time"
)

// Use this if we do not enable negative user ids

func validId(id int64) *errors.RestErr {
	if id < 0 {
		return errors.NewBadRequestError("Invalid user id: user id should not be negative")
	}
	return nil
}

func GetUserById(uId int64) (*users.User, *errors.RestErr) {

	restErr := validId(uId)
	if restErr != nil {
		return nil, restErr
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

	u.Status = users.StatusInactive
	u.DateCreated = time.Now()

	if err := usergateway.Save(u); err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateUser(isPartial bool, u *users.User) (*users.User, *errors.RestErr) {
	current, err := GetUserById(u.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if u.FirstName != "" {
			current.FirstName = u.FirstName
		}
		if u.LastName != "" {
			current.LastName = u.LastName
		}
		if u.Email != "" {
			current.Email = u.Email
		}
	} else {
		current.FirstName = u.FirstName
		current.LastName = u.LastName
		current.Email = u.Email
	}

	if err = usergateway.Update(current); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(uId int64) *errors.RestErr {
	restErr := validId(uId)
	if restErr != nil {
		return restErr
	}

	return usergateway.Delete(uId)
}

func Search(s string) ([]users.User, *errors.RestErr) {
	usersList, restErr := usergateway.FindByStatus(s)
	if restErr != nil {
		return nil, restErr
	}

	return usersList, nil
}
