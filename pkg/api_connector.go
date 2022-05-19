package pkg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type APIResponse struct {
	Payload    []byte
	StatusCode int
}

type ApiConnector interface {
	Get(path string, params map[string]string) (*APIResponse, error)
	Put(path string, body []byte, params map[string]string) (*APIResponse, error)
	Post(path string, body []byte, params map[string]string) (*APIResponse, error)
	Delete(path string, params map[string]string) (*APIResponse, error)
}

func NewApiConnector(
	client http.Client,
	provider TokenProvider,
	baseUrl string,
) ApiConnector {
	return &apiConnector{
		client:        client,
		baseUrl:       baseUrl,
		tokenProvider: provider,
	}
}

type apiConnector struct {
	client        http.Client
	baseUrl       string
	tokenProvider TokenProvider
}

func (a *apiConnector) Get(path string, params map[string]string) (*APIResponse, error) {
	req, err := a.buildRequest(http.MethodGet, path, nil, params)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	return a.executeRequest(req)
}

func (a *apiConnector) Put(path string, body []byte, params map[string]string) (*APIResponse, error) {
	req, err := a.buildRequest(http.MethodPut, path, bytes.NewReader(body), params)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	return a.executeRequest(req)
}

func (a *apiConnector) Post(path string, body []byte, params map[string]string) (*APIResponse, error) {
	req, err := a.buildRequest(http.MethodPost, path, bytes.NewReader(body), params)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	return a.executeRequest(req)
}

func (a *apiConnector) Delete(path string, params map[string]string) (*APIResponse, error) {
	req, err := a.buildRequest(http.MethodDelete, path, nil, params)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	return a.executeRequest(req)
}

func (a *apiConnector) buildRequest(method string, resourcePath string, body io.Reader, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", a.baseUrl, resourcePath), body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	values := req.URL.Query()
	for key, value := range params {
		values.Add(key, value)
	}
	req.URL.RawQuery = values.Encode()

	jwt, err := a.tokenProvider.GetToken()
	if err != nil {
		return nil, fmt.Errorf("getting token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt.Token))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, nil
}

func (a *apiConnector) executeRequest(req *http.Request) (*APIResponse, error) {
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	return &APIResponse{
		Payload:    body,
		StatusCode: resp.StatusCode,
	}, nil
}
