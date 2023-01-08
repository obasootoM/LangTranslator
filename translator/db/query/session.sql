-- name: CreateSession :one
INSERT INTO sessions (
  id,
  email,
  refresh_token,
  user_agent,
  translator_ip,
  is_blocked,
  expires_at
) VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions 
WHERE  id = $1
LIMIT 1;

-- name: UpdateSession :exec
UPDATE sessions 
set is_blocked = $2
WHERE id = $1;