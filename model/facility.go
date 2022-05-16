package model

import "time"

type Facility struct {
	Address struct {
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
		PostalCode     string `json:"postalCode"`
		Street         string `json:"street"`
		CompanyName    string `json:"companyName"`
		EmailAddresses []struct {
			Recipient string `json:"recipient"`
			Value     string `json:"value"`
		} `json:"emailAddresses"`
		ResolvedCoordinates struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"resolvedCoordinates"`
		ResolvedTimeZone struct {
			OffsetInSeconds int    `json:"offsetInSeconds"`
			TimeZoneId      string `json:"timeZoneId"`
			TimeZoneName    string `json:"timeZoneName"`
		} `json:"resolvedTimeZone"`
	} `json:"address"`
	ClosingDays []struct {
		Date       time.Time `json:"date"`
		Reason     string    `json:"reason"`
		Recurrence string    `json:"recurrence"`
	} `json:"closingDays"`
	Contact struct {
		CustomAttributes struct {
		} `json:"customAttributes"`
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		RoleDescription string `json:"roleDescription"`
	} `json:"contact"`
	CustomAttributes struct {
	} `json:"customAttributes"`
	FulfillmentProcessBuffer string `json:"fulfillmentProcessBuffer"`
	LocationType             string `json:"locationType"`
	Name                     string `json:"name"`
	PickingTimes             struct {
		Friday []struct {
			End struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"end"`
			Start struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"start"`
		} `json:"friday"`
		Monday []struct {
			End struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"end"`
			Start struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"start"`
		} `json:"monday"`
		Saturday []struct {
			End struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"end"`
			Start struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"start"`
		} `json:"saturday"`
		Sunday []struct {
			End struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"end"`
			Start struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"start"`
		} `json:"sunday"`
		Thursday []struct {
			End struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"end"`
			Start struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"start"`
		} `json:"thursday"`
		Tuesday []struct {
			End struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"end"`
			Start struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"start"`
		} `json:"tuesday"`
		Wednesday []struct {
			End struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"end"`
			Start struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"start"`
		} `json:"wednesday"`
	} `json:"pickingTimes"`
	Services []struct {
		Type string `json:"type"`
	} `json:"services"`
	Status           string    `json:"status"`
	TenantFacilityId string    `json:"tenantFacilityId"`
	Created          time.Time `json:"created"`
	LastModified     time.Time `json:"lastModified"`
	Version          int       `json:"version"`
	Configs          []struct {
		Ref string `json:"ref"`
		Rel string `json:"rel"`
	} `json:"configs"`
	Id string `json:"id"`
}

type FacilitiesResponse struct {
	Facilities []Facility `json:"facilities"`
}
