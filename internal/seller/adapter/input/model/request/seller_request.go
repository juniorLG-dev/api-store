package request

type SellerRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SellerCodeRequest struct {
	Email string `json:"email"` 
	Code  string `json:"code"`
}
