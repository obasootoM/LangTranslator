// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: client.sql

package db

import (
	"context"
)

const createClient = `-- name: CreateClient :one
INSERT INTO client (
  first_name,second_name,email,phone_number,language,password
) VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, first_name, second_name, email, phone_number, language, time, password, password_changed_at, updated_at, created_at
`

type CreateClientParams struct {
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	Email       string `json:"email"`
	PhoneNumber int32  `json:"phone_number"`
	Language    string `json:"language"`
	Password    string `json:"password"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.FirstName,
		arg.SecondName,
		arg.Email,
		arg.PhoneNumber,
		arg.Language,
		arg.Password,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.PhoneNumber,
		&i.Language,
		&i.Time,
		&i.Password,
		&i.PasswordChangedAt,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
