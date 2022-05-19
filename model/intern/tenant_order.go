package intern

import "time"

type Order struct {
	Customer  Customer  `json:"customer"`
	Shipping  Shipping  `json:"shipping"`
	OrderDate time.Time `json:"orderDate"`
	Items     []Item    `json:"items"`
	Status    string    `json:"status"`
	Id        string    `json:"Id"`
}

type Customer struct {
	Addresses []CustomerAddress `json:"Addresses"`
	Email     string            `json:"email"`
}

type PhoneNumber struct {
	Label string `json:"label"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CustomerAddress struct {
	City         string        `json:"city"`
	Country      string        `json:"country"`
	HouseNumber  string        `json:"houseNumber"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers"`
	PostalCode   string        `json:"postalCode"`
	Street       string        `json:"street"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	Salutation   string        `json:"salutation"`
}

type Shipping struct {
	ShippingMethod string    `json:"shippingMethod"`
	TargetTime     time.Time `json:"targetTime"`
	PickUpStore    string    `json:"pickUpStore"`
}

type Article struct {
	ArticleId string `json:"articleId"`
	Name      string `json:"name"`
}

type Item struct {
	Article            Article `json:"article"`
	MeasurementUnitKey string  `json:"measurementUnitKey"`
	Quantity           int     `json:"quantity"`
	ShopPrice          int     `json:"shopPrice"`
}

const SHIPPING_METHOD_DELIVERY = "DELIVERY"
const SHIPPING_METHOD_PICKUP = "PICKUP"
