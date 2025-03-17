package repository

import (
	"loja/internal/cart/adapter/output/model/db"
	"loja/internal/cart/application/domain/entities"
	"loja/internal/cart/application/domain/vo"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

type PortRepository interface {
	SaveProduct(entities.CartDomain) error
	DeleteProduct(string) error
	GetProductByID(string) (entities.CartDomain, error)
}

func (r *repository) SaveProduct(cartDomain entities.CartDomain) error {
	cartDB := db.CartDB{
		ID: cartDomain.ID.Value,
		Description: cartDomain.Description,
		Price: cartDomain.Price.Value,
		ProductID: cartDomain.ProductID,
		CustomerID: cartDomain.CustomerID,
		SellerID: cartDomain.SellerID,
	}

	err := r.db.Create(&cartDB).Error

	return err
}

func (r *repository) DeleteProduct(id string) error {
	return r.db.Delete(db.CartDB{}, "id = ?", id).Error
}

func (r *repository) GetProductByID(id string) (entities.CartDomain, error) {
	var product db.CartDB
	err := r.db.First(&product, "id = ?", id).Error

	return entities.CartDomain{
		ID: vo.ID{Value: product.ID},
		Description: product.Description,
		Price: vo.Price{Value: product.Price},
		ProductID: product.ProductID,
		CustomerID: product.CustomerID,
		SellerID: product.SellerID,
	}, err
}