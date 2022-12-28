-- name: CreateClient :one
INSERT INTO client (
  first_name,
  second_name,
  email,
  password
) VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetClient :one
SELECT * FROM client 
WHERE  email = $1
LIMIT 1;

-- name: DeleteClient :exec
DELETE FROM client 
 WHERE email = $1;