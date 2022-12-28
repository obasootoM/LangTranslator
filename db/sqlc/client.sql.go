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
  first_name,
  second_name,
  email,
  password
) VALUES ($1,$2,$3,$4)
RETURNING id, first_name, second_name, email, password, updated_at, created_at
`

type CreateClientParams struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.FirstName,
		arg.SecondName,
		arg.Email,
		arg.Password,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.Password,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteClient = `-- name: DeleteClient :exec
DELETE FROM client 
 WHERE email = $1
`

func (q *Queries) DeleteClient(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteClient, email)
	return err
}

const getClient = `-- name: GetClient :one
SELECT id, first_name, second_name, email, password, updated_at, created_at FROM client 
WHERE  email = $1
LIMIT 1
`

func (q *Queries) GetClient(ctx context.Context, email string) (Client, error) {
	row := q.db.QueryRowContext(ctx, getClient, email)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.Password,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
