package controller

import (
	"loja/internal/inventory/application/usecase"
	"loja/internal/inventory/application/dto"
	"loja/internal/inventory/application/query"
	"loja/internal/inventory/adapter/input/model/request"
	"loja/internal/inventory/adapter/input/model/response"
	"loja/internal/configuration/handler_err"

	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	createProduct usecase.CreateProduct
	getProducts query.GetProducts
	deleteProduct usecase.DeleteProduct
}

func NewInventoryController(
	createProduct usecase.CreateProduct,
	getProducts query.GetProducts,
	deleteProduct usecase.DeleteProduct,
) *controller {
	return &controller{
		createProduct: createProduct,
		getProducts: getProducts,
		deleteProduct: deleteProduct,
	}
}

type ControllerGroupInventory interface {
	CreateProduct(*gin.Context) 
	GetProducts(*gin.Context)
	DeleteProduct(*gin.Context)
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

func (ct *controller) GetProducts(c *gin.Context) {
	sellerID := c.Param("id")

	productIDInput := dto.GetProductsInput{
		SellerID: sellerID,
	}

	products, infoErr := ct.getProducts.Run(productIDInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	var productsOutput []response.ProductInventoryResponse
	for _, product := range products {
		productOutput := response.ProductInventoryResponse{
			ID: product.ID,
			Description: product.Description,
			Price: product.Price,
			Quantity: product.Quantity,
			SellerID: product.SellerID,
		}
		productsOutput = append(productsOutput, productOutput)
	}

	c.JSON(http.StatusOK, productsOutput)
}

func (ct *controller) DeleteProduct(c *gin.Context) {
	var productRequest request.ProductInventoryRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	productInput := dto.DeleteProductInput{
		ID: productRequest.ProductID,
	}

	if infoErr := ct.deleteProduct.Run(productInput, c.Request.Header.Get("Authorization")); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})
}