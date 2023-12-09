package models

import "time"

// User
type User struct {
	ID            string    `json:"id"`
	FirstName     string    `json:"fname"`
	LastName      string    `json:"lname"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Password      string    `json:"password"`
	EmailVerified bool      `json:"everify"`
	PhoneVerified bool      `json:"pverify"`
	AccessLevel   int       `json:"access_level"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
