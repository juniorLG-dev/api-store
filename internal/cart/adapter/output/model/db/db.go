package db

type CartDB struct {
	ID 					string `gorm:"primaryKey"`
	Description string
	Price 			float64
	ProductID 	string
	CustomerID  string
	SellerID 		string
}