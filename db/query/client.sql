-- name: CreateClient :one
INSERT INTO client (
  first_name,
  second_name,
  email,
  password,
  phone_number,
  language,
  currency,
  time
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetClient :one
SELECT * FROM client 
WHERE  email = $1
LIMIT 1;