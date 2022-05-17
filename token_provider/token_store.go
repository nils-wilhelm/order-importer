package token_provider

import (
	"time"
)

type TokenStore interface {
	GetToken() (*JWT, error)
	SaveToken(token JWT) error
}

func NewInMemoryTokenStore() TokenStore {
	return &inMemoryTokenStore{}
}

type inMemoryTokenStore struct {
	tokenData JWT
}

func (i *inMemoryTokenStore) SaveToken(tokenData JWT) error {
	i.tokenData = tokenData
	return nil
}

func (i *inMemoryTokenStore) GetToken() (*JWT, error) {
	return &i.tokenData, nil
}

type JWT struct {
	Token      string
	ExpireTime time.Time
}
