package api_connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	. "net/url"
	. "order-importer/token_provider"
)

type ApiConnector interface {
	Get[T any](resourcePath string) (*T, error)
	Post[T any](resourcePath string, payload []byte) (*T, error)
	Put[T any](resourcePath string, payload []byte) (*T, error)
}

func NewApiConnector(
	client http.Client,
	baseUrl string,
	token JWTToken,
) ApiConnector {
	return &apiConnector{
		client:  client,
		baseUrl: baseUrl,
		token:   token,
	}
}

type apiConnector struct {
	client  http.Client
	baseUrl string
	token   JWTToken
}

func (a *apiConnector) Get[T any](resourcePath string) (*T, error) {
	resp, err := a.client.Get(a.baseUrl + resourcePath)
	if err != nil {
		return nil, fmt.Errorf("getting resource from api: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body data: %w", err)
	}
	var respPayload T
	err = json.Unmarshal(body, &respPayload)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response data: %w", err)
	}
	return &respPayload, nil
}

func (a *apiConnector) Post[T any](resourcePath string, payload []byte) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (a *apiConnector) Put[T any](resourcePath string, payload []byte) (*T, error) {
	//TODO implement me
	panic("implement me")
}
