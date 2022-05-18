package order_converter

import (
	"time"

	"github.com/google/uuid"

	"order-importer/model/external"
	"order-importer/model/intern"
)

type OrderConverter interface {
	Convert(order intern.Order) external.Order
}

func NewOrderConverter() OrderConverter {
	return &orderConverter{}
}

type orderConverter struct {
}

func (o *orderConverter) Convert(order intern.Order) external.Order {

	return external.Order{
		Consumer: external.Consumer{
			Addresses: o.addresses(order),
			Email:     order.Customer.Email,
		},
		DeliveryPreferences: o.deliveryPrefs(order),
		OrderDate:           time.Now(),
		OrderLineItems:      o.items(order),
		Status:              "OPEN",
		TenantOrderId:       uuid.New().String(),
	}

}

func (o *orderConverter) deliveryPrefs(order intern.Order) external.DeliveryPreferences {
	deliveryPreferences := external.DeliveryPreferences{}

	if order.Shipping.ShippingMethod == intern.SHIPPING_METHOD_DELIVERY {
		deliveryPreferences.Shipping = external.Shipping{
			ServiceLevel: "DELIVERY",
			Servicetype:  "BEST_EFFORT",
		}
	} else {
		deliveryPreferences.Collect = []external.Collect{
			{
				FacilityRef: order.Shipping.PickUpStore,
				Paid:        true,
			},
		}
	}
	deliveryPreferences.TargetTime = o.targetTime()

	return deliveryPreferences
}

func (o *orderConverter) targetTime() time.Time {
	now := time.Now()
	year, month, day := now.Add(time.Hour * 24 * 2).Date()
	return time.Date(year, month, day, 12, 0, 0, 0, now.Location())
}

func (o *orderConverter) items(order intern.Order) []external.OrderLineItem {
	var orderLineItems []external.OrderLineItem

	for _, item := range order.Items {
		orderLineItems = append(orderLineItems, external.OrderLineItem{
			Article: external.Article{
				TenantArticleId: item.Article.ArticleId,
				Title:           item.Article.Name,
			},
			MeasurementUnitKey: item.MeasurementUnitKey,
			Quantity:           item.Quantity,
			ShopPrice:          item.ShopPrice,
		})
	}
	return orderLineItems
}

func (o *orderConverter) addresses(order intern.Order) []external.ConsumerAddress {
	var addresses []external.ConsumerAddress

	for _, address := range order.Customer.Addresses {

		addresses = append(addresses, external.ConsumerAddress{
			City:         address.City,
			Country:      address.Country,
			HouseNumber:  address.HouseNumber,
			PhoneNumbers: o.phoneNumbers(address),
			PostalCode:   address.PostalCode,
			Street:       address.Street,
			FirstName:    address.FirstName,
			LastName:     address.LastName,
			Salutation:   address.Salutation,
		})
	}
	return addresses
}

func (o *orderConverter) phoneNumbers(address intern.CustomerAddress) []external.PhoneNumber {
	var numbers []external.PhoneNumber
	for _, number := range address.PhoneNumbers {
		numbers = append(numbers, external.PhoneNumber{
			Type:  number.Type,
			Value: number.Value,
		})
	}
	return numbers
}
