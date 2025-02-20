package dto

type CreateProductInput struct {
	Description string
	Price       float64
	Quantity    int
	SellerID    string
}