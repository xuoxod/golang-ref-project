package collections

import (
	"encoding/json"
)

type Details interface {
	UserInfo() (string, error)
}

type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type User struct {
	UID            string  `json:"uid"`
	FirstName      string  `json:"firstname"`
	LastName       string  `json:"lastname"`
	ContactDetails Contact `json:"contactdetails"`
}

type Accounts struct {
	Users []User `json:"users"`
}

func (u User) Details() (string, error) {
	response, err := json.Marshal(u)

	if err != nil {
		return "", err
	}

	return string(response), nil
}
