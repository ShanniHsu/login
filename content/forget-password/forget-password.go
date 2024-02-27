package forget_password

type Request struct {
	Account string `json:"account"`
	Email   string `json:"email"`
}
