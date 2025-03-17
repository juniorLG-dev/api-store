package routes

import (
	"loja/internal/cart/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct controller.ControllerGroup) {
	rg.POST("/cart", ct.SaveProduct)
	rg.GET("/cart", ct.GetProducts)
	rg.DELETE("/cart", ct.DeleteProduct)
}