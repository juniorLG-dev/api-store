package controller

import (
	"loja/internal/inventory/application/usecase"
	"loja/internal/inventory/application/dto"
	"loja/internal/inventory/adapter/input/model/request"
	"loja/internal/inventory/adapter/input/model/response"
	"loja/internal/configuration/handler_err"
	"loja/internal/inventory/application/query"

	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	createProduct usecase.CreateProduct
	getProducts query.GetProducts
}

func NewInventoryController(
	createProduct usecase.CreateProduct,
	getProducts query.GetProducts,
) *controller {
	return &controller{
		createProduct: createProduct,
		getProducts: getProducts,
	}
}

type ControllerGroupInventory interface {
	CreateProduct(*gin.Context) 
	GetProducts(*gin.Context)
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

	sellerIDInput := dto.GetProductsInput{
		SellerID: sellerID,
	}

	products, infoErr := ct.getProducts.Run(sellerIDInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	var productsResponse []response.ProductInventoryResponse
	for _, product := range products {
		productInfo := response.ProductInventoryResponse{
			ID: product.ID,
			Description: product.Description,
			Price: product.Price,
			Quantity: product.Quantity,
		}

		productsResponse = append(productsResponse, productInfo)
	}

	c.JSON(http.StatusOK, productsResponse)
}