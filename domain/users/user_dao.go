package users

import (
	"github.com/danielgom/bookstore_usersapi/utils/errors"
	"strings"
	"time"
)

const (
	StatusInactive = "inactive"
)

// - in json means do not retrieve the password
type User struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
	Password    string    `json:"-"`
}

// Avoid breaking a solid principle by letting the user to validate itself

func (u *User) Validate() *errors.RestErr {

	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	//TODO: Avoid empty username and password when creating

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	u.Password = strings.TrimSpace(strings.ToLower(u.Password))
	if u.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}
