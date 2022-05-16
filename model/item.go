package model

type OrderLineItem struct {
	Article  OrderLineItemArticle `json:"article"`
	Quantity int64                `json:"quantity"`
	Id       string               `json:"id"`
}

type OrderLineItemForCreation struct {
	Article  OrderLineItemArticle `json:"article"`
	Quantity int64                `json:"quantity"`
}

type OrderLineItemArticle struct {
	TenantArticleId string `json:"tenantArticleId"`
	Title           string `json:"title"`
}
