package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ExpiredInvalidToken = errors.New("token as expired")

type Payload struct {
	ID        uuid.UUID `form:"id" json:"id"`
	Username  string    `form:"username" json:"username"`
	IssuedAt  time.Time `form:"issuedat" json:"issuedat"`
	ExpiredAt time.Time `form:"expiredat" json:"expiredat"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := Payload{
		ID:        tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return &payload, err
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ExpiredInvalidToken
	}
	return nil
}
