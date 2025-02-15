package db

type SellerDB struct {
	ID string `gorm:"primaryKey"`
	Name string
	Username string
	Email string
	Password string
}

type ProductDB struct {
	ID string `gorm:"primaryKey"`
	Description string
	Quantity int
	Price float64
	SellerID string
}