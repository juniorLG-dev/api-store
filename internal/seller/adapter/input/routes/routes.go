package routes

import (
	"loja/internal/seller/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct controller.ControllerGroup) {
	rg.POST("/register/seller", ct.RegisterSeller)
	rg.POST("/verify/seller", ct.CreateSeller)
	rg.POST("/login/seller", ct.LoginSeller)
	rg.GET("/seller/:id", ct.GetSellerByID)
	rg.GET("/seller/username/:username", ct.GetSellerByUsername)
	rg.DELETE("/delete/seller", ct.DeleteSeller)
}