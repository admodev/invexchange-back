package usermodel

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DocumentNumber string    `json:"document_number"`
	Birthdate      time.Time `json:"birthdate"`
}
