-- name: CreateClient :one
INSERT INTO client (
  first_name,second_name,email,phone_number,language,password
) VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;