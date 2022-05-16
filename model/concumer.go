package model

type ConsumerAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalCode"`
	Street      string `json:"street"`
}

type Consumer struct {
	Addresses []ConsumerAddress `json:"addresses"`
}
