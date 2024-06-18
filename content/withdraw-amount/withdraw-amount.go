package withdraw_amount

type Request struct {
	Token          string `json:"token"`
	WithdrawAmount string `json:"withdraw_amount"`
}

type Response struct {
	TempToken   string `json:"temp_token"`
	TotalAmount string `json:"total_amount"`
}
