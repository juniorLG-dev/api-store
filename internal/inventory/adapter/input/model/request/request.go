package request

type ProductInventoryRequest struct {
	ProductID   string  `json:"product_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}