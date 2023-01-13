// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID         int64     `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type Order struct {
	ID                       int64  `json:"id"`
	SourceLanguage           string `json:"source_language"`
	TargetLanguage           string `json:"target_language"`
	Translator               string `json:"translator"`
	ProofReader              string `json:"proof_reader"`
	TranslationDelivaryDate  string `json:"translation_delivary_date"`
	ProofReadingDelivaryDate string `json:"proof_reading_delivary_date"`
	ProjectEndDate           string `json:"project_end_date"`
	ServiceLevel             string `json:"service_level"`
	Profession               string `json:"profession"`
	TranslatorCategory       string `json:"translator_category"`
	DelivarySpeed            string `json:"delivary_speed"`
	TranslatorRequest        string `json:"translator_request"`
	DelivaryAddress          string `json:"delivary_address"`
}

type Profile struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Gender         string    `json:"gender"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
	AddressLine    string    `json:"address_line"`
	Country        string    `json:"country"`
	NativeLanguage string    `json:"native_language"`
	CreatedAt      time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}
