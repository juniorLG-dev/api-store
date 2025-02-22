package controller

import (
	"loja/internal/inventory/application/usecase"
	"loja/internal/inventory/application/dto"
	"loja/internal/inventory/adapter/input/model/request"
	"loja/internal/inventory/adapter/input/model/response"
	"loja/internal/configuration/handler_err"
	"loja/internal/common/decorator"

	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProductByIDInput = decorator.TokenVerifierInput[dto.GetProductByIDInput]

type controller struct {
	createProduct usecase.CreateProduct
	getProductByID decorator.Query[GetProductByIDInput, dto.GetProductByIDOutput]
}

func NewInventoryController(
	createProduct usecase.CreateProduct,
	getProductByID decorator.Query[decorator.TokenVerifierInput[dto.GetProductByIDInput], dto.GetProductByIDOutput],
) *controller {
	return &controller{
		createProduct: createProduct,
		getProductByID: getProductByID,
	}
}

type ControllerGroupInventory interface {
	CreateProduct(*gin.Context) 
	GetProductByID(*gin.Context)
}

func (ct *controller) CreateProduct(c *gin.Context) {
	var productInventory request.ProductInventoryRequest
	
	if err := c.ShouldBindJSON(&productInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	productInventoryInput := dto.CreateProductInput{
		Description: productInventory.Description,
		Price: productInventory.Price,
		Quantity: productInventory.Quantity,
	}

	if infoErr := ct.createProduct.Run(productInventoryInput, c.Request.Header.Get("Authorization")); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product created"})
}

func (ct *controller) GetProductByID(c *gin.Context) {
	productID := c.Param("id")

	productIDInput := dto.GetProductByIDInput{
		ID: productID,
	}

	product, infoErr := ct.getProductByID.Run(GetProductByIDInput{
		Token: c.Request.Header.Get("Authorization"),
		Data: productIDInput,
	})
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	productOutput := response.ProductInventoryResponse{
		ID: product.ID,
		Description: product.Description,
		Price: product.Price,
		Quantity: product.Quantity,
		SellerID: product.SellerID,
	}

	c.JSON(http.StatusOK, productOutput)
}