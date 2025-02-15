package seller_cache

type InfoSeller struct {
	Name string
	Username string
	Email string
	Password string
	Code string
}

func NewInfoSeller(
	name string,
	username string,
	email string,
	password string,
	code string,
) *InfoSeller {
	return &InfoSeller{
		Name: name,
		Username: username,
		Email: email,
		Password: password,
		Code: code,
	}
}