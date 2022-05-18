package mocks

import "order-importer/token_provider"

type TokenFetcherMock interface {
	token_provider.TokenFetcher
	FetchTokenReturns(jwt *token_provider.JWT, err error)
	FetchTokenCallCount() int
}

func NewTokenFetcherMock() TokenFetcherMock {
	return &tokenFetcherMock{}
}

type tokenFetcherMock struct {
	fetchTokenCallCount int
	jwt                 *token_provider.JWT
	err                 error
}

func (t *tokenFetcherMock) FetchTokenCallCount() int {
	return t.fetchTokenCallCount
}

func (t *tokenFetcherMock) FetchToken() (*token_provider.JWT, error) {
	t.fetchTokenCallCount++
	return t.jwt, t.err
}

func (t *tokenFetcherMock) FetchTokenReturns(jwt *token_provider.JWT, err error) {
	t.jwt = jwt
	t.err = err
}
