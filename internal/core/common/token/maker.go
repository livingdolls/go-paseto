package token

import (
	"time"

	"github.com/livingdolls/go-paseto/internal/core/entity"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

// CreateToken implements Maker.
func (p *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)

	if err != nil {
		return "", nil
	}

	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

// VerifyToken implements Maker.
func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)

	if err != nil {
		return nil, entity.ErrInvalidToken
	}

	err = payload.Valid()

	if err != nil {
		return nil, entity.ErrExpiredToken
	}

	return payload, nil
}
