package external

import (
	"time"
)

type OrderStatus string

const (
	OPEN      OrderStatus = "OPEN"
	CANCELLED OrderStatus = "CANCELLED"
	LOCKED    OrderStatus = "LOCKED"
)

type Collect struct {
	FacilityRef         string   `json:"facilityRef"`
	Paid                bool     `json:"paid"`
	SupplyingFacilities []string `json:"supplyingFacilities"`
}

type Shipping struct {
	PreferredCarriers     []string `json:"preferredCarriers,omitempty"`
	PreselectedFacilities []struct {
		FacilityRef string `json:"facilityRef,omitempty"`
	} `json:"preselectedFacilities,omitempty"`
	ServiceLevel string `json:"serviceLevel"`
	Servicetype  string `json:"servicetype"`
}

type DeliveryPreferences struct {
	Collect             []Collect  `json:"collect,omitempty"`
	Shipping            *Shipping  `json:"shipping,omitempty"`
	SupplyingFacilities []string   `json:"supplyingFacilities,omitempty"`
	TargetTime          *time.Time `json:"targetTime,omitempty"`
}

type Order struct {
	Consumer            Consumer            `json:"consumer"`
	DeliveryPreferences DeliveryPreferences `json:"deliveryPreferences"`
	OrderDate           time.Time           `json:"orderDate"`
	OrderLineItems      []OrderLineItem     `json:"orderLineItems"`
	Status              string              `json:"status"`
	TenantOrderId       string              `json:"tenantOrderId"`
}
