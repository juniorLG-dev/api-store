package repository

import (
	"loja/internal/inventory/adapter/output/model/db"
	"loja/internal/inventory/application/domain/entities"
	"loja/internal/inventory/application/domain/vo"
	
	"gorm.io/gorm"
)

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *inventoryRepository {
	return &inventoryRepository{
		db: db,
	}
}

type PortRepository interface {
	CreateProduct(entities.ProductInventory) error
	DeleteProduct(string) error
	GetProductByID(string)(entities.ProductInventory, error)
}

func (ir *inventoryRepository) CreateProduct(productInventory entities.ProductInventory) error {
	productInventoryDB := db.ProductInventoryDB{
		ID: productInventory.ID.Value,
		Description: productInventory.Description,
		Price: productInventory.Price.Value,
		Quantity: productInventory.Quantity.Value,
		SellerID: productInventory.SellerID,
	}

	err := ir.db.Create(&productInventoryDB).Error

	return err
}

func (ir *inventoryRepository) DeleteProduct(id string) error {
	err := ir.db.Delete(&db.ProductInventoryDB{}, "id = ?", id).Error
	return err
}

func (ir *inventoryRepository) GetProductByID(id string) (entities.ProductInventory, error) {
	var product db.ProductInventoryDB
	err := ir.db.First(&product, "id = ?", id).Error

	return entities.ProductInventory{
		ID: vo.ID{Value: product.ID},
		Description: product.Description,
		Price: vo.Price{Value: product.Price},
		Quantity: vo.Quantity{Value: product.Quantity},
		SellerID: product.SellerID,
	}, err
}