-- name: CreateTranslator :one
INSERT INTO translator (
    first_name,
    second_name,
    email,
    password,
    profession,
    translator_category,
    rating, 
    certified,
    source_language,
    target_language,
    timezone)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
RETURNING *;