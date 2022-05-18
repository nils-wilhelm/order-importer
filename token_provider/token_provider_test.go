package token_provider_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"order-importer/mocks"
	. "order-importer/token_provider"
)

var _ = Describe("token provider", func() {
	var tokenProvider TokenProvider
	var tokenStore TokenStore
	var tokenFetcher mocks.TokenFetcherMock

	var jwt *JWT
	var err error

	var recentJWT = &JWT{
		Token:      CORRECT_TOKEN,
		ExpireTime: time.Now().Add(time.Hour),
	}
	var fetchError = errors.New("fetching jwt")

	BeforeEach(func() {
		tokenStore = NewInMemoryTokenStore()
		tokenFetcher = mocks.NewTokenFetcherMock()
		tokenProvider = NewTokenProvider(tokenStore, tokenFetcher)
		tokenFetcher.FetchTokenReturns(recentJWT, nil)
	})
	Context("store is empty", func() {
		BeforeEach(func() {
			jwt, err = tokenProvider.GetToken()
		})
		It("calls fetcher", func() {
			Expect(tokenFetcher.FetchTokenCallCount()).To(Equal(1))
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns correct token", func() {
			Expect(jwt.Token).To(Equal(CORRECT_TOKEN))
		})
	})
	Context("store contains recent token", func() {
		BeforeEach(func() {
			_ = tokenStore.SaveToken(JWT{
				Token:      CORRECT_TOKEN,
				ExpireTime: time.Now().Add(time.Hour),
			})
			jwt, err = tokenProvider.GetToken()
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("does not call fetcher", func() {
			Expect(tokenFetcher.FetchTokenCallCount()).To(Equal(0))
		})
		It("returns recent token", func() {
			Expect(jwt).NotTo(BeNil())
			Expect(jwt.Token).To(Equal(CORRECT_TOKEN))
			Expect(jwt.ExpireTime.After(time.Now())).To(BeTrue())
		})
	})
	Context("store contains expired token", func() {
		BeforeEach(func() {
			_ = tokenStore.SaveToken(JWT{
				Token:      CORRECT_TOKEN,
				ExpireTime: time.Now().Add(-time.Hour),
			})
		})
		Context("fetcher is successful", func() {
			BeforeEach(func() {
				tokenFetcher.FetchTokenReturns(recentJWT, nil)
				jwt, err = tokenProvider.GetToken()
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("does call fetcher", func() {
				Expect(tokenFetcher.FetchTokenCallCount()).To(Equal(1))
			})
			It("returns recent token", func() {
				Expect(jwt).NotTo(BeNil())
				Expect(jwt.Token).To(Equal(CORRECT_TOKEN))
				Expect(jwt.ExpireTime.After(time.Now())).To(BeTrue())
			})
		})
		Context("fetcher is not successful", func() {
			BeforeEach(func() {
				tokenFetcher.FetchTokenReturns(nil, fetchError)
				jwt, err = tokenProvider.GetToken()
			})
			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
			})
			It("does call fetcher", func() {
				Expect(tokenFetcher.FetchTokenCallCount()).To(Equal(1))
			})
			It("returns no token", func() {
				Expect(jwt).To(BeNil())
			})
		})
	})
})
