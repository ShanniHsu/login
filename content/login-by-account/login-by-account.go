package login_by_account

type Request struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
