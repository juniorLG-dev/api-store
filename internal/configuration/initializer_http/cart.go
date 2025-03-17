package initializer_http

import (
	"loja/internal/cart/adapter/output/repository"
	"loja/internal/cart/application/usecase"
	"loja/internal/cart/adapter/input/controller"
	"loja/internal/cart/adapter/input/routes"
	"loja/internal/cart/application/query"
	"loja/internal/cart/adapter/output/gateway"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func InitCart(db *gorm.DB, router *gin.RouterGroup) {
	cartRepository := repository.NewCartRepository(db)
	cartGateway := gateway.NewCartGateway(db)
	saveProduct := usecase.NewUseCaseSaveProduct(cartRepository, cartGateway)
	getProducts := query.NewQueryGetProducts(db)
	deleteProduct := usecase.NewUseCaseDeleteProduct(cartRepository)
	cartController := controller.NewCartController(
		*saveProduct,
		*getProducts,
		*deleteProduct,
	)

	routes.InitRoutes(router, cartController)
}