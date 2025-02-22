package initializer_http

import (
	"loja/internal/inventory/adapter/output/repository"
	"loja/internal/inventory/application/usecase"
	"loja/internal/inventory/adapter/input/controller"
	"loja/internal/inventory/adapter/input/routes"
	"loja/internal/inventory/application/query"
	"loja/internal/inventory/application/dto"
	"loja/internal/common/decorator"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func InitInventory(db *gorm.DB, router *gin.RouterGroup) {
	inventoryRepository := repository.NewInventoryRepository(db)
	createProduct := usecase.NewUseCaseCreateProduct(inventoryRepository)
	getProductByID := decorator.NewTokenVerifier[dto.GetProductByIDInput, dto.GetProductByIDOutput](query.NewQueryGetProductByID(db))
	inventoryController := controller.NewInventoryController(
		*createProduct,
		getProductByID,
	)

	routes.InitRoutes(router, inventoryController)
}