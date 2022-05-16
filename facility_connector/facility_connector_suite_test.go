package facility_connector_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFacilityConnector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FacilityConnector Suite")
}
