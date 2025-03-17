package repository

import (
	"loja/internal/customer/application/domain"
	"loja/internal/customer/adapter/output/model/db"
	"loja/internal/common/domain/vo"
	
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

type PortRepository interface {
	CreateCustomer(domain.CustomerDomain) error
	GetCustomerByEmail(string) (domain.CustomerDomain, error)
	GetCustomerByID(string) (domain.CustomerDomain, error)
	UsernameExists(string) bool
	EmailExists(string) bool 
	DeleteCustomer(string) error
}

func (r *repository) CreateCustomer(customerDomain domain.CustomerDomain) error {
	customerDB := db.CustomerDB{
		ID: customerDomain.ID.Value,
		Name: customerDomain.Name,
		Username: customerDomain.Username.Value,
		Email: customerDomain.Email,
		Password: customerDomain.Password.Value,
	}

	err := r.db.Create(&customerDB).Error

	return err
}

func (r *repository) GetCustomerByEmail(email string) (domain.CustomerDomain, error) {
	var customer db.CustomerDB
	err := r.db.First(&customer, "email = ?", email).Error

	return domain.CustomerDomain{
		ID: vo.ID{Value: customer.ID},
		Name: customer.Name,
		Username: vo.Username{Value: customer.Username},
		Email: customer.Email,
		Password: vo.Password{Value: customer.Password},
	}, err
}

func (r *repository) GetCustomerByID(id string) (domain.CustomerDomain, error) {
	var customer db.CustomerDB
	err := r.db.First(&customer, "id = ?", id).Error

	return domain.CustomerDomain{
		ID: vo.ID{Value: customer.ID},
		Name: customer.Name,
		Username: vo.Username{Value: customer.Username},
		Email: customer.Email,
		Password: vo.Password{Value: customer.Password},
	}, err
}

func (r *repository) UsernameExists(username string) bool {
	var customer db.CustomerDB
	err := r.db.First(&customer, "username = ?", username).Error
	if err != nil {
		return false
	}

	return true
}

func (r *repository) EmailExists(email string) bool {
	var customer db.CustomerDB
	err := r.db.First(&customer, "email = ?", email).Error
	if err != nil {
		return false
	}

	return true
}

func (r *repository) DeleteCustomer(id string) error {
	return r.db.Delete(&db.CustomerDB{}, "id = ?", id).Error
}