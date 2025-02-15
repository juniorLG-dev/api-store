package controller 

import (
	"loja/internal/seller/application/dto"
	"loja/internal/seller/application/usecase"
	"loja/internal/seller/adapter/input/model/request"
	"loja/internal/seller/adapter/input/model/response"
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/decorator"

	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

type controller struct {
	createSeller usecase.CreateSeller
	getSellerByID usecase.Usecase[dto.GetSellerByIDInput, dto.GetSellerByIDOutput]
	getSellerByUsername usecase.GetSellerByUsername
	registerSeller usecase.RegisterSeller
	loginSeller usecase.LoginSeller
}

type ControllerGroup interface {
	CreateSeller(c *gin.Context)
	GetSellerByID(c *gin.Context)
	GetSellerByUsername(c *gin.Context)
	RegisterSeller(c *gin.Context)
	LoginSeller(c *gin.Context)
}

func NewSellerController(
	createSeller usecase.CreateSeller,
	getSellerByID usecase.Usecase[dto.GetSellerByIDInput, dto.GetSellerByIDOutput],
	getSellerByUsername usecase.GetSellerByUsername,
	registerSeller usecase.RegisterSeller,
	loginSeller usecase.LoginSeller,
) *controller {
	return &controller{
		createSeller: createSeller,
		getSellerByID: getSellerByID,
		getSellerByUsername: getSellerByUsername,
		registerSeller: registerSeller,
		loginSeller: loginSeller,
	}
}

func (ct *controller) CreateSeller(c *gin.Context) {
	var sellerCodeRequest request.SellerCodeRequest

	if err := c.ShouldBindJSON(&sellerCodeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	sellerInput := dto.CreateSellerInput{
		Email: sellerCodeRequest.Email,
		Code: sellerCodeRequest.Code,
	}

	if infoErr := ct.createSeller.Run(sellerInput); infoErr.Err != nil {
		fmt.Println(infoErr)
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

func (ct *controller) GetSellerByID(c *gin.Context) {
	id := c.Param("id")

	sellerInput := dto.GetSellerByIDInput{
		ID: id,
	}

	fmt.Println(c.Request.Header.Get("Authorization"))
	decorator.SetToken(c.Request.Header.Get("Authorization"))

	seller, infoErr := ct.getSellerByID.Run(sellerInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		c.Abort()
		return
	}

	sellerResponse := response.SellerResponse{
		ID: seller.ID,
		Name: seller.Name,
		Username: seller.Username,
		Email: seller.Email,
	}

	c.JSON(http.StatusOK, sellerResponse)
}

func (ct *controller) GetSellerByUsername(c *gin.Context) {
	username := c.Param("username")

	sellerInput := dto.GetSellerByUsernameInput{
		Username: username,
	}

	seller, infoErr := ct.getSellerByUsername.Run(sellerInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	sellerResponse := response.SellerResponse{
		ID: seller.ID,
		Name: seller.Name,
		Username: seller.Username,
		Email: seller.Email,
	}

	c.JSON(http.StatusOK, sellerResponse)
}

func (ct *controller) RegisterSeller(c *gin.Context) {
	var sellerRequest request.SellerRequest

	if err := c.ShouldBindJSON(&sellerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	sellerInput := dto.RegisterSellerInput{
		Name: sellerRequest.Name,
		Username: sellerRequest.Username,
		Email: sellerRequest.Email,
		Password: sellerRequest.Password,
	}

	if infoErr := ct.registerSeller.Run(sellerInput); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "check your email to see your verification code",
	})
}

func (ct *controller) LoginSeller(c *gin.Context) {
	var sellerRequest request.SellerRequest

	if err := c.ShouldBindJSON(&sellerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	sellerInput := dto.LoginSellerInput{
		Email: sellerRequest.Email,
		Password: sellerRequest.Password,
	}

	token, infoErr := ct.loginSeller.Run(sellerInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}