-- name: CreateOrders :one
INSERT INTO orders (
  source_language,
  target_language,
  translator,
  proof_reader,
  translation_delivary_date,
  proof_reading_delivary_date,
  project_end_date,
  service_level,
  profession,
  translator_category,
  delivary_speed,
  translator_request,
  delivary_address
)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
RETURNING *;