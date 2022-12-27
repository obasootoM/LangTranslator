package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto   *paseto.V2
	symetric []byte
}

func NewPasetoMaker(symetric string) (Maker, error) {
	if len(symetric) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid key size must be exactly %d symetric", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:   paseto.NewV2(),
		symetric: []byte(symetric),
	}
	return maker, nil
}

// CreateToken implements Maker
func (paseto *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return paseto.paseto.Encrypt(paseto.symetric, payload, nil)
}

// VerifyToken implements Maker
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symetric, payload, nil)
	if err != nil {
		return nil, paseto.ErrInvalidTokenAuth
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, err
}
