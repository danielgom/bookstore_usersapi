package users

import (
	"encoding/json"
	"log"
	"time"
)

type PublicUser struct {
	Id          int64     `json:"id"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
}

type PrivateUser struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
}

func (users Users) Marshaller(isPublic bool) []interface{} {
	res := make([]interface{}, len(users))
	for i, u := range users {
		res[i] = u.Marshaller(isPublic)
	}
	return res
}

func (u *User) Marshaller(isPublic bool) interface{} {
	if isPublic {
		return &PublicUser{
			Id:          u.Id,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}

	uJson, _ := json.Marshal(u)
	var privateUser PrivateUser
	err := json.Unmarshal(uJson, &privateUser)
	if err != nil {
		log.Fatal(err)
	}

	return privateUser
}
