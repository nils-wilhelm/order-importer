package facility_connector

import (
	. "order-importer/api_connector"
	. "order-importer/model"
)

type FacilityConnector interface {
	GetFacilities() ([]Facility, error)
}

func NewFacilityConnector() FacilityConnector {
	return &facilityConnector{}
}

type facilityConnector struct {
	apiConnector ApiConnector
}

func (f *facilityConnector) GetFacilities() ([]Facility, error) {

	var r *FacilitiesResponse

	r, err := f.apiConnector.Get[FacilitiesResponse]("/facilities")

	return facilities.Facilities, nil
}
