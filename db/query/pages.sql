-- name: CreateClientPages :one
INSERT INTO pages(
    source_language,
    target_language,
    field,
    file,
    profession,
    category,
    duration,
    additional_service
)VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;