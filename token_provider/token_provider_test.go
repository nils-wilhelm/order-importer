package token_provider_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"order-importer/model"
	. "order-importer/token_provider"
)

const (
	CORRECT_PASSWORD = "correctPassword"
	CORRECT_USERNAME = "correctUsername"
	CORRECT_API_KEY  = "correctApiKey"
	CORRECT_TOKEN    = "correctToken"
)

var _ = Describe("Token Provider", func() {
	var tokenProvider TokenProvider
	var server *httptest.Server
	var err error
	var token *JWTToken

	BeforeEach(func() {
		server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			defer request.Body.Close()
			body, _ := io.ReadAll(request.Body)
			var authdata model.TokenAuthBody
			_ = json.Unmarshal(body, &authdata)

			if authdata.Username != CORRECT_USERNAME || authdata.Password != CORRECT_PASSWORD {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte("incorrect auth data"))
				return
			}

			respBody, _ := json.Marshal(model.TokenResponse{
				IdToken:      CORRECT_TOKEN,
				Registered:   true,
				RefreshToken: "",
				ExpiresIn:    "",
			})
			writer.Write(respBody)
		}))
		tokenProvider = NewTokenProvider(
			server.URL,
			CORRECT_API_KEY,
			model.TokenAuthBody{
				Username:          CORRECT_USERNAME,
				Password:          CORRECT_PASSWORD,
				ReturnSecureToken: true,
			},
			http.Client{},
		)
	})

	Context("server responds", func() {

		Context("correct auth data", func() {
			BeforeEach(func() {
				token, err = tokenProvider.GetToken()
			})
			It("returns the correct token", func() {
				Expect(err).To(BeNil())
				Expect(*token).To(Equal(JWTToken(CORRECT_TOKEN)))
			})
		})
		Context("incorrect auth data", func() {
			BeforeEach(func() {
				tokenProvider = NewTokenProvider(
					server.URL,
					"incorrectApiKey",
					model.TokenAuthBody{
						Username:          "incorrectUser",
						Password:          "incorrectPassword",
						ReturnSecureToken: true,
					},
					http.Client{},
				)
				token, err = tokenProvider.GetToken()
			})
			It("returns no token", func() {
				Expect(err).NotTo(BeNil())
				Expect(token).To(BeNil())
			})
		})

	})

	Context("server does not respond", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				// do not respond
			}))
			tokenProvider = NewTokenProvider(server.URL, CORRECT_API_KEY, model.TokenAuthBody{}, http.Client{})
			token, err = tokenProvider.GetToken()
		})
		It("returns no token", func() {
			Expect(err).NotTo(BeNil())
			Expect(token).To(BeNil())
		})
	})
})
