package main

import (
	"log"
	"net/http"
	"order-importer/model/auth"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"order-importer/pkg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("loading .env file")
	}

	email := os.Getenv("FULFILLMENT_EMAIL")
	password := os.Getenv("FULFILLMENT_PASSWORD")
	apiKey := os.Getenv("FULFILLMENT_API_KEY")
	apiUrl := os.Getenv("FULFILLMENT_API_URL")
	credentialUrl := os.Getenv("FULFILLMENT_CREDENTIAL_URL")

	tokenProvider := pkg.NewTokenProvider(
		pkg.NewInMemoryTokenStore(),
		pkg.NewTokenFetcher(
			credentialUrl,
			apiKey,
			auth.TokenRequestPayload{
				Email:             email,
				Password:          password,
				ReturnSecureToken: true,
			},
			http.Client{},
		),
	)

	apiConnector := pkg.NewApiConnector(http.Client{}, tokenProvider, apiUrl)

	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})
	r.Handle("/orders", pkg.NewOrderHandler(apiConnector, pkg.NewOrderConverter()))
	http.ListenAndServe("localhost:8080", r)
}
