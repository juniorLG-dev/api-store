package initializer_http

import (
	"loja/internal/seller/adapter/input/controller"
	"loja/internal/seller/adapter/input/routes"
	"loja/internal/seller/adapter/output/cache"
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/seller/adapter/output/smtp"
	"loja/internal/seller/application/decorator"
	"loja/internal/seller/application/dto"
	"loja/internal/seller/application/usecase"
	"loja/internal/seller/application/query"

	"gorm.io/gorm"
	redis "github.com/redis/go-redis/v9"
	"github.com/gin-gonic/gin"

	"os"
)

func InitSeller(db *gorm.DB, rdb *redis.Client, router *gin.RouterGroup) {
	sender := os.Getenv("SENDER")
	pass := os.Getenv("PASS")

	repository := repository.NewSellerRepository(db)
	smtp := smtp.NewSMTP(sender, pass)
	cache := cache.NewCache(rdb)
	registerSeller := usecase.NewUseCaseRegisterSeller(repository, smtp, cache)
	createSeller := usecase.NewUseCaseCreateSeller(repository, smtp, cache)
	loginSeller := usecase.NewUseCaseLoginSeller(repository)
	getSellerByID := decorator.NewTokenVerifier[dto.GetSellerByIDInput, *dto.GetSellerByIDOutput](query.NewQueryGetSellerByID(db))
	getSellerByUsername := query.NewQueryGetSellerByUsername(db)
	controller := controller.NewSellerController(
		*createSeller, 
		getSellerByID,
		*getSellerByUsername,
		*registerSeller,
		*loginSeller,
	)
	routes.InitRoutes(router, controller)
}