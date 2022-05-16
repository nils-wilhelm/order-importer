package model

import "time"

type OrderStatus string

const (
	OPEN      OrderStatus = "OPEN"
	CANCELLED OrderStatus = "CANCELLED"
	LOCKED    OrderStatus = "LOCKED"
)

type OrderId string

type TOrder struct {
	Consumer       Consumer                   `json:"consumer"`
	OrderDate      time.Time                  `json:"orderDate"`
	OrderLineItems []OrderLineItemForCreation `json:"orderLineItems"`
	Status         OrderStatus                `json:"status"`
	Id             OrderId                    `json:"id"`
}

type OrderForCreation struct {
	Consumer  Consumer  `json:"consumer"`
	OrderDate time.Time `json:"orderDate"`
}

type Order struct {
	Consumer struct {
		Addresses []struct {
			AdditionalAddressInfo string `json:"additionalAddressInfo"`
			City                  string `json:"city"`
			Country               string `json:"country"`
			CustomAttributes      struct {
			} `json:"customAttributes"`
			HouseNumber  string `json:"houseNumber"`
			PhoneNumbers []struct {
				CustomAttributes struct {
				} `json:"customAttributes"`
				Label string `json:"label"`
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"phoneNumbers"`
			PostalCode  string `json:"postalCode"`
			Street      string `json:"street"`
			CompanyName string `json:"companyName"`
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			Salutation  string `json:"salutation"`
		} `json:"addresses"`
		CustomAttributes struct {
		} `json:"customAttributes"`
		Email string `json:"email"`
	} `json:"consumer"`
	CustomAttributes struct {
	} `json:"customAttributes"`
	DeliveryPreferences struct {
		Collect []struct {
			FacilityRef         string   `json:"facilityRef"`
			Paid                bool     `json:"paid"`
			SupplyingFacilities []string `json:"supplyingFacilities"`
		} `json:"collect"`
		Shipping struct {
			PreferredCarriers     []string `json:"preferredCarriers"`
			PreselectedFacilities []struct {
				FacilityRef string `json:"facilityRef"`
			} `json:"preselectedFacilities"`
			ServiceLevel string `json:"serviceLevel"`
			Servicetype  string `json:"servicetype"`
		} `json:"shipping"`
		SupplyingFacilities []string  `json:"supplyingFacilities"`
		TargetTime          time.Time `json:"targetTime"`
	} `json:"deliveryPreferences"`
	OrderDate      time.Time `json:"orderDate"`
	OrderLineItems []struct {
		Article struct {
			CustomAttributes struct {
			} `json:"customAttributes"`
			ImageUrl        string `json:"imageUrl"`
			TenantArticleId string `json:"tenantArticleId"`
			Title           string `json:"title"`
			Attributes      []struct {
				Category string `json:"category"`
				Key      string `json:"key"`
				Priority int    `json:"priority"`
				Value    string `json:"value"`
			} `json:"attributes"`
		} `json:"article"`
		CustomAttributes struct {
		} `json:"customAttributes"`
		MeasurementUnitKey string   `json:"measurementUnitKey"`
		Quantity           int      `json:"quantity"`
		ScannableCodes     []string `json:"scannableCodes"`
		ShopPrice          int      `json:"shopPrice"`
	} `json:"orderLineItems"`
	Status        string `json:"status"`
	TenantOrderId string `json:"tenantOrderId"`
}
