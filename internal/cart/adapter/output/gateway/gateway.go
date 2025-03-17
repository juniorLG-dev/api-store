package gateway

import (
	"loja/internal/inventory/adapter/output/model/db"
	
	"gorm.io/gorm"
)

type GatewayOutput struct {
	Description string
	Price 			float64
	SellerID 		string
}

type gateway struct {
	db *gorm.DB
}

func NewCartGateway(db *gorm.DB) *gateway {
	return &gateway{
		db: db,
	}
}

type PortGateway interface {
	GetProductByID(string) (GatewayOutput, error)
}

func (g *gateway) GetProductByID(id string) (GatewayOutput, error) {
	var product db.ProductInventoryDB
	err := g.db.First(&product, "id = ?", id).Error

	return GatewayOutput{
		Description: product.Description,
		Price: product.Price,
		SellerID: product.SellerID,
	}, err
}