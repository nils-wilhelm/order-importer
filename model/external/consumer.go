package external

type ConsumerAddress struct {
	AdditionalAddressInfo string        `json:"additionalAddressInfo"`
	City                  string        `json:"city"`
	Country               string        `json:"country"`
	HouseNumber           string        `json:"houseNumber"`
	PhoneNumbers          []PhoneNumber `json:"phoneNumbers"`
	PostalCode            string        `json:"postalCode"`
	Street                string        `json:"street"`
	CompanyName           string        `json:"companyName"`
	FirstName             string        `json:"firstName"`
	LastName              string        `json:"lastName"`
	Salutation            string        `json:"salutation"`
}

type Consumer struct {
	Addresses []ConsumerAddress `json:"addresses"`
	Email     string            `json:"email"`
}

type PhoneNumber struct {
	Label string `json:"label"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
