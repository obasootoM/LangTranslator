-- name: CreateProfile :one
INSERT INTO profile (
  name,
  gender,
  email,
  phone_number,
  address_line,
  country,
  native_language
) VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;


-- name: UpdateProfile :one
UPDATE profile 
  set name =$7,
  address_line = $6,
  gender = $5,
  email = $4,
  phone_number = $3,
  country = $2 
WHERE id = $1
RETURNING *;

-- name: GetProfile :one
SELECT * FROM profile 
WHERE  email = $1
LIMIT 1;

-- name: ListProfile :many
SELECT * FROM profile
ORDER BY id
LIMIT $1
OFFSET $2;
