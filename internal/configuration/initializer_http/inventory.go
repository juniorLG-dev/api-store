package initializer_http

import (
	"loja/internal/inventory/adapter/output/repository"
	"loja/internal/inventory/application/usecase"
	"loja/internal/inventory/adapter/input/controller"
	"loja/internal/inventory/adapter/input/routes"
	"loja/internal/inventory/application/query"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func InitInventory(db *gorm.DB, router *gin.RouterGroup) {
	inventoryRepository := repository.NewInventoryRepository(db)
	createProduct := usecase.NewUseCaseCreateProduct(inventoryRepository)
	getProducts := query.NewQueryGetProducts(db)
	deleteProduct := usecase.NewUseCaseDeleteProduct(inventoryRepository)
	inventoryController := controller.NewInventoryController(
		*createProduct,
		*getProducts,
		*deleteProduct,
	)

	routes.InitRoutes(router, inventoryController)
}