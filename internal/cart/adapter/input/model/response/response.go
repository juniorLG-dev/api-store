package response

type CartResponse struct {
	ID 					string 	`json:"id"`
	Description string 	`json:"description"`
	Price 			float64 `json:"price"`
	ProductID 	string  `json:"product_id"`
	CustomerID 	string  `json:"customer_id"`
	SellerID 		string  `json:"seller_id"`
}