// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Client struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	SecondName  string    `json:"second_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Language    string    `json:"language"`
	Currency    string    `json:"currency"`
	Time        string    `json:"time"`
	Password    string    `json:"password"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
