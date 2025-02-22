package dto

type GetProductByIDInput struct {
	ID string
}

type GetProductByIDOutput struct {
	ID 					string
	Description string
	Price 			float64
	Quantity 		int
	SellerID    string
}