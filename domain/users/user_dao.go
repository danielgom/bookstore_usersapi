package users

import (
	"github.com/danielgom/bookstore_usersapi/utils/errors"
	"regexp"
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
	Password    string    `json:"password"`
}

type Users []User

// Avoid breaking a solid principle by letting the user to validate itself

func (u *User) Validate(create bool) *errors.RestErr {

	if strings.TrimSpace(u.FirstName) == "" {
		return errors.NewBadRequestError("First name field cannot be empty")
	}

	if strings.TrimSpace(u.LastName) == "" {
		return errors.NewBadRequestError("Last name field cannot be empty")
	}

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if !isValidMail(u.Email) {
		return errors.NewBadRequestError("Invalid email address")
	}

	if strings.TrimSpace(strings.ToLower(u.Password)) == "" && create {
		return errors.NewBadRequestError("Invalid password: missing field or empty password")
	}

	return nil
}

func isValidMail(mail string) bool {
	expr := "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"

	emailRegex := regexp.MustCompile(expr)
	return emailRegex.MatchString(mail)
}
