package dto

type GetSellerByIDInput struct {
	ID string
}

type GetSellerByIDOutput struct {
	ID 			 string
	Name 		 string
	Username string
	Email 	 string
}