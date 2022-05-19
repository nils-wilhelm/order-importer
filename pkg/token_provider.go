package token_provider

import (
	"fmt"
	"time"
)

type TokenProvider interface {
	GetToken() (*JWT, error)
}

func NewTokenProvider(
	store TokenStore,
	fetcher TokenFetcher,
) TokenProvider {
	return &tokenProvider{
		store:   store,
		fetcher: fetcher,
	}
}

type tokenProvider struct {
	store   TokenStore
	fetcher TokenFetcher
}

func (t *tokenProvider) GetToken() (*JWT, error) {
	tokenFromStore, err := t.store.GetToken()
	if err != nil {
		return nil, fmt.Errorf("loading token from store: %w", err)
	}
	fmt.Println("token from store: ", tokenFromStore.ExpireTime)
	fmt.Println("now: ", time.Now())
	if tokenFromStore.ExpireTime.After(time.Now()) {
		fmt.Println("token still valid. use from store")
		return tokenFromStore, nil
	}
	fmt.Println("token not valid anymore. fetch new")
	fetchedToken, err := t.fetcher.FetchToken()
	if err != nil {
		return nil, fmt.Errorf("fetching token: %w", err)
	}
	err = t.store.SaveToken(*fetchedToken)
	if err != nil {
		fmt.Printf("error saving token: %s", err.Error())
	}
	return fetchedToken, nil
}
