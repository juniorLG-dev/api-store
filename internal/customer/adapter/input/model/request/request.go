package request

type CustomerRequest struct {
	Name 		 string `json:"name"`
	Username string `json:"username"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type CustomerCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}