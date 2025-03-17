package customer_cache

type InfoCustomer struct {
	Name string
	Username string
	Email string
	Password string
	Code string
}

func NewInfoCustomer(
	name string,
	username string,
	email string,
	password string,
	code string,
) *InfoCustomer {
	return &InfoCustomer{
		Name: name,
		Username: username,
		Email: email,
		Password: password,
		Code: code,
	}
}
