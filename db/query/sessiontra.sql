-- name: CreateSessionTrans :one
INSERT INTO sessionsTrans (
  id,
  email,
  refresh_token,
  user_agent,
  translator_ip,
  is_blocked,
  expires_at
) VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: GetSessionTrans :one
SELECT * FROM sessionstrans 
WHERE  email = $1
LIMIT 1;

-- name: UpdateSessionTrans :exec
UPDATE sessionstrans 
set is_blocked = $2,
email = $3
WHERE id = $1
RETURNING *;