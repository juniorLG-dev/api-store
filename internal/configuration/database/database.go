package database

import (
	sellerDB "loja/internal/seller/adapter/output/model/db"
	
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func SetupDB() (*gorm.DB, error) {
	database, err := gorm.Open(
		sqlite.Open("loja.db"),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	if err = database.AutoMigrate(
		&sellerDB.SellerDB{}, 
		&sellerDB.ProductDB{}, 
	); err != nil {
		return nil, err
	}

	return database, nil
}