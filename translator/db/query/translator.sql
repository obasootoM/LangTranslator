-- name: CreateTranslator :one
INSERT INTO translator (
    first_name,
    second_name,
    email,
    password
    )
VALUES($1,$2,$3,$4)
RETURNING *;

-- name: GetTranslator :one
SELECT * FROM translator 
WHERE  email = $1
LIMIT 1;

-- name: DeleteTranslator :exec
DELETE FROM translator 
 WHERE email = $1;