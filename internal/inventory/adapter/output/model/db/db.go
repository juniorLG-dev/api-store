package db

type ProductInventoryDB struct {
	ID 					string `gorm:"primaryKey"`
	Description string
	Price 			float64
	Quantity 		int
	SellerID 		string
}