package token_provider_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "order-importer/token_provider"
	"time"
)

var _ = Describe("in memory token store", func() {
	var store TokenStore
	var err error
	var outputJWT *JWT

	var inputJWT = JWT{
		Token:      CORRECT_TOKEN,
		ExpireTime: time.Time{},
	}

	BeforeEach(func() {
		store = NewInMemoryTokenStore()
		err = store.SaveToken(inputJWT)
		Expect(err).To(BeNil())
		outputJWT, err = store.GetToken()
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("returns correct value", func() {
		Expect(*outputJWT).To(Equal(inputJWT))
	})
})
