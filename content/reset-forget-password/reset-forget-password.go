package reset_forget_password

type Request struct {
	Code     string `json:"code"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
