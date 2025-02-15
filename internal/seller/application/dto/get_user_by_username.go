package dto

type GetSellerByUsernameInput struct {
	Username string
}

type GetSellerByUsernameOutput struct {
	ID 			 string
	Name 		 string
	Username string
	Email 	 string
}