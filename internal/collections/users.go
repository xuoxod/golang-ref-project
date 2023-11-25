package collections

import "fmt"

type Details interface {
	UserInfo() map[string]interface{}
	Info()
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

func (u User) UserInfo() map[string]interface{} {
	info := make(map[string]interface{})
	info["uid"] = u.UID
	info["fname"] = u.FirstName
	info["lname"] = u.LastName
	info["contact"] = u.ContactDetails
	return info
}

func (u User) Info() {
	fmt.Println("UID: ", u.UID)
	fmt.Println("First Name: ", u.FirstName)
	fmt.Println("Last Name: ", u.LastName)

	email := u.ContactDetails.Email
	phone := u.ContactDetails.Phone

	if email != "" {
		fmt.Println("Last Name: ", email)
	}

	if phone != "" {
		fmt.Println("Last Name: ", phone)
	}
}
