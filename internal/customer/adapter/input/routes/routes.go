package routes

import (
	"loja/internal/customer/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct controller.ControllerGroup) {
	rg.POST("/register/customer", ct.RegisterCustomer)
	rg.POST("/verify/customer", ct.CreateCustomer)
	rg.POST("/login/customer", ct.LoginCustomer)
	rg.GET("/customer/:id", ct.GetCustomerByID)
	rg.GET("/customer/username/:username", ct.GetCustomerByUsername)
	rg.DELETE("/delete/customer", ct.DeleteCustomer)
}
