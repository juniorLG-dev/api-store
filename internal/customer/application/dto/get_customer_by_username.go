package dto

type GetCustomerByUsernameInput struct {
	Username string
}

type GetCustomerByUsernameOutput struct {
	ID 			 string
	Name 		 string
	Username string
	Email 	 string
}