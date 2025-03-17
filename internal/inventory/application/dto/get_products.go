package dto

type GetProductsInput struct {
	SellerID string
}

type GetProductsOutput struct {
	ID 					string
	Description string
	Price 			float64
	Quantity 		int
	SellerID    string
}