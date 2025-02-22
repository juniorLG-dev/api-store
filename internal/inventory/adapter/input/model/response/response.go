package response

type ProductInventoryResponse struct {
	ID 					string  `json:"id"`
	Description string  `json:"description"`
	Price 			float64 `json:"price"`
	Quantity 		int     `json:"quantity"`
	SellerID    string  `json:"seller_id"`
}