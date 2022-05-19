package mocks

import (
	"order-importer/pkg"
)

type TokenFetcherMock interface {
	pkg.TokenFetcher
	FetchTokenReturns(jwt *pkg.JWT, err error)
	FetchTokenCallCount() int
}

func NewTokenFetcherMock() TokenFetcherMock {
	return &tokenFetcherMock{}
}

type tokenFetcherMock struct {
	fetchTokenCallCount int
	jwt                 *pkg.JWT
	err                 error
}

func (t *tokenFetcherMock) FetchTokenCallCount() int {
	return t.fetchTokenCallCount
}

func (t *tokenFetcherMock) FetchToken() (*pkg.JWT, error) {
	t.fetchTokenCallCount++
	return t.jwt, t.err
}

func (t *tokenFetcherMock) FetchTokenReturns(jwt *pkg.JWT, err error) {
	t.jwt = jwt
	t.err = err
}
