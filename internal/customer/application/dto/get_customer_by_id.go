package dto 

type GetCustomerByIDInput struct {
	ID string
}

type GetCustomerByIDOutput struct {
	ID 			 string
	Name 		 string
	Username string
	Email 	 string
}