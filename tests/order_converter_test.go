package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"order-importer/model/external"
	"order-importer/model/intern"
	. "order-importer/pkg"
	"time"
)

var inputPhoneNumbers = []intern.PhoneNumber{
	{
		Type:  "MOBILE",
		Value: "123456789",
	},
}
var inputAddresses = []intern.CustomerAddress{
	{
		City:         "Testhausen",
		Country:      "DE",
		HouseNumber:  "1",
		PhoneNumbers: inputPhoneNumbers,
		PostalCode:   "12345",
		Street:       "Teststrasse",
		FirstName:    "Hans",
		LastName:     "Testmann",
		Salutation:   "Herr",
	},
}
var inputCustomer = intern.Customer{
	Addresses: inputAddresses,
	Email:     "test@example.com",
}
var inputDeliveryShipping = intern.Shipping{
	ShippingMethod: intern.SHIPPING_METHOD_DELIVERY,
}

var inputPickUpShipping = intern.Shipping{
	ShippingMethod: intern.SHIPPING_METHOD_PICKUP,
	PickUpStore:    "STORE1",
}
var inputItems = []intern.Item{
	{
		Article: intern.Article{
			ArticleId: "ARTICLE1",
			Name:      "Test Article",
		},
		MeasurementUnitKey: "litre",
		Quantity:           4,
		ShopPrice:          120,
	},
}

var inputOrderDelivery = intern.Order{
	Customer:  inputCustomer,
	Shipping:  inputDeliveryShipping,
	OrderDate: time.Time{},
	Items:     inputItems,
	Status:    "OPEN",
}
var inputOrderPickUp = intern.Order{
	Customer:  inputCustomer,
	Shipping:  inputPickUpShipping,
	OrderDate: time.Time{},
	Items:     inputItems,
	Status:    "OPEN",
}

var expectedConsumerAddresses = []external.ConsumerAddress{
	{
		City:        "Testhausen",
		Country:     "DE",
		HouseNumber: "1",
		PhoneNumbers: []external.PhoneNumber{
			{
				Type:  "MOBILE",
				Value: "123456789",
			},
		},
		PostalCode: "12345",
		Street:     "Teststrasse",
		FirstName:  "Hans",
		LastName:   "Testmann",
		Salutation: "Herr",
	},
}

var expectedShippingDeliveryPreferences = external.DeliveryPreferences{
	Collect: nil,
	Shipping: &external.Shipping{
		ServiceLevel: "DELIVERY",
		Servicetype:  "BEST_EFFORT",
	},
	TargetTime: &time.Time{},
}

var expectedOrderLineItems = []external.OrderLineItem{
	{
		Article: external.Article{
			TenantArticleId: "ARTICLE1",
			Title:           "Test Article",
		},
		MeasurementUnitKey: "litre",
		Quantity:           2,
		ShopPrice:          120,
	},
}

var expectedShippingCollectPreferences = external.DeliveryPreferences{
	Collect: []external.Collect{
		{
			FacilityRef:         "STORE1",
			Paid:                true,
			SupplyingFacilities: []string{},
		},
	},
	TargetTime: &time.Time{},
}

var _ = Describe("order converter", func() {
	var orderConverter OrderConverter
	var exportOrder external.Order

	BeforeEach(func() {
		orderConverter = NewOrderConverter()
	})

	Context("shipping delivery", func() {
		BeforeEach(func() {
			exportOrder = orderConverter.Convert(inputOrderDelivery)
		})
		It("converts customer correctly", func() {
			Expect(exportOrder).NotTo(BeNil())
			Expect(exportOrder.Consumer).NotTo(BeNil())
			Expect(exportOrder.Consumer.Email).To(Equal("test@example.com"))
			Expect(exportOrder.Consumer.Addresses).To(Equal(expectedConsumerAddresses))
		})
		It("converts shipping correctly", func() {
			Expect(exportOrder).NotTo(BeNil())
			Expect(exportOrder.DeliveryPreferences).NotTo(BeNil())
			Expect(exportOrder.DeliveryPreferences.Shipping).To(Equal(expectedShippingDeliveryPreferences.Shipping))
			Expect(exportOrder.DeliveryPreferences.Collect).To(BeNil())
			Expect(exportOrder.DeliveryPreferences.SupplyingFacilities).To(BeNil())
		})
		It("converts items correctly", func() {
			Expect(exportOrder).NotTo(BeNil())
			Expect(exportOrder.OrderLineItems).NotTo(BeNil())
			Expect(exportOrder.OrderLineItems).To(HaveLen(1))
			Expect(exportOrder.OrderLineItems[0].ShopPrice).To(Equal(120))
			Expect(exportOrder.OrderLineItems[0].MeasurementUnitKey).To(Equal("litre"))
			Expect(exportOrder.OrderLineItems[0].ScannableCodes).To(BeNil())
			Expect(exportOrder.OrderLineItems[0].Article).NotTo(BeNil())
			Expect(exportOrder.OrderLineItems[0].Article.TenantArticleId).To(Equal("ARTICLE1"))
			Expect(exportOrder.OrderLineItems[0].Article.Title).To(Equal("Test Article"))
			Expect(exportOrder.OrderLineItems[0].Article.Attributes).To(HaveLen(0))
			Expect(exportOrder.OrderLineItems[0].Article.ImageUrl).To(Equal(""))
		})
		It("converts customer correctly", func() {
			Expect(exportOrder).NotTo(BeNil())
		})
	})

	Context("pickup", func() {
		BeforeEach(func() {
			exportOrder = orderConverter.Convert(inputOrderPickUp)
		})
		It("converts shipping correctly", func() {
			Expect(exportOrder).NotTo(BeNil())
			Expect(exportOrder.DeliveryPreferences).NotTo(BeNil())
			Expect(exportOrder.DeliveryPreferences.Collect).To(Equal(expectedShippingCollectPreferences.Collect))
			Expect(exportOrder.DeliveryPreferences.Shipping).To(BeNil())
			Expect(exportOrder.DeliveryPreferences.SupplyingFacilities).To(BeNil())
		})
	})

})
