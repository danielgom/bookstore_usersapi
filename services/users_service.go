package services

import (
	"github.com/danielgom/bookstore_usersapi/domain/users"
	"github.com/danielgom/bookstore_usersapi/gateways/usergateway"
	"github.com/danielgom/bookstore_usersapi/utils/cryptoutils"
	"github.com/danielgom/bookstore_utils-go/errors"
	"time"
)

var UserService userServiceInterface = &userService{}

type userService struct{}

type userServiceInterface interface {
	GetUserById(int64) (*users.User, *errors.RestErr)
	CreateUser(*users.User) (*users.User, *errors.RestErr)
	UpdateUser(*users.User) (*users.User, *errors.RestErr)
	UpdateUserPartial(*users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
	LoginUser(*users.UserLoginRequest) (*users.User, *errors.RestErr)
}

func (s *userService) GetUserById(uId int64) (*users.User, *errors.RestErr) {

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

func (s *userService) CreateUser(u *users.User) (*users.User, *errors.RestErr) {

	if err := u.Validate(true); err != nil {
		return nil, err
	}

	u.Status = users.StatusInactive
	u.DateCreated = time.Now()
	u.Password = cryptoutils.Encrypt(u.Password)

	if err := usergateway.Save(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *userService) UpdateUser(u *users.User) (*users.User, *errors.RestErr) {

	if err := u.Validate(false); err != nil {
		return nil, err
	}

	current, err := s.GetUserById(u.Id)
	if err != nil {
		return nil, err
	}

	current.FirstName = u.FirstName
	current.LastName = u.LastName
	current.Email = u.Email

	if err = usergateway.Update(current); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) UpdateUserPartial(u *users.User) (*users.User, *errors.RestErr) {

	current, err := s.GetUserById(u.Id)
	if err != nil {
		return nil, err
	}

	if u.FirstName != "" {
		current.FirstName = u.FirstName
	}
	if u.LastName != "" {
		current.LastName = u.LastName
	}
	if u.Email != "" {
		current.Email = u.Email
	}

	if err = usergateway.Update(current); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUser(uId int64) *errors.RestErr {
	restErr := validId(uId)
	if restErr != nil {
		return restErr
	}

	return usergateway.Delete(uId)
}

func (s *userService) Search(st string) (users.Users, *errors.RestErr) {
	return usergateway.FindByStatus(st)

}

func (s *userService) LoginUser(logReq *users.UserLoginRequest) (*users.User, *errors.RestErr) {
	user, err := usergateway.FindByEmailAndPassword(logReq.Email)
	if err != nil {
		return nil, err
	}

	if !cryptoutils.VerifyPassword(user.Password, logReq.Password) {
		return nil, errors.NewBadRequestError("Invalid user credentials")
	}

	return user, nil
}

func validId(id int64) *errors.RestErr {
	if id < 0 {
		return errors.NewBadRequestError("Invalid user id: user id should not be negative")
	}
	return nil
}

