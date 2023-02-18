package domain

type Product struct {
	Sku         string  `json:"sku"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
