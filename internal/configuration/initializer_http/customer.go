package initializer_http

import (
	"loja/internal/customer/adapter/input/controller"
	"loja/internal/customer/adapter/input/routes"
	"loja/internal/customer/adapter/output/cache"
	"loja/internal/customer/adapter/output/repository"
	"loja/internal/common/smtp"
	"loja/internal/common/decorator"
	"loja/internal/customer/application/dto"
	"loja/internal/customer/application/usecase"
	"loja/internal/customer/application/query"

	"gorm.io/gorm"
	redis "github.com/redis/go-redis/v9"
	"github.com/gin-gonic/gin"

	"os"
)

func InitCustomer(db *gorm.DB, rdb *redis.Client, router *gin.RouterGroup) {
	sender := os.Getenv("SENDER")
	pass := os.Getenv("PASS")

	repository := repository.NewCustomerRepository(db)
	smtp := smtp.NewSMTP(sender, pass)
	cache := cache.NewCustomerCache(rdb)
	registerCustomer := usecase.NewUseCaseRegisterCustomer(repository, smtp, cache)
	createCustomer := usecase.NewUseCaseCreateCustomer(repository, cache)
	loginCustomer := usecase.NewUseCaseLoginCustomer(repository)
	getCustomerByID := decorator.NewTokenVerifier[dto.GetCustomerByIDInput, dto.GetCustomerByIDOutput](query.NewQueryGetCustomerByID(db))
	getCustomerByUsername := decorator.NewTokenVerifier[dto.GetCustomerByUsernameInput, dto.GetCustomerByUsernameOutput](query.NewQueryGetCustomerByUsername(db))
	deleteCustomer := usecase.NewUseCaseDeleteCustomer(repository)
	controller := controller.NewCustomerController(
		*registerCustomer,
		*createCustomer,
		*loginCustomer,
		getCustomerByID,
		getCustomerByUsername,
		*deleteCustomer,
	)
	routes.InitRoutes(router, controller)
}