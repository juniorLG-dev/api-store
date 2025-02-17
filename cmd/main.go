package main

import (
	"loja/internal/configuration/cache_config"
	"loja/internal/configuration/database"
	"loja/internal/seller/adapter/input/controller"
	"loja/internal/seller/adapter/input/routes"
	"loja/internal/seller/adapter/output/cache"
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/seller/adapter/output/smtp"
	"loja/internal/seller/application/decorator"
	"loja/internal/seller/application/dto"
	"loja/internal/seller/application/usecase"
	"loja/internal/seller/application/query"

	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := database.SetupDB()
	if err != nil {
		log.Panic(err)
	}

	rdb := cache_config.SetupCacheDB("localhost:6379", "", 0)

	sender := os.Getenv("SENDER")
	pass := os.Getenv("PASS")

	repository := repository.NewSellerRepository(db)
	smtp := smtp.NewSMTP(sender, pass)
	cache := cache.NewCache(&rdb)
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
	routes.InitRoutes(&router.RouterGroup, controller)

	router.Run(":8080")
}