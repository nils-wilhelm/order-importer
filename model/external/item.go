package external

type OrderLineItem struct {
	Article            Article  `json:"article"`
	MeasurementUnitKey string   `json:"measurementUnitKey"`
	Quantity           int      `json:"quantity"`
	ScannableCodes     []string `json:"scannableCodes,omitempty"`
	ShopPrice          int      `json:"shopPrice"`
}

type ArticleAttributes struct {
	Category string `json:"category"`
	Key      string `json:"key"`
	Priority int    `json:"priority"`
	Value    string `json:"value"`
}

type Article struct {
	ImageUrl        string              `json:"imageUrl,omitempty"`
	TenantArticleId string              `json:"tenantArticleId"`
	Title           string              `json:"title"`
	Attributes      []ArticleAttributes `json:"attributes,omitempty"`
}

type OrderLineItemForCreation struct {
	Article  OrderLineItemArticle `json:"article"`
	Quantity int64                `json:"quantity"`
}

type OrderLineItemArticle struct {
	TenantArticleId string `json:"tenantArticleId"`
	Title           string `json:"title"`
}
