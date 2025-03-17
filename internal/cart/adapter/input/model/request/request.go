package request

type CartRequest struct {
	CartID    string `json:"cart_id"`
	ProductID string `json:"product_id"`
}