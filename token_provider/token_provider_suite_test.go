package token_provider_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTokenProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TokenProvider Suite")
}
