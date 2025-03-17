package controller

import (
	"loja/internal/customer/application/usecase"
	"loja/internal/customer/application/dto"
	"loja/internal/customer/adapter/input/model/request"
	"loja/internal/customer/adapter/input/model/response"
	"loja/internal/common/decorator"
	"loja/internal/configuration/handler_err"

	"github.com/gin-gonic/gin"

	"net/http"
)

type GetCustomerByUsernameInput = decorator.TokenVerifierInput[dto.GetCustomerByUsernameInput]
type GetCustomerByIDInput = decorator.TokenVerifierInput[dto.GetCustomerByIDInput]

type controller struct {
	registerCustomer usecase.RegisterCustomer
	createCustomer usecase.CreateCustomer
	loginCustomer usecase.LoginCustomer
	getCustomerByID decorator.Query[GetCustomerByIDInput, dto.GetCustomerByIDOutput]
	getCustomerByUsername decorator.Query[GetCustomerByUsernameInput, dto.GetCustomerByUsernameOutput]
	deleteCustomer usecase.DeleteCustomer
}

func NewCustomerController(
	registerCustomer usecase.RegisterCustomer,
	createCustomer usecase.CreateCustomer,
	loginCustomer usecase.LoginCustomer,
	getCustomerByID decorator.Query[decorator.TokenVerifierInput[dto.GetCustomerByIDInput], dto.GetCustomerByIDOutput],
	getCustomerByUsername decorator.Query[decorator.TokenVerifierInput[dto.GetCustomerByUsernameInput], dto.GetCustomerByUsernameOutput],
	deleteCustomer usecase.DeleteCustomer,
) *controller {
	return &controller{
		registerCustomer: registerCustomer,
		createCustomer: createCustomer,
		loginCustomer: loginCustomer,
		getCustomerByID: getCustomerByID,
		getCustomerByUsername: getCustomerByUsername,
		deleteCustomer: deleteCustomer,
	}
}

type ControllerGroup interface {
	RegisterCustomer(c *gin.Context)
	CreateCustomer(c *gin.Context)
	LoginCustomer(c *gin.Context)
	GetCustomerByID(c *gin.Context)
	GetCustomerByUsername(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

func (ct *controller) RegisterCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	customerInput := dto.RegisterCustomerInput{
		Name: customerRequest.Name,
		Username: customerRequest.Username,
		Email: customerRequest.Email,
		Password: customerRequest.Password,
	}

	infoErr := ct.registerCustomer.Run(customerInput)
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "check your email to see your verification code",
	})
}

func (ct *controller) CreateCustomer(c *gin.Context) {
	var customerCodeRequest request.CustomerCodeRequest
	if err := c.ShouldBindJSON(&customerCodeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	customerCodeInput := dto.CreateCustomerInput{
		Email: customerCodeRequest.Email,
		Code: customerCodeRequest.Code,
	}

	if infoErr := ct.createCustomer.Run(customerCodeInput); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

func (ct *controller) LoginCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
	}

	customerInput := dto.LoginCustomerInput{
		Email: customerRequest.Email,
		Password: customerRequest.Password,
	}

	token, infoErr := ct.loginCustomer.Run(customerInput)
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

func (ct *controller) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")

	customerInput := dto.GetCustomerByIDInput{
		ID: id,
	}

	customer, infoErr := ct.getCustomerByID.Run(GetCustomerByIDInput{
		Token: c.Request.Header.Get("Authorization"),
		Data: customerInput,
	})
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		c.Abort()
		return
	}

	customerResponse := response.CustomerResponse{
		ID: customer.ID,
		Name: customer.Name,
		Username: customer.Username,
		Email: customer.Email,
	}

	c.JSON(http.StatusOK, customerResponse)
}

func (ct *controller) GetCustomerByUsername(c *gin.Context) {
	username := c.Param("username")

	customerInput := dto.GetCustomerByUsernameInput{
		Username: username,
	}

	customer, infoErr := ct.getCustomerByUsername.Run(GetCustomerByUsernameInput{
		Token: c.Request.Header.Get("Authorization"),
		Data: customerInput,
	})
	if infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		c.Abort()
		return
	}

	customerResponse := response.CustomerResponse{
		ID: customer.ID,
		Name: customer.Name,
		Username: customer.Username,
		Email: customer.Email,
	}

	c.JSON(http.StatusOK, customerResponse)
}

func (ct *controller) DeleteCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid fields",
		})
		return
	}

	customerInput := dto.DeleteCustomerInput{
		Password: customerRequest.Password,
	}

	if infoErr := ct.deleteCustomer.Run(customerInput, c.Request.Header.Get("Authorization")); infoErr.Err != nil {
		msgErr := handler_err.HandlerErr(infoErr)
		c.JSON(msgErr.Status, msgErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}