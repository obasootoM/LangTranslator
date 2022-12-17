-- name: CreateClient :one
INSERT INTO client (
  first_name,second_name,email,password,phone_number,language
) VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;
-- name: GetEmail :one
SELECT * FROM client 
WHERE  email = $1
LIMIT 1;