package get_user_info

type Request struct {
	Token string `json:"token"`
}

type Response struct {
	ID        int64  `json:"ID"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	NickName  string `json:"nick_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Amount    int64  `json:"amount"`
	TempToken string `json:"temp_token"`
}
