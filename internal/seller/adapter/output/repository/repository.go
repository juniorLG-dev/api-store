package repository

import (
	"loja/internal/seller/application/domain"
	"loja/internal/seller/adapter/output/model/db"
	"loja/internal/common/domain/vo"

	"gorm.io/gorm"
)

type PortRepository interface {
	CreateSeller(domain.SellerDomain)	error
	GetSellerByEmail(string) (domain.SellerDomain, error)
	GetSellerByID(string) (domain.SellerDomain, error)
	UsernameExists(string) bool
	EmailExists(string) bool
	DeleteSeller(string) error
}

type repository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateSeller(sellerDomain domain.SellerDomain) error {
	seller := db.SellerDB{
		ID: sellerDomain.ID.Value,
		Name: sellerDomain.Name,
		Username: sellerDomain.Username.Value,
		Email: sellerDomain.Email,
		Password: sellerDomain.Password.Value,
	}

	err := r.db.Create(&seller).Error

	return err
}

func (r *repository) GetSellerByEmail(email string) (domain.SellerDomain, error) {
	var seller db.SellerDB
	err := r.db.First(&seller, "email = ?", email).Error

	sellerOutputDomain := domain.SellerDomain{
		ID: vo.ID{Value: seller.ID},
		Name: seller.Name,
		Username: vo.Username{Value: seller.Username},
		Email: seller.Email,
		Password: vo.Password{Value: seller.Password},
	}

	return sellerOutputDomain, err
}

func (r *repository) GetSellerByID(id string) (domain.SellerDomain, error) {
	var seller db.SellerDB
	err := r.db.First(&seller, "id = ?", id).Error

	sellerOutputDomain := domain.SellerDomain{
		ID: vo.ID{Value: seller.ID},
		Name: seller.Name,
		Username: vo.Username{Value: seller.Username},
		Email: seller.Email,
		Password: vo.Password{Value: seller.Password},
	}

	return sellerOutputDomain, err
}

func (r *repository) UsernameExists(username string) bool {
	var seller db.SellerDB

	err := r.db.First(&seller, "username = ?", username).Error
	if err != nil {
		return false
	}

	return true
}

func (r *repository) EmailExists(email string) bool {
	var seller db.SellerDB

	err := r.db.First(&seller, "email = ?", email).Error
	if err != nil {
		return false
	}
		
	return true
}

func (r *repository) DeleteSeller(id string) error {
	err := r.db.Delete(&db.SellerDB{}, "id = ?", id).Error
	return err
}
