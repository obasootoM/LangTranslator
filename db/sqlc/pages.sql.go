// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: pages.sql

package db

import (
	"context"
)

const createClientPages = `-- name: CreateClientPages :one
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
RETURNING id, source_language, target_language, file, profession, category, field, duration, additional_service
`

type CreateClientPagesParams struct {
	SourceLanguage    string `json:"source_language"`
	TargetLanguage    string `json:"target_language"`
	Field             string `json:"field"`
	File              string `json:"file"`
	Profession        string `json:"profession"`
	Category          string `json:"category"`
	Duration          string `json:"duration"`
	AdditionalService string `json:"additional_service"`
}

func (q *Queries) CreateClientPages(ctx context.Context, arg CreateClientPagesParams) (Page, error) {
	row := q.db.QueryRowContext(ctx, createClientPages,
		arg.SourceLanguage,
		arg.TargetLanguage,
		arg.Field,
		arg.File,
		arg.Profession,
		arg.Category,
		arg.Duration,
		arg.AdditionalService,
	)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.SourceLanguage,
		&i.TargetLanguage,
		&i.File,
		&i.Profession,
		&i.Category,
		&i.Field,
		&i.Duration,
		&i.AdditionalService,
	)
	return i, err
}
