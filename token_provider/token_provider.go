package token_provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	. "order-importer/model"
)

type JWTToken string

type TokenProvider interface {
	GetToken() (*JWTToken, error)
}

func NewTokenProvider(
	credentialUrl string,
	apiKey string,
	bodyData TokenAuthBody,
	client http.Client,
) TokenProvider {
	return &tokenProvider{
		credentialUrl: credentialUrl,
		apiKey:        apiKey,
		bodyData:      bodyData,
		client:        client,
	}
}

type tokenProvider struct {
	credentialUrl string
	apiKey        string
	bodyData      TokenAuthBody
	client        http.Client
}

func (t *tokenProvider) GetToken() (*JWTToken, error) {
	req, err := t.buildRequest()
	if err != nil {
		return nil, fmt.Errorf("could not build request: %w", err)
	}
	body, err := t.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("could not do request: %w", err)
	}
	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal token response: %w", err)
	}
	token := JWTToken(tokenResponse.IdToken)
	return &token, nil
}

func (t *tokenProvider) doRequest(req *http.Request) ([]byte, error) {
	// TODO retry
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not do request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body data: %w", err)
	}
	return body, nil
}

func (t *tokenProvider) buildRequest() (*http.Request, error) {
	postBody, err := json.Marshal(t.bodyData)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request body: %w", err)
	}
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, t.credentialUrl, responseBody)
	if err != nil {
		return nil, fmt.Errorf("could not create http request: %w", err)
	}
	req.Header.Set("Authorization", t.apiKey)
	return req, nil
}
