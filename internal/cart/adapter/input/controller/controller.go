package controller

import (
	"loja/internal/cart/adapter/input/model/request"
	"loja/internal/cart/adapter/input/model/response"
	"loja/internal/cart/application/usecase"
	"loja/internal/cart/application/query"
	"loja/internal/cart/application/dto"
	"loja/internal/configuration/handler_err"

	"github.com/gin-gonic/gin"

	"net/http"
)

type controller struct {
	saveProduct usecase.SaveProduct
	getProducts query.GetProducts
	deleteProduct usecase.DeleteProduct
}

func NewCartController(
	saveProduct usecase.SaveProduct,
	getProducts query.GetProducts,
	deleteProduct usecase.DeleteProduct,
) *controller {
	return &controller{
		saveProduct: saveProduct,
		getProducts: getProducts,
		deleteProduct: deleteProduct,
	}
}

type ControllerGroup interface {
	SaveProduct(*gin.Context)
	GetProducts(*gin.Context)
	DeleteProduct(*gin.Context)
}

func (ct *controller) SaveProduct(c *gin.Context) {
	var cartRequest request.CartRequest
	if err := c.ShouldBindJSON(&cartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
	}

	cartInput := dto.SaveProductInput{
		ProductID: cartRequest.ProductID,
	}

	token := c.Request.Header.Get("Authorization")

	if infoErr := ct.saveProduct.Run(cartInput, token); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product added to cart",
	})
}

func (ct *controller) GetProducts(c *gin.Context) {
	products, infoErr := ct.getProducts.Run(c.Request.Header.Get("Authorization"))
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	var productsResponse []response.CartResponse
	for _, product := range products {
		productInfo := response.CartResponse{
			ID: product.ID,
			Description: product.Description,
			Price: product.Price,
			ProductID: product.ProductID,
			CustomerID: product.CustomerID,
			SellerID: product.SellerID,
		}

		productsResponse = append(productsResponse, productInfo)
	}

	c.JSON(http.StatusOK, productsResponse)
}

func (ct *controller) DeleteProduct(c *gin.Context) {
	var cartRequest request.CartRequest
	if err := c.ShouldBindJSON(&cartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	cartInput := dto.DeleteProductInput{
		CartID: cartRequest.CartID,
	}

	if infoErr := ct.deleteProduct.Run(cartInput, c.Request.Header.Get("Authorization")); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})
}