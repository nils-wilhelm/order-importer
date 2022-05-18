package main

import (
	"log"
	"net/http"
	"order-importer/order_converter"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"order-importer/api_connector"
	"order-importer/handlers"
	"order-importer/model"
	. "order-importer/token_provider"
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

	tokenProvider := NewTokenProvider(
		NewInMemoryTokenStore(),
		NewTokenFetcher(
			credentialUrl,
			apiKey,
			model.TokenAuthBody{
				Email:             email,
				Password:          password,
				ReturnSecureToken: true,
			},
			http.Client{},
		),
	)

	apiConnector := api_connector.NewApiConnector(http.Client{}, tokenProvider, apiUrl)

	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})
	r.Handle("/orders", handlers.NewOrderHandler(apiConnector, order_converter.NewOrderConverter()))
	http.ListenAndServe("localhost:8080", r)
}
