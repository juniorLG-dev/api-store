package routes

import (
	"loja/internal/inventory/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct controller.ControllerGroupInventory) {
	rg.POST("/product", ct.CreateProduct)
	rg.GET("/seller/:id/products", ct.GetProducts)
}