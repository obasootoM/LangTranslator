-- name: CreateProfile :one
INSERT INTO profile (
  id,
  name,
  gender,
  email,
  phone_number,
  address_line,
  country,
  native_language
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;