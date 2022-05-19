package tests

import (
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"order-importer/model"
	"order-importer/model/external"
	. "order-importer/pkg"
)

const (
	CORRECT_PASSWORD = "correctPassword"
	CORRECT_USERNAME = "correctUsername"
	CORRECT_API_KEY  = "correctApiKey"
	CORRECT_TOKEN    = "correctToken"
)

var _ = Describe("Token Fetcher", func() {
	var tokenProvider TokenFetcher
	var server *httptest.Server
	var err error
	var jwt *JWT

	BeforeEach(func() {
		server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			defer request.Body.Close()
			body, _ := io.ReadAll(request.Body)
			var authdata model.TokenAuthBody
			_ = json.Unmarshal(body, &authdata)

			if authdata.Email != CORRECT_USERNAME || authdata.Password != CORRECT_PASSWORD {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte("incorrect auth data"))
				return
			}

			respBody, _ := json.Marshal(external.TokenResponse{
				IdToken:      CORRECT_TOKEN,
				Registered:   true,
				RefreshToken: "",
				ExpiresIn:    "3600",
			})
			writer.Write(respBody)
		}))
		tokenProvider = NewTokenFetcher(
			server.URL,
			CORRECT_API_KEY,
			model.TokenAuthBody{
				Email:             CORRECT_USERNAME,
				Password:          CORRECT_PASSWORD,
				ReturnSecureToken: true,
			},
			http.Client{},
		)
	})

	Context("server responds", func() {

		Context("correct auth data", func() {
			BeforeEach(func() {
				jwt, err = tokenProvider.FetchToken()
			})
			It("returns the correct Token", func() {
				Expect(err).To(BeNil())
				Expect(jwt.Token).To(Equal((CORRECT_TOKEN)))
			})
		})
		Context("incorrect auth data", func() {
			BeforeEach(func() {
				tokenProvider = NewTokenFetcher(
					server.URL,
					"incorrectApiKey",
					model.TokenAuthBody{
						Email:             "incorrectUser",
						Password:          "incorrectPassword",
						ReturnSecureToken: true,
					},
					http.Client{},
				)
				jwt, err = tokenProvider.FetchToken()
			})
			It("returns no Token", func() {
				Expect(err).NotTo(BeNil())
				Expect(jwt).To(BeNil())
			})
		})

	})

	Context("server does not respond", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				// do not respond
			}))
			tokenProvider = NewTokenFetcher(server.URL, CORRECT_API_KEY, model.TokenAuthBody{}, http.Client{})
			jwt, err = tokenProvider.FetchToken()
		})
		It("returns no Token", func() {
			Expect(err).NotTo(BeNil())
			Expect(jwt).To(BeNil())
		})
	})
})
