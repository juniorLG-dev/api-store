package response

type ProductResponse struct {
	ID 					string  `json:"id"`
	Description string  `json:"description"`
	Quantity    int 	  `json:"quantity"`
	Price       float64 `json:"price"`
}