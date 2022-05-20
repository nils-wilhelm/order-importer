package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	. "order-importer/model/auth"
	"strconv"
	"time"
)

type TokenFetcher interface {
	FetchToken() (*JWT, error)
}

func NewTokenFetcher(
	credentialUrl string,
	apiKey string,
	bodyData TokenRequestPayload,
	client http.Client,
) TokenFetcher {
	return &tokenFetcher{
		credentialUrl: credentialUrl,
		apiKey:        apiKey,
		bodyData:      bodyData,
		client:        client,
	}
}

type tokenFetcher struct {
	credentialUrl string
	apiKey        string
	bodyData      TokenRequestPayload
	client        http.Client
}

func (t *tokenFetcher) FetchToken() (*JWT, error) {
	req, err := t.buildRequest()
	if err != nil {
		return nil, fmt.Errorf("could not build request: %w", err)
	}
	body, err := t.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("could not do request: %w", err)
	}
	return t.parseTokenResponse(body)
}

func (t *tokenFetcher) parseTokenResponse(body []byte) (*JWT, error) {
	var tokenResponse TokenResponsePayload
	err := json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal Token response: %w", err)
	}
	expireSeconds, err := strconv.Atoi(tokenResponse.ExpiresIn)
	if err != nil {
		return nil, fmt.Errorf("parsing expiresIn: %w", err)
	}
	jwt := JWT{
		Token:      tokenResponse.IdToken,
		ExpireTime: time.Now().Add(time.Second * time.Duration(expireSeconds)),
	}
	return &jwt, nil
}

func (t *tokenFetcher) buildRequest() (*http.Request, error) {
	postBody, err := json.Marshal(t.bodyData)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request body: %w", err)
	}
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, t.credentialUrl, responseBody)
	if err != nil {
		return nil, fmt.Errorf("could not create http request: %w", err)
	}
	values := req.URL.Query()
	values.Add("key", t.apiKey)
	req.URL.RawQuery = values.Encode()
	return req, nil
}

func (t *tokenFetcher) doRequest(req *http.Request) ([]byte, error) {
	// TODO retry
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not do request: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request faield: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body data: %w", err)
	}
	return body, nil
}
