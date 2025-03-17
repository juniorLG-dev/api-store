package response

type CustomerResponse struct {
	ID 			 string `json:"id"`
	Name 		 string `json:"name"`
	Username string `json:"username"`
	Email 	 string `json:"email"`
}